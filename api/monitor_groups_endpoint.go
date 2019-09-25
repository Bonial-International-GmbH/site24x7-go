package api

type MonitorGroupsEndpoint interface {
	Get(groupID string) (*MonitorGroup, error)
	Create(group *MonitorGroup) (*MonitorGroup, error)
	Update(group *MonitorGroup) (*MonitorGroup, error)
	Delete(groupID string) error
	List() ([]*MonitorGroup, error)
}

type monitorGroupsEndpoint struct {
	client RequestFactory
}

// TODO(mohmann) implement MonitorGroupsEndpoint interface here
