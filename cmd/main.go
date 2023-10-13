package main

import (
	adapter "Telegram/pkg/adapter"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

var handlerAdapter *httpadapter.HandlerAdapter

func Handler(ctx context.Context, event *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if handlerAdapter == nil {
		handlerAdapter = adapter.SetupAdapter()
	}
	response, err := handlerAdapter.ProxyWithContext(ctx, *event)
	return &response, err
}
