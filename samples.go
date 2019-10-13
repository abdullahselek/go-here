package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/abdullahselek/go-here/here"
)

func main() {
	var httpClient = &http.Client{
		Timeout: time.Second * 15,
	}
	routingClient := here.NewRoutingClient(httpClient)
	routingParams := routingClient.Routing.CreateRoutingParams([2]float32{52.5160, 13.3779}, [2]float32{52.5206, 13.3862}, "appID", "appCode", []here.Enum{here.RouteMode.Fastest, here.RouteMode.Car, here.RouteMode.TrafficDefault})
	routes, httpResponse, err := routingClient.Routing.Route(&routingParams)
	fmt.Println(routes)
	fmt.Println(httpResponse)
	fmt.Println(err)

	geocodingClient := here.NewGeocodingClient(httpClient)
	addressBoundingBoxParams := &here.AddressInBoundingBoxParameters{SearchText: "1 main", MapView: geocodingClient.Geocoding.CreateMapView([2]float32{42.3902, -71.1293}, [2]float32{42.3312, -71.0228}), Gen: 9, AppID: "appID", AppCode: "appCode"}
	geocodingResponse, httpResponse, err := geocodingClient.Geocoding.AddressInBoundingBox(addressBoundingBoxParams)
	fmt.Println(geocodingResponse)
	fmt.Println(httpResponse)
	fmt.Println(err)

	partialAddressInformationParams := &here.PartialAddressInformationParameters{HouseNumber: 425, Street: "randolph", City: "chicago", Country: "usa", Gen: 9, AppID: "appID", AppCode: "appCode"}
	geocodingResponse, httpResponse, err = geocodingClient.Geocoding.PartialAddressInformation(partialAddressInformationParams)

	reverseGeocodingClient := here.NewReverseGeocodingClient(httpClient)
	locationParameters := reverseGeocodingClient.ReverseGeocoding.CreateAddressFromLocationParameters([2]float32{42.3902, -71.1293}, 250, here.ReverseGeocodingMode.RetrieveAddresses, 1, 9, "appID", "appCode")
	geocodingResponse, httpResponse, err = reverseGeocodingClient.ReverseGeocoding.AddressFromLocation(&locationParameters)

	landmarkParameters := reverseGeocodingClient.ReverseGeocoding.CreateLandmarksParameters([2]float32{42.3902, -71.1293}, 1, 9, "appID", "appCode")
	geocodingResponse, httpResponse, err = reverseGeocodingClient.ReverseGeocoding.Landmarks(&landmarkParameters)

	autocompleteGeocodingClient := here.NewAutocompleteGeocodingClient(httpClient)
	suggestionsParameters := autocompleteGeocodingClient.AutocompleteGeocoding.CreateDetailsForSuggestionParameters("Pariser 1 Berl", "appID", "appCode")
	autocompleteGeocodingResponse, httpResponse, err := autocompleteGeocodingClient.AutocompleteGeocoding.DetailsForSuggestion(&suggestionsParameters)
	fmt.Println(autocompleteGeocodingResponse)
	fmt.Println(httpResponse)
	fmt.Println(err)
}
