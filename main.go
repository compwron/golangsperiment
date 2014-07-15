package main

import (
	"foo/config"
	"github.com/go-martini/martini"
	"fmt"
)

func main() {
	config.AppConfiguration()

  m := martini.Classic()
  m.Post("/foo", func(params martini.Params) string {
        fmt.Println(params)
    return "Hello world!"
  })
  m.Run()
}
