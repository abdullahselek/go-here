package here

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAutocompleteGeocodingService_DetailsForSuggestion(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(readJSONFile("resources/autocomplete_details_response.json"))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := NewAutocompleteGeocodingClient(httpClient, "appID", "appCode")
	autocompleteGeocodingResponse, _, err := client.AutocompleteGeocoding.DetailsForSuggestion("Pariser 1 Berl")
	assert.NotNil(t, autocompleteGeocodingResponse)
	assert.Nil(t, err)
}
