package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v3"
)

//profile variables
type Conf struct {
	MysqlURL string `yaml:mysqlURL`
}

func GetConf(filename string) (*Conf, error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var c = new(Conf)
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return c, nil
}

/**
 * md5加密
 */
func GetMD5(password string) string {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(password))
	Result := Md5Inst.Sum(nil)
	// 以下两种输出结果一样
	fmt.Println("格式化>>>>>>>%x\n", Result)
	fmt.Println("hex解码>>>>>>>", hex.EncodeToString(Result), "\n")
	return fmt.Sprintf("%x", Result)
}

/**
 * 生成密码
 */
func GeneratePassword(mobile string) string {
	b := []byte(mobile)
	p := b[7:]
	password := "hello" + string(p)
	return GetMD5(password)
}
