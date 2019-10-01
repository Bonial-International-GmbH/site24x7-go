package main

import (
	"fmt"
	"os"

	site24x7 "github.com/Bonial-International-GmbH/site24x7-go"
)

func main() {
	config := site24x7.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RefreshToken: os.Getenv("REFRESH_TOKEN"),
	}

	client := site24x7.New(config)

	monitors, err := client.Monitors().List()
	if err != nil {
		panic(err)
	}

	for _, monitor := range monitors {
		fmt.Printf("%+v\n\n", monitor)
	}

	monitorGroups, err := client.MonitorGroups().List()
	if err != nil {
		panic(err)
	}

	for _, group := range monitorGroups {
		fmt.Printf("%+v\n\n", group)
	}

	locationProfiles, err := client.LocationProfiles().List()
	if err != nil {
		panic(err)
	}

	for _, profile := range locationProfiles {
		fmt.Printf("%+v\n\n", profile)
	}
}
