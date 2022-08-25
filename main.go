package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Sut103/HCW-SS-Viewer/aws"
	"github.com/labstack/echo/v4"
)

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

func main() {
	e := echo.New()

	e.GET("/api/screenshots", getScreenshots)

	e.Logger.Fatal(e.Start(":1323"))
}
