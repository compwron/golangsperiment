package main

import (
	"foo/json"
	"foo/uuid"
)

func main() {
	uuid.GenerateUuid()

	json.ParseKnownFormatJsonFromFile()
	json.ParseArbitraryJson()
	json.ParseArbitraryJsonUsingLibrary()
	json.ParseArbitraryJsonFromFileUsingLibrary()

	json.WriteJsonToFile()
}
