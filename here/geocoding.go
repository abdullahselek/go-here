package here

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// GeocodingService provides for HERE Geocoding api.
type GeocodingService struct {
	sling *sling.Sling
}

// SearchTextParameters parameters by search text for Geocoding Service.
type SearchTextParameters struct {
	SearchText string `url:"searchtext"`
	APIKey     string `url:"apiKey"`
	Gen        int    `url:"gen"`
}

// AddressInBoundingBoxParameters parameters by address text within given boundingbox.
type AddressInBoundingBoxParameters struct {
	SearchText string `url:"searchtext"`
	MapView    string `url:"mapview"`
	Gen        int    `url:"gen"`
	APIKey     string `url:"apiKey"`
}

// PartialAddressInformationParameters parameters by partial address information.
type PartialAddressInformationParameters struct {
	HouseNumber int    `url:"housenumber"`
	Street      string `url:"street"`
	City        string `url:"city"`
	Country     string `url:"country"`
	Gen         int    `url:"gen"`
	APIKey      string `url:"apiKey"`
}

// GeocodingResponse response model for geocoding service.
type GeocodingResponse struct {
	Response struct {
		MetaInfo struct {
			Timestamp string `json:"Timestamp"`
		} `json:"MetaInfo"`
		View []struct {
			Type   string `json:"_type"`
			ViewID int    `json:"ViewId"`
			Result []struct {
				Relevance    int    `json:"Relevance"`
				MatchLevel   string `json:"MatchLevel"`
				MatchQuality struct {
					State       int       `json:"State"`
					City        int       `json:"City"`
					Street      []float64 `json:"Street"`
					HouseNumber int       `json:"HouseNumber"`
				} `json:"MatchQuality"`
				MatchType string `json:"MatchType"`
				Location  struct {
					LocationID      string `json:"LocationId"`
					LocationType    string `json:"LocationType"`
					DisplayPosition struct {
						Latitude  float64 `json:"Latitude"`
						Longitude float64 `json:"Longitude"`
					} `json:"DisplayPosition"`
					NavigationPosition []struct {
						Latitude  float64 `json:"Latitude"`
						Longitude float64 `json:"Longitude"`
					} `json:"NavigationPosition"`
					MapView struct {
						TopLeft struct {
							Latitude  float64 `json:"Latitude"`
							Longitude float64 `json:"Longitude"`
						} `json:"TopLeft"`
						BottomRight struct {
							Latitude  float64 `json:"Latitude"`
							Longitude float64 `json:"Longitude"`
						} `json:"BottomRight"`
					} `json:"MapView"`
					Address struct {
						Label          string `json:"Label"`
						Country        string `json:"Country"`
						State          string `json:"State"`
						County         string `json:"County"`
						City           string `json:"City"`
						District       string `json:"District"`
						Street         string `json:"Street"`
						HouseNumber    string `json:"HouseNumber"`
						PostalCode     string `json:"PostalCode"`
						AdditionalData []struct {
							Value string `json:"value"`
							Key   string `json:"key"`
						} `json:"AdditionalData"`
					} `json:"Address"`
				} `json:"Location"`
			} `json:"Result"`
		} `json:"View"`
	} `json:"Response"`
}

// newGeocodingService returns a new GeocodingService.
func newGeocodingService(sling *sling.Sling) *GeocodingService {
	return &GeocodingService{
		sling: sling,
	}
}

// Search for geocode by text.
func (s *GeocodingService) Search(params *SearchTextParameters) (*GeocodingResponse, *http.Response, error) {
	geocodingResponse := new(GeocodingResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("geocode.json").QueryStruct(params).Receive(geocodingResponse, apiError)
	return geocodingResponse, resp, relevantError(err, *apiError)
}

// CreateMapView Creates mapview parameter with given latitudes and longitudes.
func (s *GeocodingService) CreateMapView(latlong0 [2]float32, latlong1 [2]float32) string {
	waypoint0 := createWaypoint(latlong0)
	waypoint1 := createWaypoint(latlong1)
	mapView := fmt.Sprintf("%s;%s", waypoint0, waypoint1)
	return mapView
}

// AddressInBoundingBox by search text within given bounding box.
func (s *GeocodingService) AddressInBoundingBox(params *AddressInBoundingBoxParameters) (*GeocodingResponse, *http.Response, error) {
	geocodingResponse := new(GeocodingResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("geocode.json").QueryStruct(params).Receive(geocodingResponse, apiError)
	return geocodingResponse, resp, relevantError(err, *apiError)
}

// PartialAddressInformation requests the latitude, longitude and details of an address based on partial address information.
func (s *GeocodingService) PartialAddressInformation(params *PartialAddressInformationParameters) (*GeocodingResponse, *http.Response, error) {
	geocodingResponse := new(GeocodingResponse)
	resp, err := s.sling.New().Get("geocode.json").QueryStruct(params).ReceiveSuccess(geocodingResponse)
	return geocodingResponse, resp, err
}
