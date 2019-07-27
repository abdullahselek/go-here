package here

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createSearchTextOkResponse() []byte {
	// Open our jsonFile
	jsonFile, err := os.Open("resources/geocoding_response.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func createAddressInBoundingBoxOkResponse() []byte {
	// Open our jsonFile
	jsonFile, err := os.Open("resources/geocoding_address_response.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func createPartialAddressInformationOkResponse() []byte {
	// Open our jsonFile
	jsonFile, err := os.Open("resources/geocoding_partial_address.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func TestGeocodingService_Route(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(createSearchTextOkResponse())
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
		w.Write(createAddressInBoundingBoxOkResponse())
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
		w.Write(createPartialAddressInformationOkResponse())
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := NewGeocodingClient(httpClient, "appID", "appCode")
	geocodingResponse, _, err := client.Geocoding.PartialAddressInformation(425, "randolph", "chicago", "usa", 9)
	assert.NotNil(t, geocodingResponse)
	assert.Nil(t, err)
}
