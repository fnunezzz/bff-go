package middleware

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func Logger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogStatus: true,
			LogURI:    true,
			BeforeNextFunc: func(c echo.Context) {
				c.Set("trace-id", uuid.New().String())
			},
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				value, _ := c.Get("trace-id").(string)
				fmt.Printf("REQUEST: uri: %v, status: %v, trace-id: %v\n", v.URI, v.Status, value)
				return nil
			},
		})
}



