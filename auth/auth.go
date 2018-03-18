package middleware

import (
	response "github.com/anraku/echo-sample/response"
	"net/http"

	"github.com/labstack/echo"
)

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
			return next(c)
		}
	}
}

// auth Access Token in Request Header
func AuthToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("auth-token")
			if token != "abc" {
				return c.JSON(http.StatusUnauthorized, response.CreateResponse(
					map[string]interface{}{
						"message": "error response",
					},
				))
			}
			return next(c)
		}
	}
}
