package s3

import (
	"errors"
	"fmt"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/s3"
)

func UploadSampleFile(bucketName string, pathInBucket string, fileContents []byte) {
	bucket := getBucket(bucketName)

	fmt.Println("\nAbout to upload to S3")
	uploadToS3(bucket, pathInBucket, fileContents)

	fmt.Println("\nAbout to download from S3")
	uploadedContents, _ := getFromS3(bucket, pathInBucket)
	fmt.Println("\nData downloaded from S3 bucket:", string(uploadedContents))
}

func authorizeToAws() aws.Auth {
	auth, err := aws.EnvAuth()
	if err != nil {
		fmt.Println("error in aws login", err)
	}
	return auth
}

func getBucket(bucketName string) (bucket *s3.Bucket) {
	auth := authorizeToAws()
	s := s3.New(auth, aws.USWest)
	return s.Bucket(bucketName)
}

func uploadToS3(bucket *s3.Bucket, pathInBucket string, fileContents []byte) {
	err := bucket.Put(pathInBucket, fileContents, "content-type", s3.Private, s3.Options{})
	if err != nil {
		fmt.Println("\n", err, "\n")
	}
}

func getFromS3(bucket *s3.Bucket, pathInBucket string) ([]byte, error) {
	data, err := bucket.Get(pathInBucket)
	if err != nil {
		fmt.Println("\n", err, "\n")
		return nil, errors.New("Couldn't Get data")
	}
	return data, err
}
