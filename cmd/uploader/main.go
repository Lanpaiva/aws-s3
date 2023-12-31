package main

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"SDFASDFASDFASDFAS",
				"sdfasdfasdfasdfasfasfasf",
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(sess)
	s3Bucket = "goBucket"

}

func main() {

	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Error reading directory")
				continue
			}
			uploadFile(files[0].Name())
		}
	}

}

func uploadFile(filename string) {
	completeFileName := fmt.Sprintf("./tmp/%s", filename)

	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error to opening file %s\n", completeFileName)
		return
	}
	defer f.Close()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("error to upload file %s\n", completeFileName)
		return
	}
	fmt.Printf("file %s uploaded sucessfully\n", completeFileName)
}
