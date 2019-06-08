package dockerui

import (
	"github.com/docker/docker/client"
	"github.com/labstack/echo"
)

type HandleFunc func(*Context) error

type Context struct {
	echo.Context
}

func (c Context) GetDockerClient() *client.Client {
	return c.Get("dc").(*client.Client)
}

func UseContext(next HandleFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(&Context{c})
	}
}
