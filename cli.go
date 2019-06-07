package dockerui

import (
	"github.com/labstack/echo"
	"github.com/urfave/cli"
)

var CliFlags = []cli.Flag{
	cli.StringFlag{Name: "listen", Value: ":9000", EnvVar: "LISTEN", Usage: "Listen address for HTTP access"},
}

func NewFlags(c *cli.Context) *Flags {
	return &Flags{
		Listen: c.String("listen"),
	}
}

type Flags struct {
	Listen string
}

func Action(c *cli.Context) error {

	flags := NewFlags(c)

	e := echo.New()

	return e.Start(flags.Listen)
}
