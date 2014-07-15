package web

import (
	"fmt"
	"net/http"
	"github.com/go-martini/martini"
)

func HandleTestPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.ParseForm())
	fmt.Println(r.PostForm())
}


