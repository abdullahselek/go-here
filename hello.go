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
	routingClient := here.NewRoutingClient(httpClient, "appId", "appCode")
	routes, httpResponse, err := routingClient.Routing.Route([2]float32{52.5160, 13.3779}, [2]float32{52.5206, 13.3862}, []here.Enum{here.RouteMode.Fastest, here.RouteMode.Car, here.RouteMode.TrafficDefault})
	fmt.Println(routes)
	fmt.Println(httpResponse)
	fmt.Println(err)

	geocodingClient := here.NewGeocodingClient(httpClient, "appId", "appCode")
	geocodingResponse, httpResponse, err := geocodingClient.Geocoding.AddressInBoundingBox("1 main", [2]float32{42.3902, -71.1293}, [2]float32{42.3312, -71.0228}, 9)
	fmt.Println(geocodingResponse)
	fmt.Println(httpResponse)
	fmt.Println(err)

	geocodingResponse, httpResponse, err = geocodingClient.Geocoding.PartialAddressInformation(425, "randolph", "chicago", "usa", 9)

	reverseGeocodingClient := here.NewReverseGeocodingClient(httpClient, "appID", "appCode")
	geocodingResponse, httpResponse, err = reverseGeocodingClient.ReverseGeocoding.AddressFromLocation([2]float32{42.3902, -71.1293}, 250, here.ReverseGeocodingMode.RetrieveAddresses, 1, 9)

	geocodingResponse, httpResponse, err = reverseGeocodingClient.ReverseGeocoding.Landmarks([2]float32{42.3902, -71.1293}, 1, 9)

	autocompleteGeocodingClient := here.NewAutocompleteGeocodingClient(httpClient, "appID", "appCode")
	autocompleteGeocodingResponse, httpResponse, err := autocompleteGeocodingClient.AutocompleteGeocoding.DetailsForSuggestion("Pariser 1 Berl")
}
