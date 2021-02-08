// Command diffmodel compares our model and the server model.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/bassosimone/apiclient/internal/fatalx"
	"github.com/bassosimone/apiclient/internal/openapi"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
)

func makeModel(data []byte) *openapi.Swagger {
	var out openapi.Swagger
	err := json.Unmarshal(data, &out)
	fatalx.OnError(err, "json.Unmarshal failed")
	// We reduce irrelevant differences by producing a common header
	return &openapi.Swagger{Paths: out.Paths}
}

func getServerModel() *openapi.Swagger {
	resp, err := http.Get("https://api.ooni.io/apispec_1.json")
	fatalx.OnError(err, "http.Get failed")
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fatalx.OnError(err, "ioutil.ReadAll failed")
	return makeModel(data)
}

func getClientModel(clientFile string) *openapi.Swagger {
	data, err := ioutil.ReadFile(clientFile)
	fatalx.OnError(err, "ioutil.ReadFile failed")
	return makeModel(data)
}

func simplifyRoundTrip(rt *openapi.RoundTrip) {
	// This is a quirk that needs to be fixed upstream. We are not
	// going to focus on it for now so that we reduce noise.
	rt.Consumes, rt.Produces = nil, nil

	// Normalize the used name when a parameter is in body
	for _, param := range rt.Parameters {
		if param.In == "body" {
			param.Name = "body"
		}
	}

	// Sort parameters so the comparison does not depend on order.
	sort.SliceStable(rt.Parameters, func(i, j int) bool {
		left, right := rt.Parameters[i].Name, rt.Parameters[j].Name
		return strings.Compare(left, right) < 0
	})

	// Normalize description of 200 response
	rt.Responses.Successful.Description = "all good"
}

func simplifyInPlace(path *openapi.Path) *openapi.Path {
	if path.Get != nil {
		simplifyRoundTrip(path.Get)
	}
	if path.Post != nil {
		simplifyRoundTrip(path.Post)
	}
	return path
}

func jsonify(model interface{}) string {
	data, err := json.MarshalIndent(model, "", "    ")
	fatalx.OnError(err, "json.MarshalIndent failed")
	return string(data)
}

type diffable struct {
	name  string
	value string
}

func computediff(server, client *diffable) string {
	d := gotextdiff.ToUnified(server.name, client.name, server.value, myers.ComputeEdits(
		span.URIFromPath(server.name),
		server.value,
		client.value,
	))
	return fmt.Sprint(d)
}

// maybediff emits the diff between the server and the client and
// returns the length of the diff itself in bytes.
func maybediff(key string, server, client *openapi.Path) int {
	diff := computediff(&diffable{
		name:  fmt.Sprintf("server%s.json", key),
		value: jsonify(simplifyInPlace(server)),
	}, &diffable{
		name:  fmt.Sprintf("client%s.json", key),
		value: jsonify(simplifyInPlace(client)),
	})
	if diff != "" {
		fmt.Printf("%s", diff)
	}
	return len(diff)
}

func compare(clientFile string) int {
	var code int
	serverModel, clientModel := getServerModel(), getClientModel(clientFile)
	for key := range serverModel.Paths {
		if _, found := clientModel.Paths[key]; !found {
			delete(serverModel.Paths, key)
			continue
		}
		if maybediff(key, serverModel.Paths[key], clientModel.Paths[key]) > 0 {
			code = 1
		}
	}
	return code
}

func main() {
	os.Exit(compare("swagger.json"))
}
