package framework

import (
	"github.com/jinzhu/gorm"
	"sync"
)

var db *gorm.DB
var dbOnce sync.Once

func GetDB() (*gorm.DB, error) {
	var err error
	dbOnce.Do(func() {
		db, err = gorm.Open("mysql", "root:123456@tcp(localhost:3333)/test?charset=utf8&parseTime=True&loc=Local")
	})
	return db, err
}
