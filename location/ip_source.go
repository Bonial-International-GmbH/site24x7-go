package location

import (
	"fmt"
	"net"
	"strings"

	"github.com/Bonial-International-GmbH/site24x7-go/api"
)

const (
	// DefaultBaseDNSDomain is the domain where all Site24x7 check location
	// IPs are aggregated under.
	DefaultBaseDNSDomain = "enduserexp.com"
)

// IPSource provides the origin IP addresses for a Site24x7 check location.
type IPSource interface {
	// LookupIPs looks up all Site24x7 IPs associated with given location. The
	// resulting slice of net.IP values can be used for whitelisting IP addresses
	// in the firewalls of the endpoints monitors are configured for.
	LookupIPs(location *api.Location) ([]string, error)
}

// StaticIPSource looks up check location IPs from a static map.
type StaticIPSource struct {
	// LocationIPs is a mapping between location ID and the list of IP
	// addresses associated with this location.
	LocationIPs map[string][]string
}

// LookupIPs implements IPSource.
func (s *StaticIPSource) LookupIPs(location *api.Location) ([]string, error) {
	return s.LocationIPs[location.LocationID], nil
}

// DNSIPSource looks up check location IPs using DNS.
type DNSIPSource struct {
	// BaseDNSDomain is the DNS domain the aggregates all check location IPs
	// under subdomains with the naming scheme {{CityName}}-{{CountryCode}}.
	BaseDNSDomain string
}

// NewDefaultDNSIPSource creates a new *DNSIPSource value for the default
// BaseDNSDomain.
func NewDefaultDNSIPSource() *DNSIPSource {
	return &DNSIPSource{
		BaseDNSDomain: DefaultBaseDNSDomain,
	}
}

// LookupIPs implements IPSource.
func (s *DNSIPSource) LookupIPs(location *api.Location) ([]string, error) {
	countryCode, err := parseCountryCode(location)
	if err != nil {
		return nil, err
	}

	ips, err := s.lookupIPs(location.CityName, countryCode)
	if err != nil {
		return nil, fmt.Errorf("failed to lookup IPs for location %q: %v", location.DisplayName, err)
	}

	return netIPsToStrings(ips), nil
}

// lookupIPs performs IP lookup by city and country code and tries to handle
// some inconsistencies in the naming of the city specific DNS subdomains.
//
// This builds the DNS name that resolves to all IP addresses
// for given location. DNS names have the following form:
//
//   {{CityName}}-{{CountryCode}}.enduserexp.com
//
// For example, the DNS name for `Tel Aviv` in Israel is:
//
//   telaviv-il.enduserexp.com
//
// However, the DNS name for `Rio de Janeiro` in Brazil is not:
//
//   riodejaneiro-br.enduserexp.com
//
// but rather
//
//    rio-br.enduserexp.com
//
// For these cases, the following sequence of DNS lookups is attempted, bailing
// out on the first successful one:
//
//   1. riodejaneiro-br.enduserexp.com
//   2. riode-br.enduserexp.com
//   3. rio-br.enduserexp.com
func (s *DNSIPSource) lookupIPs(cityName string, countryCode string) (ips []net.IP, err error) {
	words := strings.Split(cityName, " ")

	for len(words) > 0 {
		cityName := strings.Join(words, "")
		host := fmt.Sprintf("%s-%s.%s", strings.ToLower(cityName), countryCode, s.BaseDNSDomain)

		ips, err = net.LookupIP(host)
		if err == nil {
			break
		}

		words = words[0 : len(words)-1]
	}

	return ips, err
}

func netIPsToStrings(netIPs []net.IP) []string {
	ips := make([]string, len(netIPs))
	for i, ip := range netIPs {
		if len(ip) == 0 {
			continue
		}

		ips[i] = ip.String()
	}

	return ips
}

// parseCountryCode parses the country code from the locations' display name.
// The display name has the format: {{CityName}} - {{CountryCode}}.
func parseCountryCode(location *api.Location) (string, error) {
	parts := strings.Split(location.DisplayName, " - ")
	if len(parts) != 2 || parts[1] == "" {
		return "", fmt.Errorf("failed to parse country code from location name %q", location.DisplayName)
	}

	return strings.ToLower(parts[1]), nil
}
