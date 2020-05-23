package here

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/dghubble/sling"
)

// FleetTelematicsService provides for HERE routing api.
type FleetTelematicsService struct {
	sling *sling.Sling
}

// DestinationParams params
type DestinationParams struct {
	Coordinates [2]float32
	Text        string
}

// FleetTelematicsParams parameters for Fleet Telematics Service.
type FleetTelematicsParams struct {
	Start         string `url:"start"`
	Destination1  string `url:"destination1,omitempty"`
	Destination2  string `url:"destination2,omitempty"`
	Destination3  string `url:"destination3,omitempty"`
	Destination4  string `url:"destination4,omitempty"`
	Destination5  string `url:"destination5,omitempty"`
	Destination6  string `url:"destination6,omitempty"`
	Destination7  string `url:"destination7,omitempty"`
	Destination8  string `url:"destination8,omitempty"`
	Destination9  string `url:"destination9,omitempty"`
	Destination10 string `url:"destination10,omitempty"`
	End           string `url:"end"`
	APIKey        string `url:"apikey"`
	Modes         string `url:"mode"`
	Departure     string `url:"departure"`
}

// FleetTelematicsWaypointsSequenceResponse model for fleet telematics service response
type FleetTelematicsWaypointsSequenceResponse struct {
	Results []struct {
		Waypoints []struct {
			ID                   string    `json:"id"`
			Lat                  float64   `json:"lat"`
			Lon                  float64   `json:"lon"`
			Sequence             uint8     `json:"sequence"`
			EstimatedArrival     time.Time `json:"estimatedArrival"`
			EstimatedDeparture   time.Time `json:"estimatedDeparture"`
			FulFilledConstraints []string  `json:"fulfilledConstraints"`
		} `json:"waypoints"`
		Distance         string `json:"distance"`
		Time             string `json:"time"`
		InterConnections []struct {
			FromWaypoint string  `json:"fromWaypoint"`
			ToWaypoint   string  `json:"toWaypoint"`
			Distance     float64 `json:"distance"`
			Time         float64 `json:"time"`
			Rest         float64 `json:"rest"`
			Waiting      float64 `json:"waiting"`
		} `json:"interconnections"`
		Description   string `json:"description"`
		TimeBreakdown struct {
			Driving int `json:"driving"`
			Service int `json:"service"`
			Rest    int `json:"rest"`
			Waiting int `json:"waiting"`
		}
	}
	Errors             []string `json:"errors"`
	ProcessingTimeDesc string   `json:"processingTimeDesc"`
	ResponseCode       string   `json:"responseCode"`
}

// newFleetTelematicsService returns a new RoutingService.
func newFleetTelematicsService(sling *sling.Sling) *FleetTelematicsService {
	return &FleetTelematicsService{
		sling: sling,
	}
}

// Returns destinations as a formatted string.
func createDestination(destination DestinationParams) string {
	if destination.Text != "" {
		return fmt.Sprintf("%s;%f,%f;", destination.Text, destination.Coordinates[0], destination.Coordinates[1])
	}
	return fmt.Sprintf("%f,%f;", destination.Coordinates[0], destination.Coordinates[1])
}

// CreateFleetTelematicsParams creates fleet telematics parameters struct.
func (s *FleetTelematicsService) CreateFleetTelematicsParams(start DestinationParams, end DestinationParams, destinations []DestinationParams, apiKey string, modes []Enum) FleetTelematicsParams {
	var buffer bytes.Buffer
	for _, routeMode := range modes {
		mode := Enum.ValueOfRouteMode(routeMode)
		buffer.WriteString(mode + ";")
	}
	routeModes := buffer.String()
	routeModes = routeModes[:len(routeModes)-1]
	fleetTelematicsParams := FleetTelematicsParams{
		Start:     createDestination(start),
		APIKey:    apiKey,
		Modes:     routeModes,
		Departure: "now",
		End:       createDestination(end),
	}

	for i, destination := range destinations {
		stringDestination := createDestination(destination)
		concatenated := "Destination" + strconv.Itoa(i+1)
		reflect.ValueOf(&fleetTelematicsParams).Elem().FieldByName(concatenated).SetString(stringDestination)
	}

	fmt.Printf("%v", fleetTelematicsParams)
	return fleetTelematicsParams
}

// FindSequence with given parameters.
func (s *FleetTelematicsService) FindSequence(params *FleetTelematicsParams) (*FleetTelematicsWaypointsSequenceResponse, *http.Response, error) {
	routeSequence := new(FleetTelematicsWaypointsSequenceResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("findsequence.json").QueryStruct(params).Receive(routeSequence, apiError)
	return routeSequence, resp, relevantError(err, *apiError)
}
