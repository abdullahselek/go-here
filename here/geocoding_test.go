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

	client := NewGeocodingClient(httpClient)
	searchTextParams := &SearchTextParameters{SearchText: "200 S Mathilda Sunnyvale CA", APIKey: "apiKey", Gen: 9}
	geocodingResponse, _, err := client.Geocoding.Search(searchTextParams)
	assert.NotNil(t, geocodingResponse)
	assert.Nil(t, err)
}

func TestGeocodingService_AddressInBoundingBox(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(readJSONFile("resources/geocoding_address_response.json"))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := NewGeocodingClient(httpClient)
	params := &AddressInBoundingBoxParameters{SearchText: "1 main", MapView: client.Geocoding.CreateMapView([2]float32{42.3902, -71.1293}, [2]float32{42.3312, -71.0228}), Gen: 9, APIKey: "apiKey"}
	geocodingResponse, _, err := client.Geocoding.AddressInBoundingBox(params)
	assert.NotNil(t, geocodingResponse)
	assert.Nil(t, err)
}

func TestGeocodingService_PartialAddressInformation(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(readJSONFile("resources/geocoding_partial_address.json"))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := NewGeocodingClient(httpClient)
	params := &PartialAddressInformationParameters{HouseNumber: 425, Street: "randolph", City: "chicago", Country: "usa", Gen: 9, APIKey: "apiKey"}
	geocodingResponse, _, err := client.Geocoding.PartialAddressInformation(params)
	assert.NotNil(t, geocodingResponse)
	assert.Nil(t, err)
}
