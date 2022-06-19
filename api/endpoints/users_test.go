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
			name:         "list users",
			expectedVerb: "GET",
			expectedPath: "/users",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_users.json"),
			fn: func(t *testing.T, c rest.Client) {
				groups, err := NewUsers(c).List()
				require.NoError(t, err)

				expected := []*api.Users{
					{
						EmailAddress: "users.list@test.com",
						DisplayName:  "User Tester",
						UserID:       "09856666678",
						UserGroup: []string{
							"Admin",
							"SRE",
						},
					},
					{
						EmailAddress: "users.list2@test.com",
						DisplayName:  "User Tester2",
						UserID:       "34567890",
						UserGroup: []string{
							"SRE",
						},
					},
					{
						EmailAddress: "users.list3@test.com",
						DisplayName:  "User Tester3",
						UserID:       "34567111890",
						UserGroup: []string{
							"SRE",
						},
					},
				}

				assert.Equal(t, expected, groups)
			},
		},
	})
}
