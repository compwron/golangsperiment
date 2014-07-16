package main

import (
	"foo/json"
	"foo/uuid"
	"foo/s3"
)

func main() {
	uuid.GenerateUuid()

	json.ParseKnownFormatJsonFromFile()
	json.ParseArbitraryJson()
	json.ParseArbitraryJsonUsingLibrary()
	json.ParseArbitraryJsonFromFileUsingLibrary()
	json.WriteJsonToFile()

	s3.UploadSampleFile()
}
