package here

import (
	"net/http"
	"net/url"

	"github.com/dghubble/sling"
)

// AutocompleteGeocodingService provides for HERE AutocompleteGeocoding api.
type AutocompleteGeocodingService struct {
	sling *sling.Sling
}

// DetailsForSuggestionParameters parameters by search text for Geocoding Service.
type DetailsForSuggestionParameters struct {
	Query  string `url:"query"`
	APIKey string `url:"apikey"`
}

// AutocompleteGeocodingResponse response model for autocomplete geocoding service.
type AutocompleteGeocodingResponse struct {
	Suggestions []struct {
		Label       string `json:"label"`
		Language    string `json:"language"`
		CountryCode string `json:"countryCode"`
		LocationID  string `json:"locationId"`
		Address     struct {
			Country     string `json:"country"`
			State       string `json:"state"`
			County      string `json:"county"`
			City        string `json:"city"`
			District    string `json:"district"`
			Street      string `json:"street"`
			HouseNumber string `json:"houseNumber,omitempty"`
			PostalCode  string `json:"postalCode"`
		} `json:"address,omitempty"`
		MatchLevel string `json:"matchLevel"`
	} `json:"suggestions"`
}

// newAutocompleteGeocodingService returns a new AutocompleteGeocodingService.
func newAutocompleteGeocodingService(sling *sling.Sling) *AutocompleteGeocodingService {
	return &AutocompleteGeocodingService{
		sling: sling,
	}
}

// CreateDetailsForSuggestionParameters creates DetailsForSuggestionParameters parameters by search text for Geocoding Service.
func (s *AutocompleteGeocodingService) CreateDetailsForSuggestionParameters(query string, apiKey string) DetailsForSuggestionParameters {
	detailsForSuggestionParameters := DetailsForSuggestionParameters{Query: url.QueryEscape(query), APIKey: apiKey}
	return detailsForSuggestionParameters
}

// DetailsForSuggestion returns a list of address suggestions for the search text.
func (s *AutocompleteGeocodingService) DetailsForSuggestion(params *DetailsForSuggestionParameters) (*AutocompleteGeocodingResponse, *http.Response, error) {
	autocompleteGeocodingResponse := new(AutocompleteGeocodingResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("suggest.json").QueryStruct(&params).Receive(autocompleteGeocodingResponse, apiError)
	return autocompleteGeocodingResponse, resp, relevantError(err, *apiError)
}
