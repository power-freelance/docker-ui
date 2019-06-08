package dockerui

import "github.com/docker/docker/api/types"

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
