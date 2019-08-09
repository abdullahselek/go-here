package here

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// ReverseGeocodingService provides for HERE ReverseGeocoding api.
type ReverseGeocodingService struct {
	sling   *sling.Sling
	AppID   string
	AppCode string
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
func newReverseGeocodingService(sling *sling.Sling, appID string, appCode string) *ReverseGeocodingService {
	return &ReverseGeocodingService{
		sling:   sling,
		AppID:   appID,
		AppCode: appCode,
	}
}

// Returns prox as a formatted string.
func createProx(latlong [2]float32, diameter int) string {
	prox := fmt.Sprintf("%f,%f,%d", latlong[0], latlong[1], diameter)
	return prox
}

// AddressFromLocation returns address or addresses from given location.
func (s *ReverseGeocodingService) AddressFromLocation(latlong [2]float32, diameter int, mode ReverseGeocodingMode, maxResults int, gen int) (*GeocodingResponse, *http.Response, error) {
	addressFromLocationParameters := &AddressFromLocationParameters{Prox: createProx(latlong, diameter), Mode: ReverseGeocodingMode.String(mode), MaxResults: maxResults, Gen: gen, AppID: s.AppID, AppCode: s.AppCode}
	geocodingResponse := new(GeocodingResponse)
	resp, err := s.sling.New().Get("reversegeocode.json").QueryStruct(addressFromLocationParameters).ReceiveSuccess(geocodingResponse)
	return geocodingResponse, resp, err
}

// Landmarks returns details of landmarks near to a given latitude and longitude.
func (s *ReverseGeocodingService) Landmarks(latlong [2]float32, diameter int, gen int) (*GeocodingResponse, *http.Response, error) {
	landmarksParameters := &LandmarksParameters{Prox: createProx(latlong, diameter), Mode: ReverseGeocodingMode.String(RetrieveLandmarks), Gen: gen, AppID: s.AppID, AppCode: s.AppCode}
	geocodingResponse := new(GeocodingResponse)
	resp, err := s.sling.New().Get("reversegeocode.json").QueryStruct(landmarksParameters).ReceiveSuccess(geocodingResponse)
	return geocodingResponse, resp, err
}
