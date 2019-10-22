package main

import (
	"Demo/if/models"
	_ "Demo/if/routers"
	"fmt"
	"net/http"
)

func main() {
	blog := models.NewBlog()
	blog.GetBlogs()
	fmt.Printf("blogs: %+v\n", blog)
	blog.Save(blog, map[string]interface{}{"id": 16, "title": "xxx"})
	fmt.Printf("blogs: %+v\n", blog)
	//title := blog.Get(blog, "Title")
	//fmt.Printf("blog title: %+v\n", title)

	server := &http.Server{Addr: "127.0.0.1:8888"}
	http.HandleFunc("/favicon.ico", favicon)
	server.ListenAndServe()
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "statics/favicon.ico")
}
