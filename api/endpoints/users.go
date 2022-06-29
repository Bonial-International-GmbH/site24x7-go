package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type Users interface {
	List() ([]*api.User, error)
}

type users struct {
	client rest.Client
}

func NewUsers(client rest.Client) Users {
	return &users{
		client: client,
	}
}

func (c *users) List() ([]*api.User, error) {
	users := []*api.User{}
	err := c.client.
		Get().
		Resource("users").
		Do().
		Into(&users)

	return users, err
}
