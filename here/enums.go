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

type weatherProductList struct {
	Observation         Enum
	Forecast7days       Enum
	Forecast7DaysSimple Enum
	ForecastHourly      Enum
	ForecastAstronomy   Enum
	Alerts              Enum
	NwsAlerts           Enum
}

// WeatherProduct for public use
var WeatherProduct = &weatherProductList{
	Observation:         0,
	Forecast7days:       1,
	Forecast7DaysSimple: 2,
	ForecastHourly:      3,
	ForecastAstronomy:   4,
	Alerts:              5,
	NwsAlerts:           6,
}

// ValueOfWeatherProduct returns value for WeatherProduct.
func (mode Enum) ValueOfWeatherProduct() string {
	modes := [...]string{
		"observation",
		"forecast_7days",
		"forecast_7days_simple",
		"forecast_hourly",
		"forecast_astronomy",
		"alerts",
		"nws_alerts",
	}
	if mode < WeatherProduct.Observation || mode > WeatherProduct.NwsAlerts {
		return "Unknown"
	}
	return modes[mode]
}
