package here

// Route modes for Routing API.
type Route int

type list struct {
	Fastest         Route
	Car             Route
	TrafficDisabled Route
	Enabled         Route
	Pedestrian      Route
	PublicTransport Route
	Truck           Route
	TrafficDefault  Route
	Bicycle         Route
}

// Mode for public use
var Mode = &list{
	Fastest:         0,
	Car:             1,
	TrafficDisabled: 2,
	Enabled:         3,
	Pedestrian:      4,
	PublicTransport: 5,
	Truck:           6,
	TrafficDefault:  7,
	Bicycle:         8,
}

func (routeMode Route) String() string {
	modes := [...]string{
		"fastest",
		"car",
		"traffic:disabled",
		"enabled",
		"pedestrian",
		"publicTransport",
		"truck",
		"traffic:default",
		"bicycle"}
	if routeMode < Mode.Fastest || routeMode > Mode.TrafficDefault {
		return "Unknown"
	}
	return modes[routeMode]
}

// ReverseGeocodingMode modes for ReverseGeocoding API.
type ReverseGeocodingMode int

const (
	RetrieveAddresses ReverseGeocodingMode = 0
	RetrieveAreas     ReverseGeocodingMode = 1
	RetrieveLandmarks ReverseGeocodingMode = 2
	RetrieveAll       ReverseGeocodingMode = 3
	TrackPosition     ReverseGeocodingMode = 4
)

func (reverseGeocodingMode ReverseGeocodingMode) String() string {
	modes := [...]string{
		"retrieveAddresses",
		"retrieveAreas",
		"retrieveLandmarks",
		"retrieveAll",
		"trackPosition"}
	if reverseGeocodingMode < RetrieveAddresses || reverseGeocodingMode > TrackPosition {
		return "Unknown"
	}
	return modes[reverseGeocodingMode]
}
