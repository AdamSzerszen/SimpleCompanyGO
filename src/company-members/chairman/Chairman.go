package chairman

import (
	"../../servers"
	"../../tasks"
)

type IChairman interface {
	GetFullId() string
}

type Chairman struct {
	id          string
	taskServer  servers.ITaskServer
	taskFactory *tasks.ITaskFactory
}

func (chairman Chairman) GetFullId() string {
	return chairman.id
}
