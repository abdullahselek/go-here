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

func TestGeocodingService_Route(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(createSearchTextOkResponse())
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := NewGeocodingClient(httpClient)
	geocodingResponse, _, err := client.Geocoding.Search("200 S Mathilda Sunnyvale CA", "appId", "appCode", 9)
	assert.NotNil(t, geocodingResponse)
	assert.Nil(t, err)
}
