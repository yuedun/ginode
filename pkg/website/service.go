package website

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type (
	WebsiteService interface {
		GetWebsiteList(offset, limit int, search Website) (website []Website, err error)
		CreateWebsite(website *Website) (err error)
		UpdateWebsite(website *Website) (err error)
		DeleteWebsite(websiteID int) (err error)
	}
)
type websiteService struct {
	mysql *gorm.DB
}

func NewService(mysql *gorm.DB) WebsiteService {
	return &websiteService{
		mysql: mysql,
	}
}

func (u *websiteService) GetWebsiteList(offset, limit int, search Website) (websites []Website, err error) {
	if search.Name != "" {
		u.mysql = u.mysql.Where("name like ?", search.Name+"%")
	}
	if search.Category != "" {
		u.mysql = u.mysql.Where("category =?", search.Category)
	}
	err = u.mysql.Offset(offset).Limit(limit).Find(&websites).Error
	if err != nil {
		return websites, err
	}
	return websites, nil
}

func (u *websiteService) CreateWebsite(website *Website) (err error) {
	err = u.mysql.Create(website).Error
	fmt.Println(website)
	if err != nil {
		return err
	}
	return nil
}

func (u *websiteService) UpdateWebsite(website *Website) (err error) {
	err = u.mysql.Model(website).Update(website).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *websiteService) DeleteWebsite(websiteID int) (err error) {
	u.mysql.Where("id = ?", websiteID).Delete(Website{})
	if err != nil {
		return err
	}
	return nil
}
