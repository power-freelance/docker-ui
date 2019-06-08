package dockerui

import (
	"net/http"

	"github.com/docker/docker/api/types"

	"github.com/labstack/echo"
)

func EndpointIndex(c *Context) error {
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

func EndpointImageList(hc *Context) error {
	ctx := hc.Request().Context()
	dc := hc.GetDockerClient()
	images, err := dc.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return err
	}

	return hc.JSON(http.StatusOK, images)
}

func EndpointContainerList(hc *Context) error {
	// Parse request params
	params := new(ContainerListParams)
	if err := hc.Bind(params); err != nil {
		return err
	}

	ctx := hc.Request().Context()
	dc := hc.GetDockerClient()
	containers, err := dc.ContainerList(ctx, params.GetOpts())
	if err != nil {
		return err
	}

	return hc.JSON(http.StatusOK, containers)
}

func EndpointComposeProjectList(hc *Context) error {
	ctx := hc.Request().Context()
	dc := hc.GetDockerClient()
	containers, err := dc.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}

	// TODO: parse container labels and create docker-compose projects list

	return hc.JSON(http.StatusOK, containers)
}
