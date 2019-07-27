package here

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouteMode(t *testing.T) {
	mode := RouteMode.String(Fastest)
	assert.Equal(t, mode, "fastest")
	mode = RouteMode.String(TrafficDefault)
	assert.Equal(t, mode, "traffic:default")
}

func TestReverseGeocodingMode(t *testing.T) {
	mode := ReverseGeocodingMode.String(RetrieveAddresses)
	assert.Equal(t, mode, "retrieveAddresses")
	mode = ReverseGeocodingMode.String(TrackPosition)
	assert.Equal(t, mode, "trackPosition")
}
