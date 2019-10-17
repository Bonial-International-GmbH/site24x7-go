site24x7-go
===========

[![Build Status](https://travis-ci.org/Bonial-International-GmbH/site24x7-go.svg?branch=master)](https://travis-ci.org/Bonial-International-GmbH/site24x7-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Bonial-International-GmbH/site24x7-go?style=flat)](https://goreportcard.com/report/github.com/Bonial-International-GmbH/site24x7-go)
[![GoDoc](https://godoc.org/github.com/Bonial-International-GmbH/site24x7-go?status.svg)](https://godoc.org/github.com/Bonial-International-GmbH/site24x7-go)

An API client for [Site24x7](https://www.site24x7.com) written in go. Currently
the following endpoints are implemented:

* IT Automations
* Location Profiles
* Monitor Groups
* Monitors
* Notification Profiles
* Threshold Profiles
* User Groups

If you want to add support for other endpoints as well or want to add missing
fields to an existing endpoint, we are happy to accept Pull Requests.

Installation
------------

```
go get -u github.com/Bonial-International-GmbH/site24x7-go
```

Usage Example
-------------

The API client needs the OAuth2 client ID, client secret and refresh token to
authenticate against Site24x7 and obtain OAuth access tokens. Refer to the
[official documentation](https://www.site24x7.com/help/api/#authentication) for
instructions on how to obtain these credentials.

Here is a very simple example of creating a website monitor:

```go
package main

import (
	"fmt"
	"os"
	"time"

	site24x7 "github.com/Bonial-International-GmbH/site24x7-go"
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	apierrors "github.com/Bonial-International-GmbH/site24x7-go/api/errors"
	"github.com/Bonial-International-GmbH/site24x7-go/backoff"
)

func main() {
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

	monitor := &api.Monitor{
		DisplayName: "my monitor",
		Website:     "https://example.com",
	}

	monitor, err := client.Monitors().Create(monitor)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Monitor %q created (ID: %s)\n", monitor.DisplayName, monitor.MonitorID)
	fmt.Printf("%+v\n\n", monitor)

	monitorID := "123"

	monitor, err := client.Monitors().Get(monitorID)
	if apierrors.IsNotFound(err) {
		fmt.Printf("monitor %s not found\n", monitorID)
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("%+v\n\n", monitor)
	}
}
```

Refer to the
[godoc](https://godoc.org/github.com/Bonial-International-GmbH/site24x7-go) for
all available endpoints and
[API types](https://godoc.org/github.com/Bonial-International-GmbH/site24x7-go/api).

License
-------

The source code of site24x7-go is released under the MIT License. See the
bundled LICENSE file for details.
