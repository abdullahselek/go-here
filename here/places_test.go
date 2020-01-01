package here

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlacesService_PlaceCategories(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(readJSONFile("resources/place_categories_response.json"))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := NewPlacesClient(httpClient)
	params := client.Places.CreatePlaceCategoriesParameters([2]float32{52.5160, 13.3779}, "apiKey")
	categories, _, err := client.Places.PlaceCategories(&params)
	assert.NotNil(t, categories)
	assert.Nil(t, err)
}
