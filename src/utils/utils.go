package utils

import "github.com/labstack/echo/v4"

func RespondWithError(c echo.Context, statusCode int, message string, err error) error {
	return c.JSON(statusCode, map[string]string{
		"error":   message,
		"details": err.Error(),
	})
}
