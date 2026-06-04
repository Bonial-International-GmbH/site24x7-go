package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

// Users is the interface for managing Site24x7 user accounts.
type Users interface {
	Get(userID string) (*api.User, error)
	Create(user *api.User) (*api.User, error)
	Update(user *api.User) (*api.User, error)
	Delete(userID string) error
	List() ([]*api.User, error)
}

type users struct {
	client rest.Client
}

// NewUsers creates a new Users endpoint client.
func NewUsers(client rest.Client) Users {
	return &users{
		client: client,
	}
}

func (c *users) Get(userID string) (*api.User, error) {
	user := &api.User{}
	err := c.client.
		Get().
		Resource("users").
		ResourceID(userID).
		Do().
		Into(user)

	return user, err
}

func (c *users) Create(user *api.User) (*api.User, error) {
	newUser := &api.User{}
	err := c.client.
		Post().
		Resource("users").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(user).
		Do().
		Into(newUser)

	return newUser, err
}

func (c *users) Update(user *api.User) (*api.User, error) {
	updatedUser := &api.User{}
	err := c.client.
		Put().
		Resource("users").
		ResourceID(user.UserID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(user).
		Do().
		Into(updatedUser)

	return updatedUser, err
}

func (c *users) Delete(userID string) error {
	return c.client.
		Delete().
		Resource("users").
		ResourceID(userID).
		Do().
		Err()
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
