package here

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouteMode(t *testing.T) {
	mode := Enum.ValueOfRouteMode(RouteMode.Fastest)
	assert.Equal(t, mode, "fastest")
	mode = Enum.ValueOfRouteMode(RouteMode.TrafficDefault)
	assert.Equal(t, mode, "traffic:default")
}

func TestReverseGeocodingMode(t *testing.T) {
	mode := Enum.ValueOfReverseGeocodingMode(ReverseGeocodingMode.RetrieveAddresses)
	assert.Equal(t, mode, "retrieveAddresses")
	mode = Enum.ValueOfReverseGeocodingMode(ReverseGeocodingMode.TrackPosition)
	assert.Equal(t, mode, "trackPosition")
}

func TestWeatherProduct(t *testing.T) {
	product := Enum.ValueOfWeatherProduct(WeatherProduct.Observation)
	assert.Equal(t, product, "observation")
	product = Enum.ValueOfWeatherProduct(WeatherProduct.NwsAlerts)
	assert.Equal(t, product, "nws_alerts")
}
