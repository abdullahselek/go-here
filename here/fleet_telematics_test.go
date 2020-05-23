package here

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFleetTelematicsService_Route(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(readJSONFile("resources/fleet_telematics_find_sequence_response.json"))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client := NewFleetTelematicsClient(httpClient)

	start := DestinationParams{
		Coordinates: [2]float32{-25.643787, -49.158607},
		Text:        "Start",
	}

	end := DestinationParams{
		Coordinates: [2]float32{-20.778591, -51.591198},
		Text:        "End",
	}

	destinations := []DestinationParams{
		{
			Coordinates: [2]float32{-25.644764, -49.158558},
			Text:        "Destination1"},
		{
			Coordinates: [2]float32{-22.98319, -49.903282},
			Text:        "Destination2"},
		{
			Coordinates: [2]float32{-20.778555, -51.591164},
			Text:        "Destination3"},
	}

	params := client.FleetTelematics.CreateFleetTelematicsParams(start, end, destinations, "apiKey", []Enum{RouteMode.Fastest, RouteMode.Truck})
	routes, _, err := client.FleetTelematics.FindSequence(&params)
	// TO-DO IMPROVE VALIDATIONS
	assert.NotNil(t, routes)
	assert.Nil(t, err)
}
