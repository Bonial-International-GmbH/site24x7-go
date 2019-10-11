package fake

import (
	"errors"
	"testing"

	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestClientMonitorsCreate(t *testing.T) {
	c := NewClient()

	monitor := &api.Monitor{
		DisplayName: "foo",
	}

	c.FakeMonitors.On("Create", mock.Anything).Return(monitor, nil).Once()
	c.FakeMonitors.On("Create", mock.Anything).Return(nil, errors.New("whoops")).Once()

	result, err := c.Monitors().Create(&api.Monitor{})

	require.NoError(t, err)

	assert.Equal(t, monitor, result)

	_, err = c.Monitors().Create(&api.Monitor{})

	require.Error(t, err)

	assert.Equal(t, errors.New("whoops"), err)

	c.FakeMonitors.AssertExpectations(t)
}
