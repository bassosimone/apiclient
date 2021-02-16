package apiclient_test

import (
	"context"
	"fmt"
	"log"

	"github.com/bassosimone/apiclient"
)

func ExampleURLsAPI_Call() {
	api := apiclient.NewURLsAPI(&apiclient.Client{})
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
	api := apiclient.NewTorTargetsAPI(&apiclient.Client{})
	request := &apiclient.TorTargetsRequest{}
	ctx := context.Background()
	response, err := api.Call(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d", len(response))
}
