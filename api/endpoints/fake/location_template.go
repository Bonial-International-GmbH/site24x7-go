package fake

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.LocationTemplate = &LocationTemplate{}

type LocationTemplate struct {
	mock.Mock
}

func (e *LocationTemplate) Get() (*api.LocationTemplate, error) {
	args := e.Called()
	if obj, ok := args.Get(0).(*api.LocationTemplate); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
