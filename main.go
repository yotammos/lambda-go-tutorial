package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

var invokeCount = 0
var myObjects []*s3.Object

func init() {
	svc := s3.New(session.New())
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String("yotammosbucket"),
	}
	result, _ := svc.ListObjectsV2(input)
	myObjects = result.Contents
}

func LambdaHandler() (int, error) {
	invokeCount = invokeCount + 1
	log.Print(myObjects)
	return invokeCount, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
