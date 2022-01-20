package main

import "C"

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

type _falBaseMessage struct {
	Message string `json:"_fal"`
}

type _falRequest struct {
	Input string `json:"input"`
}

type _falResponse struct {
	Success bool   `json:"success"`
	Output  string `json:"output"`
}

type _functionCall struct {
	ctx        *context.Context
	client     *lambda.Client
	arn        string
	msg        *string
	isAsync    bool
	collectLog bool
}

type _functionCallResponse struct {
	statusCode int32
	err        *string
	logs       *string
	output     *string
}

func performCall(call *_functionCall) (*_functionCallResponse, error) {
	/* ... */

	input := lambda.InvokeInput{
		FunctionName:   &call.arn,
		InvocationType: types.InvocationTypeRequestResponse,
		LogType:        types.LogTypeNone,
	}

	if call.isAsync {
		input.InvocationType = types.InvocationTypeEvent
	}

	if call.collectLog {
		input.LogType = types.LogTypeTail
	}

	if call.msg != nil {
		payload, err := json.Marshal(_falBaseMessage{Message: *call.msg})
		if err != nil {
			return nil, err
		}

		input.Payload = payload
	}

	invokeOutput, err := call.client.Invoke(*call.ctx, &input)
	if err != nil {
		return nil, err
	}

	/* ... */

	callResponse := _functionCallResponse{
		statusCode: invokeOutput.StatusCode,
		err:        invokeOutput.FunctionError,
	}

	if !call.isAsync {
		var response _falBaseMessage
		err = json.Unmarshal(invokeOutput.Payload, &response)
		if err != nil {
			return nil, err
		}

		callResponse.output = &response.Message
	}

	if call.collectLog && invokeOutput.LogResult != nil {
		logs, err := base64.StdEncoding.DecodeString(*invokeOutput.LogResult)
		if err != nil {
			panic(err)
		}

		slogs := string(logs)
		callResponse.logs = &slogs
	}

	return &callResponse, nil
}

func main() {
	ctx := context.Background()
	cfg, _ := config.LoadDefaultConfig(ctx)
	client := lambda.NewFromConfig(cfg)

	falRequest, err := json.Marshal(_falRequest{Input: "8"})
	if err != nil {
		panic(err)
	}

	msg := string(falRequest)
	response, err := performCall(&_functionCall{
		ctx:        &ctx,
		client:     client,
		arn:        "arn:aws:lambda:us-east-1:503265333675:function:plus-one-test-handler",
		msg:        &msg,
		isAsync:    false,
		collectLog: true,
	})

	/* ... */

	fmt.Printf("\nSTATUS CODE: \n%d\n", response.statusCode)

	if response.err != nil {
		fmt.Printf("\nERRORS: \n%s\n", *response.err)
	}

	if response.logs != nil {
		fmt.Printf("\nLOGS: \n%s\n", *response.logs)
	}

	if response.output != nil {
		fmt.Printf("\nOUTPUTS: \n%s\n", *response.output)
	}
}

/* ....... */

func callWrapper(action func() string) *C.char {
	return C.CString(action())
}

func callWithParamWrapper(cval *C.char, action func(string) string) *C.char {
	/* receive a callback function and perform the (C) char* -> (Go) string -> (C) char*
	 *
	 * conversion cval is a pointer allocated by python and tracked by it's
	 * gargabe collector it should NOT be free'd, otherwise python will abort
	 */
	return C.CString(action(C.GoString(cval)))
}

/* ....... */

//export __FAL_stringFuncCall
func __FAL_stringFuncCall(cval *C.char) *C.char {
	return callWithParamWrapper(
		cval,
		func(val string) string {
			return fmt.Sprintf("Received: %s", val)
		},
	)
}
