package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

type AWSService struct {
	s3Client   *s3.Client
	bucketName string
}

func New(bucketName string) *AWSService {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("err loading config")
	}
	return &AWSService{
		s3Client:   s3.NewFromConfig(cfg),
		bucketName: bucketName,
	}
}

func (awsSvc AWSService) UploadFile(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("close", err)
		}

		if err = os.Remove(file.Name()); err != nil {
			fmt.Println("Remove:", err)
		}

	}()

	_, err = awsSvc.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(awsSvc.bucketName),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		return err
	}

	return nil
}
