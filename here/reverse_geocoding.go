package here

import "github.com/dghubble/sling"

// ReverseGeocodingService provides for HERE ReverseGeocoding api.
type ReverseGeocodingService struct {
	sling   *sling.Sling
	AppID   string
	AppCode string
}

// newGeocodingService returns a new GeocodingService.
func newReverseGeocodingService(sling *sling.Sling, appID string, appCode string) *ReverseGeocodingService {
	return &ReverseGeocodingService{
		sling:   sling,
		AppID:   appID,
		AppCode: appCode,
	}
}
