package middleware

import (
	"os"

	"github.com/labstack/echo/middleware"
)

// setting application log config
func OutputAppLog(filepath string) middleware.LoggerConfig {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: f,
	}
}
