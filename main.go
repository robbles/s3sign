package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Args struct {
	Region string        `arg:"help:the AWS region"`
	Expiry time.Duration `arg:"help:how long the URL should stay valid for"`
	Bucket string        `arg:"positional,required,help:the S3 bucket name"`
	Key    string        `arg:"positional,required,help:the S3 key"`
}

func main() {
	args := Args{
		Region: "us-west-1",
		Expiry: time.Hour * 24,
	}
	arg.MustParse(&args)

	session := session.New(&aws.Config{
		Region: &args.Region,
	})
	client := s3.New(session)

	req, _ := client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &args.Bucket,
		Key:    &args.Key,
	})
	url, err := req.Presign(args.Expiry)

	if err != nil {
		log.Println("Failed to sign request", err)
		os.Exit(1)
	}

	fmt.Print(url)
}
