// +build integration

package location

import (
	"errors"
	"testing"

	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDNSIPSource(t *testing.T) {
	tests := []struct {
		name        string
		location    *api.Location
		expectedErr error
	}{
		{
			name: "lookup one word",
			location: &api.Location{
				LocationID:  "66",
				DisplayName: "Helsinki - FI",
				CityName:    "Helsinki",
			},
		},
		{
			name: "lookup multiword",
			location: &api.Location{
				LocationID:  "15",
				DisplayName: "Tel Aviv - IL",
				CityName:    "Tel Aviv",
			},
		},
		{
			name: "lookup multiword",
			location: &api.Location{
				LocationID:  "15",
				DisplayName: "Tel Aviv - IL",
				CityName:    "Tel Aviv",
			},
		},
		{
			name: "lookup multiword edge case",
			location: &api.Location{
				LocationID:  "16",
				DisplayName: "Rio de Janeiro - BR",
				CityName:    "Rio de Janeiro",
			},
		},
		{
			name: "lookup invalid multiword",
			location: &api.Location{
				LocationID:  "42",
				DisplayName: "Foo Bar Baz - DE",
				CityName:    "Foo Bar Baz",
			},
			expectedErr: errors.New(`failed to lookup IPs for location "Foo Bar Baz - DE": lookup foo-de.enduserexp.com: no such host`),
		},
		{
			name: "missing country code",
			location: &api.Location{
				LocationID:  "1",
				DisplayName: "Frankfurt",
				CityName:    "Frankfurt",
			},
			expectedErr: errors.New(`failed to parse country code from location name "Frankfurt"`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewDefaultDNSIPSource()

			ips, err := s.LookupIPs(test.location)
			if test.expectedErr != nil {
				require.Error(t, err)
				assert.Equal(t, test.expectedErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
				assert.NotEmpty(t, ips)
			}
		})
	}
}
