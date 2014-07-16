package main

import (
	"foo/json"
	"foo/s3"
	"foo/uuid"
)

func main() {
	uuid.GenerateUuid()

	json.ParseKnownFormatJsonFromFile()
	json.ParseArbitraryJson()
	json.ParseArbitraryJsonUsingLibrary()
	json.ParseArbitraryJsonFromFileUsingLibrary()
	json.WriteJsonToFile()

	bucketName := "someNewBuckete88c2f9e-b081-4f72-8559-13b9e8001d48" // existing bucket created via fmt.Sprintf("someNewBucket%s", uuid.GenerateUuid())
	s3.UploadSampleFile(bucketName, "somewhereInS3Bucket", []byte("some data to put in the file"))
}
