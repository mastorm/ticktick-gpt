package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/mastorm/ticktick-gpt/api"
)

func main() {
	serveMux := api.ServeMux()
	adapter := httpadapter.New(serveMux).ProxyWithContext
	lambda.Start(adapter)
}
