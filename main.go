package main

import (
	"foo/config"
	"net/http"
	"foo/web"
)

func main() {
	config.AppConfiguration()
	http.HandleFunc("/test", web.HandleTestPost)
	http.ListenAndServe(":8082", nil)
}
