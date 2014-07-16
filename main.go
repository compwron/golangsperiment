package main

import (
	"foo/config"
	"foo/uuid"
)

func main() {
	config.AppConfiguration()
	uuid.GenerateUuid()

	// goji.Post("/hello/:name", hello)
	// goji.Serve()
}
