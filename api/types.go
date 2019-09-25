package api

type Status int

const (
	Down           Status = 0
	Up             Status = 1
	Trouble        Status = 2
	Suspended      Status = 5
	Maintenance    Status = 7
	Discovery      Status = 9
	DiscoveryError Status = 10
)

type ValueAndSeverity struct {
	Value    string `json:"value"`
	Severity Status `json:"severity"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ActionRef struct {
	ActionID  string `json:"action_id"`
	AlertType Status `json:"alert_type"`
}

type Monitor struct {
	MonitorID             string            `json:"monitor_id,omitempty"`
	DisplayName           string            `json:"display_name"`
	Type                  string            `json:"type"`
	Website               string            `json:"website"`
	CheckFrequency        string            `json:"check_frequency"`
	HTTPMethod            string            `json:"http_method"`
	AuthUser              string            `json:"auth_user"`
	AuthPass              string            `json:"auth_pass"`
	MatchingKeyword       *ValueAndSeverity `json:"matching_keyword,omitempty"`
	UnmatchingKeyword     *ValueAndSeverity `json:"unmatching_keyword,omitempty"`
	MatchRegex            *ValueAndSeverity `json:"match_regex,omitempty"`
	MatchCase             bool              `json:"match_case"`
	UserAgent             string            `json:"user_agent"`
	CustomHeaders         []Header          `json:"custom_headers"`
	Timeout               int               `json:"timeout"`
	LocationProfileID     string            `json:"location_profile_id"`
	NotificationProfileID string            `json:"notification_profile_id"`
	ThresholdProfileID    string            `json:"threshold_profile_id"`
	MonitorGroups         []string          `json:"monitor_groups"`
	UserGroupIDs          []string          `json:"user_group_ids"`
	ActionIDs             []ActionRef       `json:"action_ids,omitempty"`
	UseNameServer         bool              `json:"use_name_server"`
}

// @TODO(mohmann): add necessary fields
type MonitorGroup struct{}

// @TODO(mohmann): add necessary fields
type NotificationProfile struct{}

// @TODO(mohmann): add necessary fields
type LocationProfile struct{}

// @TODO(mohmann): add necessary fields
type ThresholdProfile struct{}

// @TODO(mohmann): add necessary fields
type UserGroup struct{}
