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
	crud[api.User]
}

// NewUsers creates a new Users endpoint client.
func NewUsers(client rest.Client) Users {
	return &users{crud[api.User]{client: client, resource: "users"}}
}

func (c *users) Get(userID string) (*api.User, error)     { return c.get(userID) }
func (c *users) Create(user *api.User) (*api.User, error) { return c.create(user) }
func (c *users) Update(user *api.User) (*api.User, error) { return c.update(user.UserID, user) }
func (c *users) Delete(userID string) error               { return c.delete(userID) }
func (c *users) List() ([]*api.User, error)               { return c.list() }
