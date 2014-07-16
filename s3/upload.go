package s3

import (
	"errors"
	"fmt"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/s3"
	"log"
	"foo/uuid"
)

func UploadSampleFile() {
	s3Response, err := uploadToS3()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nResponse to s3 upload: ", s3Response)
}

func uploadToS3() (string, error) {
	auth, err := aws.EnvAuth()
	fmt.Println("auth response", auth)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Couldn't auth to Amazon AWS")
	}

	s := s3.New(auth, aws.USWest)
	fmt.Println("new s3 response", s)

	bucketName := fmt.Sprintf("someNewBucket%s", uuid.GenerateUuid())
	fmt.Println("Bucket name: ", bucketName)
	
	bucket := s.Bucket(bucketName)
	fmt.Println(bucket, "")

	// someData := []byte(`{"foo":"bar","baz":6,"stuff":["a","b"], "isTrue":false}`)
	err = bucket.PutBucket(s3.Private)

	// name string, data []byte, datatype string, permission s3.ACL) (string, error) {
	// err = bucket.Put("somenewname", someData, "txt", s3.PublicRead, s3.Options{})
	if err != nil {
		fmt.Println("\n", err, "\n")
		return "", errors.New("Couldn't Upload That!")
	}
	return "it worked", nil
	// return "foo", nil
}
