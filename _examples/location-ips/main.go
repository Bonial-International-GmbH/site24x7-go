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

	ipProvider, err := location.NewDefaultProfileIPProvider(client)
	if err != nil {
		panic(err)
	}

	profiles, err := client.LocationProfiles().List()
	if err != nil {
		panic(err)
	}

	for _, profile := range profiles {
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
