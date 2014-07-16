package s3

import (
	"fmt"
	"os"
)

type S3Configuration struct {
	AccessKeyId     string
	SecretAccessKey string
}

func importKeysFromEnvironment() S3Configuration {
	// os.Setenv("AccessKeyId", "1")
	// os.Setenv("SecretAccessKey", "2")
	c := S3Configuration{os.Getenv("AccessKeyId"), os.Getenv("SecretAccessKey")}
	return c
}

func UploadSampleFile() {
	c := importKeysFromEnvironment()
	fmt.Println("\nConfig from environment: ", c)
}
