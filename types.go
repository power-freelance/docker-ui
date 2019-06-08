package dockerui

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

type ContainerListParams struct {
	All    bool     `query:"all"`
	Status []string `query:"status[]"`
}

func (p ContainerListParams) GetOpts() types.ContainerListOptions {
	f := filters.NewArgs()

	if len(p.Status) != 0 {
		for _, status := range p.Status {
			f.Add("status", status)
		}
	}

	return types.ContainerListOptions{
		All:     p.All,
		Filters: f,
	}
}

type ComposeProject struct {
	ConfigHash string
	Name       string
	Version    string
	Services   map[string]ComposeService
}

type ComposeService struct {
	Name       string
	Containers map[int]types.Container
}
