package apiclient_test

import (
	"context"
	"fmt"
	"log"

	"github.com/bassosimone/apiclient"
	"github.com/bassosimone/apiclient/model"
)

func ExampleClient_URLs() {
	clnt := &apiclient.Client{}
	request := &model.URLsRequest{
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
	request := &model.TorTargetsRequest{}
	ctx := context.Background()
	response, err := clnt.TorTargets(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d", len(response))
}
