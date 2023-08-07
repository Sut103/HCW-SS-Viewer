package main

import (
	"context"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/Sut103/HCW-SS-Viewer/aws"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

var echoLambda *echoadapter.EchoLambda

type Template struct {
	templates *template.Template
}

func init() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.GET("/", getScreenshots)

	echoLambda = echoadapter.New(e)
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func getScreenshots(c echo.Context) error {
	screenshots, err := aws.Scan()
	if err != nil {
		log.Println(time.Now(), "getScreenshots(): ", err)
		return echo.ErrInternalServerError
	}

	// 暫定で配列からランダムに6個とりだす
	ret_screenshots := make([]aws.Screenshot, 0)
	for i := 0; i < 6; i++ {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		ret_screenshots = append(ret_screenshots, screenshots[rand.Intn(len(screenshots))])
	}

	return c.Render(http.StatusOK, "main", ret_screenshots)
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(HandleRequest)
}
