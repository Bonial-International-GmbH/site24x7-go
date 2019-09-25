package api

type UserGroupsEndpoint interface {
	Get(groupID string) (*UserGroup, error)
	Create(group *UserGroup) (*UserGroup, error)
	Update(group *UserGroup) (*UserGroup, error)
	Delete(groupID string) error
	List() ([]*UserGroup, error)
}

type userGroupsEndpoint struct {
	client RequestFactory
}

// TODO(mohmann) implement UserGroupsEndpoint interface here
