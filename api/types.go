package api

import (
	"encoding/json"
	"strconv"
)

type Status int

// Custom unmarshaller that allows the value to be both a
// string and an integer, always unmarshals into integer
//
// The Site24x7 API has a bug where it accepts integers, but
// returns them as strings.
func (status *Status) UnmarshalJSON(rawValue []byte) error {
	if rawValue[0] != '"' {
		return json.Unmarshal(rawValue, (*int)(status))
	}

	var valueAsString string
	if err := json.Unmarshal(rawValue, &valueAsString); err != nil {
		return err
	}

	valueAsInt, err := strconv.Atoi(valueAsString)
	if err != nil {
		return err
	}

	*status = Status(valueAsInt)
	return nil
}

const (
	Down               Status = 0
	Up                 Status = 1
	Trouble            Status = 2
	Critical           Status = 3
	Suspended          Status = 5
	Maintenance        Status = 7
	Discovery          Status = 9
	ConfigurationError Status = 10
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
	CustomHeaders         []Header          `json:"custom_headers,omitempty"`
	Timeout               int               `json:"timeout"`
	LocationProfileID     string            `json:"location_profile_id"`
	NotificationProfileID string            `json:"notification_profile_id"`
	ThresholdProfileID    string            `json:"threshold_profile_id"`
	MonitorGroups         []string          `json:"monitor_groups,omitempty"`
	UserGroupIDs          []string          `json:"user_group_ids,omitempty"`
	ActionIDs             []ActionRef       `json:"action_ids,omitempty"`
	UseNameServer         bool              `json:"use_name_server"`
	UpStatusCodes         string            `json:"up_status_codes"`
}

// MonitorGroup organizes Monitor resources into groups.
type MonitorGroup struct {
	GroupID              string   `json:"group_id,omitempty"`
	DisplayName          string   `json:"display_name"`
	Description          string   `json:"description,omitempty"`
	Monitors             []string `json:"monitors,omitempty"`
	HealthThresholdCount int      `json:"health_threshold_count,omitempty"`
	DependencyReourceID  string   `json:"dependency_resource_id,omitempty"`
	SuppressAlert        bool     `json:"suppress_alert"`
}

// NotificationProfile allows tweaking when alerts have to be sent out.
type NotificationProfile struct {
	ProfileID                   string   `json:"profile_id,omitempty"`
	ProfileName                 string   `json:"profile_name"`
	RcaNeeded                   bool     `json:"rca_needed"`
	NotifyAfterExecutingActions bool     `json:"notify_after_executing_actions"`
	DowntimeNotificationDelay   int      `json:"downtime_notification_delay,omitempty"`
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

// LocationTemplate holds locations Site24x7 performs their monitor checks
// from.
type LocationTemplate struct {
	Locations []*Location `json:"locations"`
}

// Location is a physical location Site24x7 performs monitor checks from. The
// LocationID field maps to the IDs used in the PrimaryLocation and
// SecondaryLocations fields of LocationProfile values.
type Location struct {
	LocationID  string `json:"location_id"`
	CountryName string `json:"country_name"`
	DisplayName string `json:"display_name"`
	UseIPV6     bool   `json:"use_ipv6"`
	CityName    string `json:"city_name"`
	CityShort   string `json:"city_short"`
	Continent   string `json:"continent"`
}

// ThresholdProfile help the alarms engine to decide if a specific resource has to be declared critical or down
type ThresholdProfile struct {
	ProfileID              string `json:"profile_id,omitempty"`
	Type                   string `json:"type"`
	ProfileName            string `json:"profile_name"`
	DownLocationThreshold  int    `json:"down_location_threshold"`
	WebsiteContentModified bool   `json:"website_content_modified"`
}

// UserGroup help organize individuals so that they receive alerts and reports based on their responsibility.
type UserGroup struct {
	UserGroupID      string   `json:"user_group_id,omitempty"`
	DisplayName      string   `json:"display_name"`
	Users            []string `json:"users,omitempty"`
	AttributeGroupID string   `json:"attribute_group_id,omitempty"`
}

type Users struct {
	EmailAddress string   `json:"email_address,omitempty"`
	DisplayName  string   `json:"display_name"`
	UserID       string   `json:"user_id,omitempty"`
	UserGroup    []string `json:"user_groups,omitempty"`
}

// ITAutomation prioritize and remediate routine actions automatically,
// increase IT efficiency and streamline your processes to reduce performance degrade
type ITAutomation struct {
	ActionID               string `json:"action_id,omitempty"`
	ActionName             string `json:"action_name"`
	ActionUrl              string `json:"action_url"`
	ActionTimeout          int    `json:"action_timeout"`
	ActionType             int    `json:"action_type"`
	ActionMethod           string `json:"action_method"`
	SuppressAlert          bool   `json:"suppress_alert,omitempty"`
	SendIncidentParameters bool   `json:"send_incident_parameters"`
	SendCustomParameters   bool   `json:"send_custom_parameters"`
	CustomParameters       string `json:"custom_parameters"`
	SendInJsonFormat       bool   `json:"send_in_json_format"`
	AuthMethod             string `json:"auth_method,omitempty"`
	Username               string `json:"username,omitempty"`
	Password               string `json:"password,omitempty"`
	OAuth2Provider         string `json:"oauth2_provider,omitempty"`
	UserAgent              string `json:"user_agent,omitempty"`
}

// MonitorsStatus describes the response for the current status endpoint as
// defined here: https://www.site24x7.com/help/api/#retrieve-current-status.
type MonitorsStatus struct {
	Monitors []*MonitorStatus `json:"monitors"`
}

// MonitorStatus describes a monitor status response as defined here:
// https://www.site24x7.com/help/api/#retrieve-current-status.
type MonitorStatus struct {
	Name           string   `json:"name"`
	MonitorID      string   `json:"monitor_id"`
	MonitorType    string   `json:"monitor_type"`
	Status         Status   `json:"status"`
	LastPolledTime string   `json:"last_polled_time"`
	Unit           string   `json:"unit"`
	OutageID       string   `json:"outage_id"`
	DowntimeMillis string   `json:"downtime_millis"`
	DownReason     string   `json:"down_reason"`
	Duration       string   `json:"duration"`
	ServerType     string   `json:"server_type"`
	Tags           []string `json:"tags"`
}

// CurrentStatusListOptions hold the options that can be specified to filter
// current monitor statuses.
type CurrentStatusListOptions struct {
	APMRequired       *bool   `url:"apm_required,omitempty"`
	GroupRequired     *bool   `url:"group_required,omitempty"`
	SuspendedRequired *bool   `url:"suspended_required,omitempty"`
	LocationsRequired *bool   `url:"locations_required,omitempty"`
	StatusRequired    *string `url:"status_required,omitempty"`
}
