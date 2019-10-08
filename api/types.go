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

// Monitor performance of websites and internet services
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

// MonitorGroup organizes Monitor resources into groups.
type MonitorGroup struct {
	GroupID              string   `json:"group_id,omitempty"`
	DisplayName          string   `json:"display_name"`
	Description          string   `json:"description,omitempty"`
	Monitors             []string `json:"monitors,omitempty"`
	HealthThresholdCount int      `json:"health_threshold_count,omitempty"`
	DependencyReourceID  string   `json:"dependency_resource_id,omitempty"`
	SuppressAlert        bool     `json:"suppress_alert,omitempty"`
}

// NotificationProfile allows tweaking when alerts have to be sent out.
type NotificationProfile struct {
	ProfileID                   string   `json:"profile_id,omitempty"`
	ProfileName                 string   `json:"profile_name"`
	RcaNeeded                   bool     `json:"rca_needed"`
	NotifyAfterExecutingActions bool     `json:"notify_after_executing_actions"`
	DowntimeNotificationDelay   []int    `json:"downtime_notification_delay,omitempty"`
	PersistentNotification      int      `json:"persistent_notification,omitempty"`
	EscalationUserGroupId       string   `json:"escalation_user_group_id,omitempty"`
	EscalationWaitTime          int      `json:"escalation_wait_time"`
	EscalationAutomations       []string `json:"escalation_automations,omitempty"`
	EscalationServices          []string `json:"escalation_services,omitempty"`
	TemplateID                  string   `json:"template_id,omitempty"`
}

// LocationProfiles make it convenient to set monitoring locations consistently across many websites or monitors
type LocationProfile struct {
	ProfileID          string   `json:"profile_id,omitempty"`
	ProfileName        string   `json:"profile_name"`
	PrimaryLocation    string   `json:"primary_location"`
	SecondaryLocations []string `json:"secondary_locations,omitempty"`
	RestrictAltLoc     bool     `json:"restrict_alt_loc"`
}

// ThresholdProfile help the alarms engine to decide if a specific resource has to be declared critical or down
type ThresholdProfile struct {
	ProfileID              string `json:"profile_id"`
	Type                   string `json:"type"`
	ProfileName            string `json:"profile_name"`
	DownLocationThreshold  int    `json:"down_location_threshold"`
	WebsiteContentModified bool   `json:"website_content_modified,omitempty"`
}

// UserGroup help organize individuals so that they receive alerts and reports based on their responsibility.
type UserGroup struct {
	UserGroupID      string   `json:"user_group_id,omitempty"`
	DisplayName      string   `json:"display_name"`
	Users            []string `json:"users"`
	AttributeGroupID string   `json:"attribute_group_id,omitempty"`
}
