package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func AddTraceId(next echo.HandlerFunc) echo.HandlerFunc {
	var traceHeader string = "trace-id"



	return func(c echo.Context) error {
		trace := c.Request().Header.Get(traceHeader)
		
		if trace == "" {
			trace = uuid.New().String()
			c.Request().Header.Set(traceHeader, trace)
		}

		c.Set(traceHeader, trace)
        c.Response().Header().Set(traceHeader, trace)


        if err := next(c); err != nil { //exec main process
            c.Error(err)
        }
        return nil
    }
}