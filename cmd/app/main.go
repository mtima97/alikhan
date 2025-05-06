package main

import (
	"alikhan2/internal"
	"github.com/labstack/echo/v4"
	"net/http"
)

type EmailRequest struct {
	Email string `json:"email"`
}

func main() {
	// БД
	db := internal.Init()

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	e.POST("/save-email", func(c echo.Context) error {
		req := new(EmailRequest)

		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, response("Ошибка "+err.Error()))
		}

		if err := db.Save(req.Email); err != nil {
			return c.JSON(http.StatusInternalServerError, response(err.Error()))
		}

		return c.JSON(http.StatusOK, response("Почта получена успешно"))
	})

	e.GET("/get-emails", func(c echo.Context) error {
		result, err := db.Get()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response(err.Error()))
		}

		return c.JSON(http.StatusOK, result)
	})

	if err := e.Start(":8009"); err != nil {
		e.Logger.Fatal()
		db.Close()
	}
}

func response(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}
