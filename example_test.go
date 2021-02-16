package apiclient_test

import (
	"context"
	"fmt"
	"log"

	"github.com/bassosimone/apiclient"
)

func ExampleURLsAPI_Call() {
	var api apiclient.URLsAPI
	request := &apiclient.URLsRequest{
		CountryCode: "IT",
		Limit:       14,
	}
	ctx := context.Background()
	response, err := api.Call(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", response)
}

func ExampleTorTargetsAPI_Call() {
	api := apiclient.TorTargetsAPI{
		Authorizer: apiclient.NewStaticAuthorizer("valid-token-here"),
	}
	request := &apiclient.TorTargetsRequest{}
	ctx := context.Background()
	response, err := api.Call(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d", len(response))
}
