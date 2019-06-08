package dockerui

import (
	"github.com/docker/docker/client"
	"github.com/labstack/echo"
)

func MiddlewareDockerClient(dc *client.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("dc", dc)
			return next(c)
		}
	}
}
