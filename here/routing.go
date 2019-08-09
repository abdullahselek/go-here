package here

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// RoutingService provides for HERE routing api.
type RoutingService struct {
	sling   *sling.Sling
	AppID   string
	AppCode string
}

// RoutingParams parameters for Routing Service.
type RoutingParams struct {
	Waypoint0 string `url:"waypoint0"`
	Waypoint1 string `url:"waypoint1"`
	AppID     string `url:"app_id"`
	AppCode   string `url:"app_code"`
	Modes     string `url:"mode"`
	Departure string `url:"departure"`
}

// RoutingResponse model for routing service.
type RoutingResponse struct {
	Response struct {
		MetaInfo struct {
			Timestamp           time.Time `json:"timestamp"`
			MapVersion          string    `json:"mapVersion"`
			ModuleVersion       string    `json:"moduleVersion"`
			InterfaceVersion    string    `json:"interfaceVersion"`
			AvailableMapVersion []string  `json:"availableMapVersion"`
		} `json:"metaInfo"`
		Route []struct {
			Waypoint []struct {
				LinkID         string `json:"linkId"`
				MappedPosition struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"mappedPosition"`
				OriginalPosition struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"originalPosition"`
				Type           string  `json:"type"`
				Spot           float64 `json:"spot"`
				SideOfStreet   string  `json:"sideOfStreet"`
				MappedRoadName string  `json:"mappedRoadName"`
				Label          string  `json:"label"`
				ShapeIndex     int     `json:"shapeIndex"`
				Source         string  `json:"source"`
			} `json:"waypoint"`
			Mode struct {
				Type           string        `json:"type"`
				TransportModes []string      `json:"transportModes"`
				TrafficMode    string        `json:"trafficMode"`
				Feature        []interface{} `json:"feature"`
			} `json:"mode"`
			Leg []struct {
				Start struct {
					LinkID         string `json:"linkId"`
					MappedPosition struct {
						Latitude  float64 `json:"latitude"`
						Longitude float64 `json:"longitude"`
					} `json:"mappedPosition"`
					OriginalPosition struct {
						Latitude  float64 `json:"latitude"`
						Longitude float64 `json:"longitude"`
					} `json:"originalPosition"`
					Type           string  `json:"type"`
					Spot           float64 `json:"spot"`
					SideOfStreet   string  `json:"sideOfStreet"`
					MappedRoadName string  `json:"mappedRoadName"`
					Label          string  `json:"label"`
					ShapeIndex     int     `json:"shapeIndex"`
					Source         string  `json:"source"`
				} `json:"start"`
				End struct {
					LinkID         string `json:"linkId"`
					MappedPosition struct {
						Latitude  float64 `json:"latitude"`
						Longitude float64 `json:"longitude"`
					} `json:"mappedPosition"`
					OriginalPosition struct {
						Latitude  float64 `json:"latitude"`
						Longitude float64 `json:"longitude"`
					} `json:"originalPosition"`
					Type           string  `json:"type"`
					Spot           float64 `json:"spot"`
					SideOfStreet   string  `json:"sideOfStreet"`
					MappedRoadName string  `json:"mappedRoadName"`
					Label          string  `json:"label"`
					ShapeIndex     int     `json:"shapeIndex"`
					Source         string  `json:"source"`
				} `json:"end"`
				Length     int `json:"length"`
				TravelTime int `json:"travelTime"`
				Maneuver   []struct {
					Position struct {
						Latitude  float64 `json:"latitude"`
						Longitude float64 `json:"longitude"`
					} `json:"position"`
					Instruction string `json:"instruction"`
					TravelTime  int    `json:"travelTime"`
					Length      int    `json:"length"`
					ID          string `json:"id"`
					Type        string `json:"_type"`
				} `json:"maneuver"`
			} `json:"leg"`
			Summary struct {
				Distance    int      `json:"distance"`
				TrafficTime int      `json:"trafficTime"`
				BaseTime    int      `json:"baseTime"`
				Flags       []string `json:"flags"`
				Text        string   `json:"text"`
				TravelTime  int      `json:"travelTime"`
				Type        string   `json:"_type"`
			} `json:"summary"`
		} `json:"route"`
		Language string `json:"language"`
	} `json:"response"`
}

// newRoutingService returns a new RoutingService.
func newRoutingService(sling *sling.Sling, appID string, appCode string) *RoutingService {
	return &RoutingService{
		sling:   sling,
		AppID:   appID,
		AppCode: appCode,
	}
}

// Returns waypoints as a formatted string.
func createWaypoint(waypoint [2]float32) string {
	waypoints := fmt.Sprintf("%f,%f", waypoint[0], waypoint[1])
	return waypoints
}

// Creates routing parameters struct.
func createRoutingParams(waypoint0 [2]float32, waypoint1 [2]float32, appID string, appCode string, modes []Route) RoutingParams {
	stringWaypoint0 := createWaypoint(waypoint0)
	stringWaypoint1 := createWaypoint(waypoint1)
	var buffer bytes.Buffer
	for _, routeMode := range modes {
		mode := Route.String(routeMode)
		buffer.WriteString(mode + ";")
	}
	routeModes := buffer.String()
	routeModes = routeModes[:len(routeModes)-1]
	routingParams := RoutingParams{
		Waypoint0: stringWaypoint0,
		Waypoint1: stringWaypoint1,
		AppID:     appID,
		AppCode:   appCode,
		Modes:     routeModes,
		Departure: "now",
	}
	return routingParams
}

// Route with given parameters.
func (s *RoutingService) Route(waypoint0 [2]float32, waypoint1 [2]float32, modes []Route) (*RoutingResponse, *http.Response, error) {
	routingParams := createRoutingParams(waypoint0, waypoint1, s.AppID, s.AppCode, modes)
	routes := new(RoutingResponse)
	resp, err := s.sling.New().Get("verify_credentials.json").QueryStruct(routingParams).ReceiveSuccess(routes)
	return routes, resp, err
}
