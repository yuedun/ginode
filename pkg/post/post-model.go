package post

import "time"

type Post struct {
	ID          int       `json:"id"`
	Mobile      string    `json:"mobile"`
	UserName    string    `json:"userName"`
	Gender      string    `json:"gender"`
	Addr        string    `json:"addr"`
	Email       string    `json:"email"`
	Status      int       `json:"status"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// 设置User的表名为`post`
func (Post) TableName() string {
	return "post"
}
