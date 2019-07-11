package here

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Client is a HERE client for making HERE API requests.
type Client struct {
	sling *sling.Sling
	// HERE API Services
	Routing   *RoutingService
	Geocoding *GeocodingService
}

// NewRoutingClient returns a new RoutingClient.
func NewRoutingClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base("https://route.api.here.com/routing/7.2/")
	return &Client{
		sling:   base,
		Routing: newRoutingService(base.New()),
	}
}

// NewGeocodingClient returns a new GeocodingClient.
func NewGeocodingClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base("https://geocoder.api.here.com/6.2/geocode.json")
	return &Client{
		sling:     base,
		Geocoding: newGeocodingService(base.New()),
	}
}
