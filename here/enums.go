package here

// RouteMode modes for Routing API.
type RouteMode int

const (
	Fastest         RouteMode = 0
	Car             RouteMode = 1
	TrafficDisabled RouteMode = 2
	Enabled         RouteMode = 3
	Pedestrian      RouteMode = 4
	PublicTransport RouteMode = 5
	Truck           RouteMode = 6
	TrafficDefault  RouteMode = 7
	Bicycle         RouteMode = 8
)

func (routeMode RouteMode) String() string {
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
	if routeMode < Fastest || routeMode > TrafficDefault {
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
