package main

import (
	"fmt"
	"os"

	site24x7 "github.com/Bonial-International-GmbH/site24x7-go"
	"github.com/Bonial-International-GmbH/site24x7-go/api"
)

func main() {
	config := site24x7.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RefreshToken: os.Getenv("REFRESH_TOKEN"),
	}

	client := site24x7.New(config)

	options := &api.CurrentStatusListOptions{
		APMRequired:       api.Bool(false),
		SuspendedRequired: api.Bool(true),
		StatusRequired:    api.String(fmt.Sprintf("%d", api.Down)),
	}

	monitorsStatus, err := client.CurrentStatus().List(options)
	if err != nil {
		panic(err)
	}

	for _, monitorStatus := range monitorsStatus.Monitors {
		fmt.Printf("%#v\n\n", monitorStatus)
	}
}
