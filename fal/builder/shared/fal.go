package main

import "C"

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {}

func callWrapper(action func() string) *C.char {
	return C.CString(action())
}

func callWithParamWrapper(cval *C.char, action func(string) string) *C.char {
	// receive a callback function and perform the (C) char* -> (Go) string -> (C) char*

	// conversion cval is a pointer allocated by python and tracked by it's
	// gargabe collector it should NOT be free'd, otherwise python will abort
	return C.CString(action(C.GoString(cval)))
}

//export __FAL_stringFuncCall
func __FAL_stringFuncCall(cval *C.char) *C.char {
	return callWithParamWrapper(
		cval,
		func(val string) string {
			return fmt.Sprintf("Received: %s", val)
		},
	)
}

//export __FAL_listS3Buckets
func __FAL_listS3Buckets() *C.char {
	ctx := context.Background()

	cfg, _ := config.LoadDefaultConfig(ctx)
	client := s3.NewFromConfig(cfg)
	params := s3.ListBucketsInput{}

	return callWrapper(
		func() string {
			result, err := client.ListBuckets(ctx, &params)
			if err != nil {
				return fmt.Sprintf("%s", err)
			}

			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("Found %d buckets\n", len(result.Buckets)))

			for _, bucket := range result.Buckets {
				b.WriteString(fmt.Sprintln(*bucket.Name))
			}

			return b.String()
		},
	)
}
