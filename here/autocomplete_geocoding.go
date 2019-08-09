package here

import "github.com/dghubble/sling"

// AutocompleteGeocodingService provides for HERE AutocompleteGeocoding api.
type AutocompleteGeocodingService struct {
	sling   *sling.Sling
	AppID   string
	AppCode string
}

// newAutocompleteGeocodingService returns a new AutocompleteGeocodingService.
func newAutocompleteGeocodingService(sling *sling.Sling, appID string, appCode string) *AutocompleteGeocodingService {
	return &AutocompleteGeocodingService{
		sling:   sling,
		AppID:   appID,
		AppCode: appCode,
	}
}
