package web

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
)

func HandleTestPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.ParseForm())
	fmt.Println(r.PostForm())
}
