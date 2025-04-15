package aion

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Main() {
	e := echo.New()
	// ヘルスチェック
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
