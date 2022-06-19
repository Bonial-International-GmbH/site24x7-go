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
		APIBaseURL:   "https://www.site24x7.eu/api",
		TokenURL:     "https://accounts.zoho.eu/oauth/v2/token",

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

	users, err := client.Users().List()
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Println(user.DisplayName)
		fmt.Println(user.EmailAddress)
		fmt.Println(user.UserID)
	}
}
