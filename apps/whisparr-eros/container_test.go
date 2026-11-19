package main

import (
	"testing"
	"time"

	helpers "github.com/home-operations/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/home-operations/whisparr-eros:rolling")
	helpers.RequireHTTPEndpoint(t, image, helpers.HTTPTestConfig{
		Port:    "6969",
		Path:    "/ping",
		Timeout: 2 * time.Minute,
	}, nil)
}
