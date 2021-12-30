package main

import (
	"C"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//export __addAndMultiplies
func __addAndMultiplies(value int) int {
	return (value + 1) * value
}

//export __listS3Buckets
func __listS3Buckets() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)

	client := s3.NewFromConfig(cfg)
	params := s3.ListBucketsInput{}

	result, err := client.ListBuckets(ctx, &params)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %d buckets\n", len(result.Buckets))
	for _, bucket := range result.Buckets {
		fmt.Println(*bucket.Name)
	}
}

func main() {}
