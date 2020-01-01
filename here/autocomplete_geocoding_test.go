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

	client := NewAutocompleteGeocodingClient(httpClient)
	params := client.AutocompleteGeocoding.CreateDetailsForSuggestionParameters("Pariser 1 Berl", "apiKey")
	autocompleteGeocodingResponse, _, err := client.AutocompleteGeocoding.DetailsForSuggestion(&params)
	assert.Equal(t, len(autocompleteGeocodingResponse.Suggestions), 5)
	assert.Nil(t, err)
}
