package main

import (
	"foo/uuid"
	"foo/json"
)

func main() {
	uuid.GenerateUuid()
	
	json.ParseKnownFormatJsonFromFile()
	json.ParseArbitraryJson()
	json.ParseArbitraryJsonUsingLibrary()
	// json.WriteJsonToFile()


	// goji.Post("/hello/:name", hello)
	// goji.Serve()
}
