package post

import (
	"fmt"

	"gorm.io/gorm"
)

type (
	PostService interface {
		GetPostInfo() (user Post, err error)
		GetPostInfoBySQL() (user Post, err error)
		CreatePost(user *Post) (err error)
		UpdatePost(userID int, user *Post) (err error)
		DeletePost(userID int) (err error)
	}
)
type postService struct {
	mysql *gorm.DB
}

func NewService(mysql *gorm.DB) PostService {
	return &postService{
		mysql: mysql,
	}
}

func (u *postService) GetPostInfo() (post Post, err error) {
	err = u.mysql.First(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func (u *postService) GetPostInfoBySQL() (post Post, err error) {
	err = u.mysql.Raw("select * from post where id=?", post.ID).Scan(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func (u *postService) CreatePost(post *Post) (err error) {
	err = u.mysql.Create(post).Error
	fmt.Println(post)
	if err != nil {
		return err
	}
	return nil
}

func (u *postService) UpdatePost(userID int, post *Post) (err error) {
	err = u.mysql.Model(post).Where("id = ?", userID).Updates(post).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *postService) DeletePost(userID int) (err error) {
	u.mysql.Where("id = ?", userID).Delete(Post{})
	if err != nil {
		return err
	}
	return nil
}
