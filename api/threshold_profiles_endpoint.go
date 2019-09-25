package api

type ThresholdProfilesEndpoint interface {
	Get(profileID string) (*ThresholdProfile, error)
	Create(profile *ThresholdProfile) (*ThresholdProfile, error)
	Update(profile *ThresholdProfile) (*ThresholdProfile, error)
	Delete(profileID string) error
	List() ([]*ThresholdProfile, error)
}

type thresholdProfilesEndpoint struct {
	client RequestFactory
}

// TODO(mohmann) implement ThresholdProfilesEndpoint interface here
