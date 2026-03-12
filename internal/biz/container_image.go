package biz

import (
	"github.com/acepanel/panel/v3/internal/http/request"
	"github.com/acepanel/panel/v3/pkg/types"
)

type ContainerImageRepo interface {
	List() ([]types.ContainerImage, error)
	Exist(name string) (bool, error)
	Pull(req *request.ContainerImagePull) error
	Remove(id string) error
	Prune() error
}
