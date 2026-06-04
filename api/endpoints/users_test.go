package endpoints

import (
	"testing"

	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsers(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "create user",
			expectedVerb: "POST",
			expectedPath: "/users",
			expectedBody: fixture(t, "requests/create_user.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				user := &api.User{
					DisplayName:   "Jane Doe",
					EmailAddress:  "jane.doe@example.com",
					Role:          2,
					NotifyMedium:  []int{1, 3},
					SelectionType: 0,
					UserGroups:    []string{"123456789012", "987654321098"},
					AlertSettings: &api.AlertSettings{
						EmailFormat:    1,
						AlertStartTime: "00:00",
						AlertEndTime:   "00:00",
						Down:           true,
						Trouble:        true,
						Up:             true,
						AppLogs:        false,
					},
				}

				_, err := NewUsers(c).Create(user)
				require.NoError(t, err)
			},
		},
		{
			name:         "get user",
			expectedVerb: "GET",
			expectedPath: "/users/112233445566",
			statusCode:   200,
			responseBody: fixture(t, "responses/get_user.json"),
			fn: func(t *testing.T, c rest.Client) {
				user, err := NewUsers(c).Get("112233445566")
				require.NoError(t, err)

				expected := &api.User{
					UserID:        "112233445566",
					DisplayName:   "Jane Doe",
					EmailAddress:  "jane.doe@example.com",
					Role:          2,
					NotifyMedium:  []int{1, 3},
					SelectionType: 0,
					UserGroups:    []string{"123456789012", "987654321098"},
					AlertSettings: &api.AlertSettings{
						EmailFormat:    1,
						AlertStartTime: "00:00",
						AlertEndTime:   "00:00",
						Down:           true,
						Trouble:        true,
						Up:             true,
						AppLogs:        false,
					},
				}

				assert.Equal(t, expected, user)
			},
		},
		{
			name:         "list users",
			expectedVerb: "GET",
			expectedPath: "/users",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_users.json"),
			fn: func(t *testing.T, c rest.Client) {
				userList, err := NewUsers(c).List()
				require.NoError(t, err)

				expected := []*api.User{
					{
						UserID:        "112233445566",
						DisplayName:   "Jane Doe",
						EmailAddress:  "jane.doe@example.com",
						Role:          2,
						NotifyMedium:  []int{1, 3},
						SelectionType: 0,
						UserGroups:    []string{"123456789012", "987654321098"},
						AlertSettings: &api.AlertSettings{
							EmailFormat:    1,
							AlertStartTime: "00:00",
							AlertEndTime:   "00:00",
							Down:           true,
							Trouble:        true,
							Up:             true,
							AppLogs:        false,
						},
					},
					{
						UserID:        "223344556677",
						DisplayName:   "John Admin",
						EmailAddress:  "john.admin@example.com",
						Role:          1,
						NotifyMedium:  []int{1},
						SelectionType: 0,
					},
				}

				assert.Equal(t, expected, userList)
			},
		},
		{
			name:         "update user",
			expectedVerb: "PUT",
			expectedPath: "/users/112233445566",
			expectedBody: fixture(t, "requests/update_user.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				user := &api.User{
					UserID:        "112233445566",
					DisplayName:   "Jane Smith",
					EmailAddress:  "jane.smith@example.com",
					Role:          1,
					NotifyMedium:  []int{1},
					SelectionType: 0,
					UserGroups:    []string{"123456789012"},
					AlertSettings: &api.AlertSettings{
						EmailFormat:    1,
						AlertStartTime: "09:00",
						AlertEndTime:   "17:00",
						Down:           true,
						Trouble:        false,
						Up:             true,
						AppLogs:        false,
					},
				}

				_, err := NewUsers(c).Update(user)
				require.NoError(t, err)
			},
		},
		{
			name:         "delete user",
			expectedVerb: "DELETE",
			expectedPath: "/users/112233445566",
			statusCode:   200,
			fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewUsers(c).Delete("112233445566"))
			},
		},
	})
}
