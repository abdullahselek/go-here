package here

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// ReverseGeocodingService provides for HERE ReverseGeocoding api.
type ReverseGeocodingService struct {
	sling *sling.Sling
}

// AddressFromLocationParameters used for fetching address from location.
type AddressFromLocationParameters struct {
	Prox       string `url:"prox"`
	Mode       string `url:"mode"`
	MaxResults int    `url:"maxresults"`
	Gen        int    `url:"gen"`
	AppID      string `url:"app_id"`
	AppCode    string `url:"app_code"`
}

// LandmarksParameters used for fetching address from location.
type LandmarksParameters struct {
	Prox    string `url:"prox"`
	Mode    string `url:"mode"`
	Gen     int    `url:"gen"`
	AppID   string `url:"app_id"`
	AppCode string `url:"app_code"`
}

// newReverseGeocodingService returns a new GeocodingService.
func newReverseGeocodingService(sling *sling.Sling) *ReverseGeocodingService {
	return &ReverseGeocodingService{
		sling: sling,
	}
}

// Returns prox as a formatted string.
func createProx(latlong [2]float32, diameter int) string {
	prox := fmt.Sprintf("%f,%f,%d", latlong[0], latlong[1], diameter)
	return prox
}

// CreateAddressFromLocationParameters creates AddressFromLocationParameters used for fetching address from location.
func (s *ReverseGeocodingService) CreateAddressFromLocationParameters(latlong [2]float32, diameter int, mode Enum, maxResults int, gen int, appID string, appCode string) AddressFromLocationParameters {
	addressFromLocationParameters := AddressFromLocationParameters{Prox: createProx(latlong, diameter), Mode: Enum.ValueOfReverseGeocodingMode(mode), MaxResults: maxResults, Gen: gen, AppID: appID, AppCode: appCode}
	return addressFromLocationParameters
}

// CreateLandmarksParameters creates LandmarksParameters used for fetching address from location.
func (s *ReverseGeocodingService) CreateLandmarksParameters(latlong [2]float32, diameter int, gen int, appID string, appCode string) LandmarksParameters {
	landmarksParameters := LandmarksParameters{Prox: createProx(latlong, diameter), Mode: Enum.ValueOfReverseGeocodingMode(ReverseGeocodingMode.RetrieveLandmarks), Gen: gen, AppID: appID, AppCode: appCode}
	return landmarksParameters
}

// AddressFromLocation returns address or addresses from given location.
func (s *ReverseGeocodingService) AddressFromLocation(params *AddressFromLocationParameters) (*GeocodingResponse, *http.Response, error) {
	geocodingResponse := new(GeocodingResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("reversegeocode.json").QueryStruct(params).Receive(geocodingResponse, apiError)
	return geocodingResponse, resp, relevantError(err, *apiError)
}

// Landmarks returns details of landmarks near to a given latitude and longitude.
func (s *ReverseGeocodingService) Landmarks(params *LandmarksParameters) (*GeocodingResponse, *http.Response, error) {
	geocodingResponse := new(GeocodingResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("reversegeocode.json").QueryStruct(params).Receive(geocodingResponse, apiError)
	return geocodingResponse, resp, relevantError(err, *apiError)
}
