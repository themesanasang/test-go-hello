package main

import (
	"os"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
	   port = "1323"
	   fmt.Println("No Port In Heroku" + port)
	}
	return ":" + port 
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(getPort()))
}
