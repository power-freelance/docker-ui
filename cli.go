package dockerui

import (
	"log"

	"github.com/docker/docker/client"
	"github.com/labstack/echo"
	"github.com/urfave/cli"
)

var CliFlags = []cli.Flag{
	cli.StringFlag{Name: "listen", Value: ":9000", EnvVar: "LISTEN", Usage: "Listen address for HTTP access"},
	cli.StringFlag{Name: "staticRoot", Value: "web/dist", EnvVar: "STATIC_ROOT", Usage: "Static files root dir"},
}

func NewFlags(c *cli.Context) *Flags {
	return &Flags{
		Listen:     c.String("listen"),
		StaticRoot: c.String("staticRoot"),
	}
}

type Flags struct {
	Listen     string
	StaticRoot string
}

func Action(c *cli.Context) error {

	// Parse cli flags
	flags := NewFlags(c)

	// Connect to docker
	dc, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	defer func() {
		if err := dc.Close(); err != nil {
			log.Fatalf("error close docker client: %s", err)
		}
	}()

	// Init echo web-router
	e := echo.New()

	// Set docker client in every request context
	e.Use(HttpMiddlewareDockerClient(dc))

	// Web app serve
	e.Static("/", flags.StaticRoot)

	// API endpoints
	api := e.Group("/api")
	{
		api.GET("/", UseHttpContext(HttpIndex))
		api.GET("/image", UseHttpContext(HttpImageList))
		api.GET("/container", UseHttpContext(HttpContainerList))
	}

	return e.Start(flags.Listen)
}
