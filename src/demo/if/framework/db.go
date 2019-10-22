package framework

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"sync"
)

func init() {
	_, _ = GetDB()
}

var db *xorm.Engine
var dbOnce sync.Once

func GetDB() (*xorm.Engine, error) {
	var err error
	dbOnce.Do(func() {
		db, err = xorm.NewEngine("mysql", "root:123456@tcp(localhost:3333)/test?charset=utf8&parseTime=True&loc=Local")
	})
	return db, err
}

//var db *sqlx.DB
//var dbOnce sync.Once
//
//func GetDB() (*sqlx.DB, error) {
//	var err error
//	dbOnce.Do(func() {
//		db, err = sqlx.Open("mysql", "root:123456@tcp(localhost:3333)/test?charset=utf8&parseTime=True&loc=Local")
//	})
//	return db, err
//}
