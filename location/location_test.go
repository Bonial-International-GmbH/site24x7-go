package location

import (
	"testing"

	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProfileIPProvider(t *testing.T) {
	tests := []struct {
		name        string
		locations   []*api.Location
		locationIPs map[string][]string
		profile     *api.LocationProfile
		expected    []string
		expectedErr error
	}{
		{
			name: "looks up IPs for primary location",
			locations: []*api.Location{
				{LocationID: "123"},
			},
			locationIPs: map[string][]string{
				"123": []string{"1.2.3.4", "5.6.7.8"},
			},
			profile: &api.LocationProfile{
				PrimaryLocation: "123",
			},
			expected: []string{"1.2.3.4", "5.6.7.8"},
		},
		{
			name: "looks up IPs for primary and secondary locations",
			locations: []*api.Location{
				{LocationID: "123"},
				{LocationID: "456"},
				{LocationID: "789"},
			},
			locationIPs: map[string][]string{
				"123": []string{"127.0.0.1"},
				"456": []string{"1.2.3.4", "5.6.7.8"},
				"789": []string{"1.1.1.1"},
			},
			profile: &api.LocationProfile{
				PrimaryLocation:    "456",
				SecondaryLocations: []string{"123", "789"},
			},
			expected: []string{"1.2.3.4", "5.6.7.8", "127.0.0.1", "1.1.1.1"},
		},
		{
			name: "unknown locations produce empty ip list",
			profile: &api.LocationProfile{
				PrimaryLocation: "123",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := &ProfileIPProvider{
				Locations: test.locations,
				IPSource: &StaticIPSource{
					LocationIPs: test.locationIPs,
				},
			}

			ips, err := p.GetLocationIPs(test.profile)
			if test.expectedErr != nil {
				require.Error(t, err)
				assert.Equal(t, test.expectedErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, ips)
			}
		})
	}
}
