package main

import (
	"context"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Sut103/HCW-SS-Viewer/aws"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

var echoLambda *echoadapter.EchoLambda
var MAX_HIGHT int = 600

type Template struct {
	templates *template.Template
}

type Response struct {
	URL         string `json:"url"`
	OriginalURL string `json:"original_url"`
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
	res_screenshots := make([]Response, 0)
	for i := 0; i < 6; i++ {

		rand.New(rand.NewSource(time.Now().UnixNano()))
		random_num := rand.Intn(len(screenshots))
		url := screenshots[random_num].URL

		// Attachmentsが複数あるため、URLに対応するAttachmentを特定する必要がある
		for _, attachment := range screenshots[random_num].ChannelMessage.Attachments {
			if attachment.ProxyURL == url {
				width := strconv.Itoa((attachment.Width * MAX_HIGHT) / attachment.Height)
				height := strconv.Itoa(MAX_HIGHT)

				res_screenshots = append(res_screenshots, Response{
					URL:         url + "?width=" + width + "&height=" + height,
					OriginalURL: url,
				})

				break
			}
		}
	}

	return c.Render(http.StatusOK, "main", res_screenshots)
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(HandleRequest)
}
