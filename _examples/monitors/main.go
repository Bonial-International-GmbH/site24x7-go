package main

import (
	"fmt"
	"os"
	"time"

	site24x7 "github.com/Bonial-International-GmbH/site24x7-go"
	"github.com/Bonial-International-GmbH/site24x7-go/backoff"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	config := site24x7.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RefreshToken: os.Getenv("REFRESH_TOKEN"),

		// RetryConfig is optional. If omitted, backoff.DefaultRetryConfig will
		// be used.
		RetryConfig: &backoff.RetryConfig{
			MinWait:    1 * time.Second,
			MaxWait:    30 * time.Second,
			MaxRetries: 4,
			CheckRetry: backoff.DefaultRetryPolicy,
			Backoff:    backoff.DefaultBackoff,
		},
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

	notificationProfiles, err := client.NotificationProfiles().List()
	if err != nil {
		panic(err)
	}

	for _, profile := range notificationProfiles {
		fmt.Printf("%+v\n\n", profile)
	}

	itAutomations, err := client.ITAutomations().List()
	if err != nil {
		panic(err)
	}

	for _, automation := range itAutomations {
		fmt.Printf("%+v\n\n", automation)
	}
}
