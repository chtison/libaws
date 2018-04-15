package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	r, err := Handler(events.APIGatewayProxyRequest{})
	if err != nil {
		t.Error(err)
	}
	t.Log(r.Body)
}
