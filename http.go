package dockerui

import (
	"net/http"

	"github.com/docker/docker/api/types"

	"github.com/docker/docker/client"
	"github.com/labstack/echo"
)

type HttpContext struct {
	echo.Context
}

func (hc HttpContext) GetDockerClient() *client.Client {
	return hc.Get("dc").(*client.Client)
}

func UseHttpContext(next func(*HttpContext) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		hc := &HttpContext{c}
		return next(hc)
	}
}

func HttpMiddlewareDockerClient(dc *client.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("dc", dc)
			return next(c)
		}
	}
}

func HttpIndex(c *HttpContext) error {
	ctx := c.Request().Context()
	dc := c.GetDockerClient()
	info, err := dc.Info(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"docker": map[string]interface{}{
			"version": dc.ClientVersion(),
			"info":    info,
		},
		"echo": map[string]interface{}{
			"version": echo.Version,
		},
	})
}

func HttpImageList(hc *HttpContext) error {
	ctx := hc.Request().Context()
	dc := hc.GetDockerClient()
	images, err := dc.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return err
	}

	return hc.JSON(http.StatusOK, images)
}

func HttpContainerList(hc *HttpContext) error {
	ctx := hc.Request().Context()
	dc := hc.GetDockerClient()
	containers, err := dc.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}

	return hc.JSON(http.StatusOK, containers)
}
