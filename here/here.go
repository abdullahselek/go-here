package here

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Client is a HERE client for making HERE API requests.
type Client struct {
	sling *sling.Sling
	// HERE API Services
	Routing          *RoutingService
	Geocoding        *GeocodingService
	ReverseGeocoding *ReverseGeocodingService
}

// NewRoutingClient returns a new RoutingClient.
func NewRoutingClient(httpClient *http.Client, appID string, appCode string) *Client {
	base := sling.New().Client(httpClient).Base("https://route.api.here.com/routing/7.2/")
	return &Client{
		sling:   base,
		Routing: newRoutingService(base.New(), appID, appCode),
	}
}

// NewGeocodingClient returns a new GeocodingClient.
func NewGeocodingClient(httpClient *http.Client, appID string, appCode string) *Client {
	base := sling.New().Client(httpClient).Base("https://geocoder.api.here.com/6.2/geocode.json")
	return &Client{
		sling:     base,
		Geocoding: newGeocodingService(base.New(), appID, appCode),
	}
}

// NewReverseGeocodingClient returns a new ReverseGeocodingClient.
func NewReverseGeocodingClient(httpClient *http.Client, appID string, appCode string) *Client {
	base := sling.New().Client(httpClient).Base("https://reverse.geocoder.api.here.com/6.2/")
	return &Client{
		sling:            base,
		ReverseGeocoding: newReverseGeocodingService(base.New(), appID, appCode),
	}
}
