package here

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testingHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewTLSServer(handler)
	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return cli, s.Close
}

func createOkResponse() []byte {
	// Open our jsonFile
	jsonFile, err := os.Open("resources/routing_response.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func TestRoutingService_Route(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(createOkResponse())
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := NewRoutingClient(httpClient)
	routes, _, _ := client.Routing.Route([2]float32{52.5160, 13.3779}, [2]float32{52.5206, 13.3862}, "appId", "appCode", []RouteMode{Fastest, Car, TrafficDefault})
	assert.NotNil(t, routes)
	t.Log(routes)
}
