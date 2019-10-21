package main

import (
	"Demo/if/models"
	_ "Demo/if/routers"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func main() {
	blog := models.NewBlog()
	blog.GetBlogs()
	fmt.Printf("blogs: %+v", models.GetBlogs())

	server := &http.Server{Addr: "127.0.0.1:8888"}
	http.HandleFunc("/favicon.ico", favicon)
	server.ListenAndServe()
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "statics/favicon.ico")
}
