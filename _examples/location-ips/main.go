package main

import (
	"fmt"
	"os"

	site24x7 "github.com/Bonial-International-GmbH/site24x7-go"
	"github.com/Bonial-International-GmbH/site24x7-go/location"
)

func main() {
	config := site24x7.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RefreshToken: os.Getenv("REFRESH_TOKEN"),
	}

	client := site24x7.New(config)

	// Using the ProfileIPProvider it is possible to retrieve a list of all IP
	// addresses associated with the locations of a LocationProfile. This is
	// useful it you want to do dynamic server-side IP whitelisting of Site24x7
	// check origins.
	ipProvider, err := location.NewDefaultProfileIPProvider(client)
	if err != nil {
		panic(err)
	}

	profiles, err := client.LocationProfiles().List()
	if err != nil {
		panic(err)
	}

	for _, profile := range profiles {
		// This will lookup all IP addresses associated to the locations
		// configured in the location profile.
		ips, err := ipProvider.GetLocationIPs(profile)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s:\n", profile.ProfileName)
		for _, ip := range ips {
			fmt.Printf("  %s\n", ip)
		}
		fmt.Println()
	}
}
