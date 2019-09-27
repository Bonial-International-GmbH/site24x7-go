package main

import (
	"context"
	"fmt"
	"os"

	site24x7 "github.com/Bonial-International-GmbH/site24x7-go"
	"github.com/Bonial-International-GmbH/site24x7-go/oauth"
)

func main() {
	config := oauth.NewConfig(
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		os.Getenv("REFRESH_TOKEN"),
	)

	httpClient := config.Client(context.Background())

	client := site24x7.NewClient(httpClient)

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

}
