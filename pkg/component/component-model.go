package component

import "time"

type Component struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// 设置Component的表名为`component`
func (Component) TableName() string {
	return "component"
}
