package endpoints

import "github.com/Bonial-International-GmbH/site24x7-go/rest"

// crud is a generic helper that implements the standard five REST operations
// (get, create, update, delete, list) for a given resource type T. Endpoint
// structs embed this type and delegate their exported interface methods to
// these unexported helpers.
type crud[T any] struct {
	client   rest.Client
	resource string
}

func (c *crud[T]) get(id string) (*T, error) {
	out := new(T)
	return out, c.client.Get().Resource(c.resource).ResourceID(id).Do().Into(out)
}

func (c *crud[T]) create(v *T) (*T, error) {
	out := new(T)
	return out, c.client.Post().Resource(c.resource).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(v).Do().Into(out)
}

func (c *crud[T]) update(id string, v *T) (*T, error) {
	out := new(T)
	return out, c.client.Put().Resource(c.resource).ResourceID(id).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(v).Do().Into(out)
}

func (c *crud[T]) delete(id string) error {
	return c.client.Delete().Resource(c.resource).ResourceID(id).Do().Err()
}

func (c *crud[T]) list() ([]*T, error) {
	var out []*T
	return out, c.client.Get().Resource(c.resource).Do().Into(&out)
}
