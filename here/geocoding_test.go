package here

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeocodingService_Route(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(readJSONFile("resources/geocoding_response.json"))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := NewGeocodingClient(httpClient, "appID", "appCode")
	geocodingResponse, _, err := client.Geocoding.Search("200 S Mathilda Sunnyvale CA", 9)
	assert.NotNil(t, geocodingResponse)
	assert.Nil(t, err)
}

func TestGeocodingService_AddressInBoundingBox(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(readJSONFile("resources/geocoding_address_response.json"))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := NewGeocodingClient(httpClient, "appID", "appCode")
	geocodingResponse, _, err := client.Geocoding.AddressInBoundingBox("1 main", [2]float32{42.3902, -71.1293}, [2]float32{42.3312, -71.0228}, 9)
	assert.NotNil(t, geocodingResponse)
	assert.Nil(t, err)
}

func TestGeocodingService_PartialAddressInformation(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(readJSONFile("resources/geocoding_partial_address.json"))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := NewGeocodingClient(httpClient, "appID", "appCode")
	geocodingResponse, _, err := client.Geocoding.PartialAddressInformation(425, "randolph", "chicago", "usa", 9)
	assert.NotNil(t, geocodingResponse)
	assert.Nil(t, err)
}
