package here

// Enum a kind of alias for int values.
type Enum int

type routeModeList struct {
	Fastest         Enum
	Car             Enum
	TrafficDisabled Enum
	Enabled         Enum
	Pedestrian      Enum
	PublicTransport Enum
	Truck           Enum
	TrafficDefault  Enum
	Bicycle         Enum
}

// RouteMode for public use.
var RouteMode = &routeModeList{
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

// ValueOfRouteMode returns value for RouteMode.
func (mode Enum) ValueOfRouteMode() string {
	modes := [...]string{
		"fastest",
		"car",
		"traffic:disabled",
		"enabled",
		"pedestrian",
		"publicTransport",
		"truck",
		"traffic:default",
		"bicycle",
	}
	if mode < RouteMode.Fastest || mode > RouteMode.TrafficDefault {
		return "Unknown"
	}
	return modes[mode]
}

type reverseGeocodingList struct {
	RetrieveAddresses Enum
	RetrieveAreas     Enum
	RetrieveLandmarks Enum
	RetrieveAll       Enum
	TrackPosition     Enum
}

// ReverseGeocodingMode for public use
var ReverseGeocodingMode = &reverseGeocodingList{
	RetrieveAddresses: 0,
	RetrieveAreas:     1,
	RetrieveLandmarks: 2,
	RetrieveAll:       3,
	TrackPosition:     4,
}

// ValueOfReverseGeocodingMode returns value for ReverseGeocodingMode.
func (mode Enum) ValueOfReverseGeocodingMode() string {
	modes := [...]string{
		"retrieveAddresses",
		"retrieveAreas",
		"retrieveLandmarks",
		"retrieveAll",
		"trackPosition",
	}
	if mode < ReverseGeocodingMode.RetrieveAddresses || mode > ReverseGeocodingMode.TrackPosition {
		return "Unknown"
	}
	return modes[mode]
}
