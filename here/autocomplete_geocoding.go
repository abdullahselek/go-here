package here

import (
	"net/http"
	"net/url"

	"github.com/dghubble/sling"
)

// AutocompleteGeocodingService provides for HERE AutocompleteGeocoding api.
type AutocompleteGeocodingService struct {
	sling   *sling.Sling
	AppID   string
	AppCode string
}

// DetailsForSuggestionParameters parameters by search text for Geocoding Service.
type DetailsForSuggestionParameters struct {
	Query   string `url:"query"`
	AppID   string `url:"app_id"`
	AppCode string `url:"app_code"`
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
func newAutocompleteGeocodingService(sling *sling.Sling, appID string, appCode string) *AutocompleteGeocodingService {
	return &AutocompleteGeocodingService{
		sling:   sling,
		AppID:   appID,
		AppCode: appCode,
	}
}

// DetailsForSuggestion returns a list of address suggestions for the search text.
func (s *AutocompleteGeocodingService) DetailsForSuggestion(query string) (*AutocompleteGeocodingResponse, *http.Response, error) {
	detailsForSuggestionParameters := &DetailsForSuggestionParameters{Query: url.QueryEscape(query), AppID: s.AppID, AppCode: s.AppCode}
	autocompleteGeocodingResponse := new(AutocompleteGeocodingResponse)
	resp, err := s.sling.New().Get("suggest.json").QueryStruct(detailsForSuggestionParameters).ReceiveSuccess(autocompleteGeocodingResponse)
	return autocompleteGeocodingResponse, resp, err
}
