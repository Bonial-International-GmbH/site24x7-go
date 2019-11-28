package location

import (
	site24x7 "github.com/Bonial-International-GmbH/site24x7-go"
	"github.com/Bonial-International-GmbH/site24x7-go/api"
)

// ProfileIPProvider provides the check location IP addresses for location
// profiles.
type ProfileIPProvider struct {
	Locations []*api.Location
	IPSource  IPSource
}

// NewDefaultProfileIPProvider creates a new *ProfileIPProvider which performs
// IP lookups via DNS. The available locations are obtained via the
// site24x7.Client. Returns an error of retrieving the location list fails.
func NewDefaultProfileIPProvider(client site24x7.Client) (*ProfileIPProvider, error) {
	locationTemplate, err := client.LocationTemplate().Get()
	if err != nil {
		return nil, err
	}

	p := &ProfileIPProvider{
		Locations: locationTemplate.Locations,
		IPSource:  NewDefaultDNSIPSource(),
	}

	return p, nil
}

// GetLocationIPs performs a lookup of the check IPs of all primary and
// secondary locations associated with the location profile.
func (p *ProfileIPProvider) GetLocationIPs(profile *api.LocationProfile) ([]string, error) {
	locationIDs := append([]string{profile.PrimaryLocation}, profile.SecondaryLocations...)

	var ips []string

	for _, locationID := range locationIDs {
		location, found := p.getLocation(locationID)
		if !found {
			continue
		}

		locationIPs, err := p.IPSource.LookupIPs(location)
		if err != nil {
			return nil, err
		}

		ips = append(ips, locationIPs...)
	}

	return ips, nil
}

func (p *ProfileIPProvider) getLocation(locationID string) (*api.Location, bool) {
	for _, location := range p.Locations {
		if location.LocationID == locationID {
			return location, true
		}
	}

	return nil, false
}
