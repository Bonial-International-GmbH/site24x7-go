package api

type NotificationProfilesEndpoint interface {
	Get(profileID string) (*NotificationProfile, error)
	Create(profile *NotificationProfile) (*NotificationProfile, error)
	Update(profile *NotificationProfile) (*NotificationProfile, error)
	Delete(profileID string) error
	List() ([]*NotificationProfile, error)
}

type notificationProfilesEndpoint struct {
	client RequestFactory
}

// TODO(mohmann) implement NotificationProfilesEndpoint interface here
