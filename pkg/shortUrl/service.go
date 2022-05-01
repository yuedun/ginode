package shortUrl

import (
	"crypto/md5"
	"fmt"

	"gorm.io/gorm"
)

type (
	ShortURLService interface {
		GetLongByShort(username string) (shortUrl ShortUrl, err error)
		Long2Short(req *Long2ShortRequest) (shortUrl string, err error)
	}
)
type shortURLService struct {
	mysql *gorm.DB
}

func NewService(mysql *gorm.DB) ShortURLService {
	return &shortURLService{
		mysql: mysql,
	}
}

// 根据短链获取长链
func (this *shortURLService) GetLongByShort(shortURL string) (shortUrl ShortUrl, err error) {
	if err = this.mysql.Where("short_url = ?", shortURL).Find(&shortUrl).Error; err != nil {
		return shortUrl, err
	}
	return shortUrl, nil
}

// 长链转短链
func (this *shortURLService) Long2Short(req *Long2ShortRequest) (shortUrl string, err error) {
	urlMd5 := fmt.Sprintf("%x", md5.Sum([]byte(req.OriginUrl)))
	var short ShortUrl
	short.HashCode = urlMd5
	err = this.mysql.Find(&short, short).Error
	if err != nil || short.ID == 0 {
		err = nil
		// 数据库中没有记录，重新生成一个新的短url
		shortUrl, errRet := this.generateShortUrl(req, urlMd5)
		if errRet != nil {
			return "", errRet
		}
		return shortUrl, nil
	}
	return short.ShortUrl, nil
}

func (this *shortURLService) generateShortUrl(req *Long2ShortRequest, hashcode string) (shortUrl string, err error) {
	var shortRecord ShortUrl
	shortRecord.OriginUrl = req.OriginUrl
	shortRecord.HashCode = hashcode
	result := this.mysql.Create(&shortRecord)
	if result.Error != nil {
		return "", err
	}
	// 0-9a-zA-Z 六十二进制（base62）
	insertId := shortRecord.ID
	// mongodb的_id不能转换成int64，此处需要使用mysql自增id实现，或使用自定义mysqldb 自增id
	shortUrlTrans := this.transTo62(int64(insertId))
	this.mysql.First(&shortRecord)
	shortRecord.ShortUrl = shortUrlTrans
	err = this.mysql.Save(shortRecord).Error
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return shortRecord.ShortUrl, nil
}

// 将十进制转换为62进制   0-9a-zA-Z 六十二进制
func (this *shortURLService) transTo62(id int64) string {
	// 1 -- > 1
	// 10-- > a
	// 61-- > Z
	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var shortUrl []byte
	for {
		var result byte
		number := id % 62
		result = charset[number]
		var tmp []byte
		tmp = append(tmp, result)
		shortUrl = append(tmp, shortUrl...)
		id = id / 62
		if id == 0 {
			break
		}
	}
	return string(shortUrl)
}
