package middlewares

import (
	"bitbucket.org/BBeamnantapong/cooking-server/core"
	"github.com/labstack/echo/v4"
)

func Core(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &core.Context{c}
		return h(cc)
	}
}
