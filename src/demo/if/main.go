package main

import (
	_ "Demo/if/routers"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Blog2 struct {
	Id    int
	Title string
}

func main() {
	//icecream := framework.GetIcecream()
	//icecream.DB, _ = gorm.Open("mysql", "root:123456@tcp(localhost:3333)/test?charset=utf8&parseTime=True&loc=Local")
	//defer icecream.DB.Close()

	var db *sqlx.DB
	db, _ = sqlx.Open("mysql", "root:123456@tcp(localhost:3333)/test?charset=utf8&parseTime=True&loc=Local")
	err := db.Ping()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	//p := Blog2{}
	// this will pull the first place directly into p
	//err = db.Get(&p, "SELECT id,title FROM blog LIMIT 1")

	//rows := db.QueryRowx(`SELECT 1`)
	//fmt.Printf("rows: %+v", rows)

	ret := make([]map[string]interface{}, 0)
	rows, err := db.Queryx("SELECT id,title,date,test FROM blog WHERE id=16 limit 10")
	for rows.Next() {
		results := make(map[string]interface{})
		rows.MapScan(results)
		for key, value := range results {
			if value == nil {
				results[key] = ""
			} else {
				results[key] = string(value.([]byte))
			}
		}
		ret = append(ret, results)
	}
	fmt.Printf("rows: %+v", ret)

	server := &http.Server{Addr: "127.0.0.1:8888"}
	http.HandleFunc("/favicon.ico", favicon)
	server.ListenAndServe()
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "statics/favicon.ico")
}
