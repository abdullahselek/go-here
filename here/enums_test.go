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
