package web

import (
	"fmt"
	"net/http"
)

func HandleTestPost(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Form)
}
