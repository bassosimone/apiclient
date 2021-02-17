package apiclient_test

import (
	"context"
	"fmt"
	"log"

	"github.com/bassosimone/apiclient"
)

func ExampleClient_URLs() {
	clnt := &apiclient.Client{}
	request := &apiclient.URLsRequest{
		CountryCode: "IT",
		Limit:       14,
	}
	ctx := context.Background()
	response, err := clnt.URLs(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", response)
}

func ExampleClient_TorTargets() {
	clnt := &apiclient.Client{}
	request := &apiclient.TorTargetsRequest{}
	ctx := context.Background()
	response, err := clnt.TorTargets(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d", len(response))
}
