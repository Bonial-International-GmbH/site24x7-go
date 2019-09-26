package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type MonitorGroupsEndpoint interface {
	Get(groupID string) (*api.MonitorGroup, error)
	Create(group *api.MonitorGroup) (*api.MonitorGroup, error)
	Update(group *api.MonitorGroup) (*api.MonitorGroup, error)
	Delete(groupID string) error
	List() ([]*api.MonitorGroup, error)
}

type monitorGroupsEndpoint struct {
	client rest.Client
}

// TODO(mohmann) implement MonitorGroupsEndpoint interface here
