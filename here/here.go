package here

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Client is a HERE client for making HERE API requests.
type Client struct {
	sling *sling.Sling
	// HERE API Services
	Routing               *RoutingService
	FleetTelematics       *FleetTelematicsService
	Geocoding             *GeocodingService
	ReverseGeocoding      *ReverseGeocodingService
	AutocompleteGeocoding *AutocompleteGeocodingService
	Weather               *WeatherService
	Places                *PlacesService
}

// NewRoutingClient returns a new RoutingClient.
func NewRoutingClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base("https://route.api.here.com/routing/7.2/")
	return &Client{
		sling:   base,
		Routing: newRoutingService(base.New()),
	}
}

// NewFleetTelematicsClient returns a new NewFleetTelematicsClient.
func NewFleetTelematicsClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base("https://wse.ls.hereapi.com/2/")
	return &Client{
		sling:           base,
		FleetTelematics: newFleetTelematicsService(base.New()),
	}
}

// NewGeocodingClient returns a new GeocodingClient.
func NewGeocodingClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base("https://geocoder.api.here.com/6.2/")
	return &Client{
		sling:     base,
		Geocoding: newGeocodingService(base.New()),
	}
}

// NewReverseGeocodingClient returns a new ReverseGeocodingClient.
func NewReverseGeocodingClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base("https://reverse.geocoder.api.here.com/6.2/")
	return &Client{
		sling:            base,
		ReverseGeocoding: newReverseGeocodingService(base.New()),
	}
}

// NewAutocompleteGeocodingClient returns a new AutocompleteGeocodingService.
func NewAutocompleteGeocodingClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base("https://autocomplete.geocoder.api.here.com/6.2/")
	return &Client{
		sling:                 base,
		AutocompleteGeocoding: newAutocompleteGeocodingService(base.New()),
	}
}

// NewWeatherClient returns a new WeatherService.
func NewWeatherClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base("https://weather.api.here.com/weather/1.0/")
	return &Client{
		sling:   base,
		Weather: newWeatherService(base.New()),
	}
}

// NewPlacesClient returns a new PlacesService.
func NewPlacesClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base("https://places.ls.hereapi.com/places/v1/")
	return &Client{
		sling:  base,
		Places: newPlacesService(base.New()),
	}
}
