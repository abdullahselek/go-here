package here

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutingService_Route(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(readJSONFile("resources/routing_response.json"))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := NewRoutingClient(httpClient, "appId", "appCode")
	routes, _, err := client.Routing.Route([2]float32{52.5160, 13.3779}, [2]float32{52.5206, 13.3862}, []RouteMode{Fastest, Car, TrafficDefault})
	assert.NotNil(t, routes)
	assert.Nil(t, err)
}
