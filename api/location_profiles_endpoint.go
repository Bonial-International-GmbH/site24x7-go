package api

type LocationProfilesEndpoint interface {
	Get(profileID string) (*LocationProfile, error)
	Create(profile *LocationProfile) (*LocationProfile, error)
	Update(profile *LocationProfile) (*LocationProfile, error)
	Delete(profileID string) error
	List() ([]*LocationProfile, error)
}

type locationProfilesEndpoint struct {
	client RequestFactory
}

// TODO(mohmann) implement LocationProfilesEndpoint interface here
