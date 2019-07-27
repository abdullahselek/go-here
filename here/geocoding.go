package here

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// GeocodingService provides for HERE Geocoding api.
type GeocodingService struct {
	sling   *sling.Sling
	AppID   string
	AppCode string
}

// Parameters by search text for Geocoding Service.
type SearchTextParameters struct {
	SearchText string `url:"searchtext"`
	AppID      string `url:"app_id"`
	AppCode    string `url:"app_code"`
	Gen        int    `url:"gen"`
}

// Parameters by address text within given boundingbox.
type AddressInBoundingBoxParameters struct {
	SearchText string `url:"searchtext"`
	MapView    string `url:"mapview"`
	Gen        int    `url:"gen"`
	AppID      string `url:"app_id"`
	AppCode    string `url:"app_code"`
}

// Parameters by partial address information.
type PartialAddressInformationParameters struct {
	HouseNumber int    `url:"housenumber"`
	Street      string `url:"street"`
	City        string `url:"city"`
	Country     string `url:"country"`
	Gen         int    `url:"gen"`
	AppID       string `url:"app_id"`
	AppCode     string `url:"app_code"`
}

// Response model for geocoding service.
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
func newGeocodingService(sling *sling.Sling, appID string, appCode string) *GeocodingService {
	return &GeocodingService{
		sling:   sling,
		AppID:   appID,
		AppCode: appCode,
	}
}

// Geocode by search text.
func (s *GeocodingService) Search(text string, gen int) (*GeocodingResponse, *http.Response, error) {
	searchTextParams := &SearchTextParameters{SearchText: text, AppID: s.AppID, AppCode: s.AppCode, Gen: gen}
	geocodingResponse := new(GeocodingResponse)
	resp, err := s.sling.New().Get("geocode.json").QueryStruct(searchTextParams).ReceiveSuccess(geocodingResponse)
	return geocodingResponse, resp, err
}

// Creates mapview parameter with given latitudes and longitudes.
func createMapView(latlong0 [2]float32, latlong1 [2]float32) string {
	waypoint0 := createWaypoint(latlong0)
	waypoint1 := createWaypoint(latlong1)
	mapView := fmt.Sprintf("%s;%s", waypoint0, waypoint1)
	return mapView
}

// Geocode by search text within given bounding box.
func (s *GeocodingService) AddressInBoundingBox(searchText string, latlong0 [2]float32, latlong1 [2]float32, gen int) (*GeocodingResponse, *http.Response, error) {
	searchTextParams := &AddressInBoundingBoxParameters{SearchText: searchText, MapView: createMapView(latlong0, latlong1), Gen: gen, AppID: s.AppID, AppCode: s.AppCode}
	geocodingResponse := new(GeocodingResponse)
	resp, err := s.sling.New().Get("geocode.json").QueryStruct(searchTextParams).ReceiveSuccess(geocodingResponse)
	return geocodingResponse, resp, err
}

// Request the latitude, longitude and details of an address based on partial address information.
func (s *GeocodingService) PartialAddressInformation(houseNumber int, street string, city string, country string, gen int) (*GeocodingResponse, *http.Response, error) {
	partialAddressInformationParameters := &PartialAddressInformationParameters{HouseNumber: houseNumber, Street: street, City: city, Country: country, Gen: gen, AppID: s.AppID, AppCode: s.AppCode}
	geocodingResponse := new(GeocodingResponse)
	resp, err := s.sling.New().Get("geocode.json").QueryStruct(partialAddressInformationParameters).ReceiveSuccess(geocodingResponse)
	return geocodingResponse, resp, err
}
