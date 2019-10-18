package models

import "time"

type Blog struct {
	ID       int
	Title    string
	Date     time.Time `gorm:"default:'0000-00-00'"`
	Content  string
	Type     int
	Views    int
	IsDelete int
}

func (Blog) TableName() string {
	return "blog"
}
