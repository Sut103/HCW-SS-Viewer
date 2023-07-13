package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Sut103/HCW-SS-Viewer/aws"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

var echoLambda *echoadapter.EchoLambda

func init() {
	e := echo.New()
	e.GET("/", getScreenshots)

	echoLambda = echoadapter.New(e)
}

func getScreenshots(c echo.Context) error {
	screenshots, err := aws.Scan()
	if err != nil {
		log.Println(time.Now(), "getScreenshots(): ", err)
		return echo.ErrInternalServerError
	}

	jsonByte, err := json.Marshal(screenshots)
	if err != nil {
		log.Println(time.Now(), "getScreenshots(): ", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, string(jsonByte))
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(HandleRequest)
}
