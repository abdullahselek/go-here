package here

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// PlacesService provides for HERE Places api.
type PlacesService struct {
	sling *sling.Sling
}

// PlaceCategoriesParameters parameters to get categories a list of place categories available for a given location.
type PlaceCategoriesParameters struct {
	At     string `url:"at"`
	APIKey string `url:"apiKey"`
}

// PlaceCategoriesResponse model for place categories.
type PlaceCategoriesResponse struct {
	Items []struct {
		ID     string        `json:"id"`
		Title  string        `json:"title"`
		Icon   string        `json:"icon"`
		Type   string        `json:"type"`
		Href   string        `json:"href"`
		System string        `json:"system"`
		Within []interface{} `json:"within"`
	} `json:"items"`
}

// newPlacesService returns a new PlacesService.
func newPlacesService(sling *sling.Sling) *PlacesService {
	return &PlacesService{
		sling: sling,
	}
}

// CreatePlaceCategoriesParameters creates place categories parameters struct.
func (s *PlacesService) CreatePlaceCategoriesParameters(coordinates [2]float32, apiKey string) PlaceCategoriesParameters {
	at := fmt.Sprintf("%f,%f", coordinates[0], coordinates[1])
	placeCategoriesParameters := PlaceCategoriesParameters{
		At:     at,
		APIKey: apiKey,
	}
	return placeCategoriesParameters
}

// PlaceCategories returns place categories.
func (s *PlacesService) PlaceCategories(placeCategoriesParameters *PlaceCategoriesParameters) (*PlaceCategoriesResponse, *http.Response, error) {
	categories := new(PlaceCategoriesResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("categories/places").QueryStruct(placeCategoriesParameters).Receive(categories, apiError)
	return categories, resp, relevantError(err, *apiError)
}
