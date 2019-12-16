package endpoints

import (
	"fmt"
	"testing"

	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCurrentStatus(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "list current status",
			expectedVerb: "GET",
			expectedPath: "/current_status",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_current_status.json"),
			fn: func(t *testing.T, c rest.Client) {
				status, err := NewCurrentStatus(c).List(nil)
				require.NoError(t, err)

				expected := &api.MonitorsStatus{
					Monitors: []*api.MonitorStatus{
						{
							Name:           "Server1",
							DownReason:     "maintenance",
							Duration:       "12 days 4 Hrs 55 Mins",
							Status:         api.Down,
							LastPolledTime: "2015-07-21T16:37:41+0530",
							ServerType:     "WINDOWS",
							MonitorType:    "SERVER",
							Unit:           "%",
							Tags:           []string{"down_tag"},
							MonitorID:      "355000001863001",
						},
						{
							Name:           "Site1",
							Status:         api.Trouble,
							LastPolledTime: "2015-07-21T15:30:35+0530",
							MonitorType:    "URL",
							Unit:           "ms",
							OutageID:       "1526624941082",
							DowntimeMillis: "12885615",
							DownReason:     "Response time from California - IN exceeded 2000 ms.",
							Duration:       "3 Hrs 35 Mins ",
							MonitorID:      "355000001863103",
						},
					},
				}

				assert.Equal(t, expected, status)
			},
		},
		{
			name:         "list current status with options",
			expectedVerb: "GET",
			expectedPath: "/current_status?apm_required=false&status_required=0%2C2&suspended_required=true",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_current_status.json"),
			fn: func(t *testing.T, c rest.Client) {
				options := &api.CurrentStatusListOptions{
					APMRequired:       api.Bool(false),
					SuspendedRequired: api.Bool(true),
					StatusRequired:    api.String(fmt.Sprintf("%d,%d", api.Down, api.Trouble)),
				}

				_, err := NewCurrentStatus(c).List(options)
				require.NoError(t, err)
			},
		},
	})
}
