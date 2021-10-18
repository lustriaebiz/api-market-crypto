package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "Halo bro"
		return ctx.String(http.StatusOK, data)
	})

	r.Start(":8000")
}
