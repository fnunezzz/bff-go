package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func Gzip() echo.MiddlewareFunc { 
	return middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	  })
}