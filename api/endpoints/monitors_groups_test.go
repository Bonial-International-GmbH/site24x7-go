package endpoints

import (
	"testing"

	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMonitorGroups(t *testing.T) {
	runEndpointTests(t, []*endpointTest{
		{
			name:         "create monitor group",
			expectedVerb: "POST",
			expectedPath: "/monitor_groups",
			expectedBody: fixture(t, "requests/create_monitor_group.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				group := &api.MonitorGroup{
					DisplayName:         "foo group",
					DependencyReourceID: "123",
					Description:         "This is foo group",
					Monitors: []string{
						"726000000002460",
						"726000000002464",
					},
					SuppressAlert:        true,
					HealthThresholdCount: 10,
				}

				_, err := NewMonitorGroupsEndpoint(c).Create(group)
				require.NoError(t, err)
			},
		},
		{
			name:         "get monitor group",
			expectedVerb: "GET",
			expectedPath: "/monitor_groups/113770000041271035",
			statusCode:   200,
			responseBody: fixture(t, "responses/get_monitor_group.json"),
			fn: func(t *testing.T, c rest.Client) {
				group, err := NewMonitorGroupsEndpoint(c).Get("113770000041271035")
				require.NoError(t, err)

				expected := &api.MonitorGroup{
					GroupID:     "113770000041271035",
					DisplayName: "Group1",
					Description: "Group all IDC monitors.",
					Monitors: []string{
						"726000000002460",
						"726000000002464",
					},
					DependencyReourceID:  "123",
					SuppressAlert:        true,
					HealthThresholdCount: 1,
				}

				assert.Equal(t, expected, group)
			},
		},
		{
			name:         "list monitor groups",
			expectedVerb: "GET",
			expectedPath: "/monitor_groups",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_monitor_groups.json"),
			fn: func(t *testing.T, c rest.Client) {
				groups, err := NewMonitorGroupsEndpoint(c).List()
				require.NoError(t, err)

				expected := []*api.MonitorGroup{
					{
						GroupID:     "797300000123437",
						DisplayName: "misc",
						Description: "checks for misc sites",
						Monitors: []string{
							"13370000004999063",
							"79730133704999073",
							"79730000001337083",
							"12340000005000031",
						},
						HealthThresholdCount: 1,
					},
					{
						GroupID:              "79123400003075053",
						DisplayName:          "api",
						HealthThresholdCount: 1,
					},
					{
						GroupID:              "79730456703075223",
						DisplayName:          "web",
						HealthThresholdCount: 1,
					},
				}

				assert.Equal(t, expected, groups)
			},
		},
		{
			name:         "update monitor group",
			expectedVerb: "PUT",
			expectedPath: "/monitor_groups/123",
			expectedBody: []byte(`{"group_id":"123","display_name":"foo"}`),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				group := &api.MonitorGroup{
					GroupID:     "123",
					DisplayName: "foo",
				}

				_, err := NewMonitorGroupsEndpoint(c).Update(group)
				require.NoError(t, err)
			},
		},
		{
			name:         "delete monitor group",
			expectedVerb: "DELETE",
			expectedPath: "/monitor_groups/123",
			statusCode:   200,
			fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewMonitorGroupsEndpoint(c).Delete("123"))
			},
		},
	})
}
