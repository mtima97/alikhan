package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type EmailRequest struct {
	Email string `json:"email"`
}

func main() {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	e.POST("/save-email", func(c echo.Context) error {
		req := new(EmailRequest)

		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Ошибка " + err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Почта получена успешно",
		})
	})

	e.Logger.Fatal(e.Start(":8009"))
}
