package main

import (
        "fmt"
        "net/http"

        "github.com/zenazn/goji"
        "github.com/zenazn/goji/web"
)
var c  = make(chan int )

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	c <- 42
        fmt.Fprintf(w, 200)
}

func main() {
        goji.Post("/hello/:name", hello)
        goji.Serve()
}
