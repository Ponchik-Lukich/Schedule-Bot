package adapter

import (
	"Telegram/pkg/handlers/test"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gin-gonic/gin"
)

func SetupAdapter() *httpadapter.HandlerAdapter {
	r := gin.Default()
	r.GET("/", test.Test)
	return httpadapter.New(r)
}
