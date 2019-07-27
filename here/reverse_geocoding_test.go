package here

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseGeocodingService_AddressFromLocation(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(readJSONFile("resources/reverse_geocoding_address_response.json"))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := NewReverseGeocodingClient(httpClient, "appID", "appCode")
	geocodingResponse, _, err := client.ReverseGeocoding.AddressFromLocation([2]float32{42.3902, -71.1293}, 250, RetrieveAddresses, 1, 9)
	assert.NotNil(t, geocodingResponse)
	assert.Nil(t, err)
}
