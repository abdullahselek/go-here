package main

import (
	"fmt"
	"go-here/here"
	"net/http"
	"time"
)

func main() {
	var httpClient = &http.Client{
		Timeout: time.Second * 15,
	}
	routingClient := here.NewRoutingClient(httpClient)
	routingParams := routingClient.Routing.CreateRoutingParams([2]float32{52.5160, 13.3779}, [2]float32{52.5206, 13.3862}, "appID", []here.Enum{here.RouteMode.Fastest, here.RouteMode.Car, here.RouteMode.TrafficDefault})
	routes, httpResponse, err := routingClient.Routing.Route(&routingParams)
	fmt.Println(routes)
	fmt.Println(httpResponse)
	fmt.Println(err)

	geocodingClient := here.NewGeocodingClient(httpClient)
	addressBoundingBoxParams := here.AddressInBoundingBoxParameters{SearchText: "1 main", MapView: geocodingClient.Geocoding.CreateMapView([2]float32{42.3902, -71.1293}, [2]float32{42.3312, -71.0228}), Gen: 9, APIKey: "appKey"}
	geocodingResponse, httpResponse, err := geocodingClient.Geocoding.AddressInBoundingBox(&addressBoundingBoxParams)
	fmt.Println(geocodingResponse)
	fmt.Println(httpResponse)
	fmt.Println(err)

	partialAddressInformationParams := here.PartialAddressInformationParameters{HouseNumber: 425, Street: "randolph", City: "chicago", Country: "usa", Gen: 9, APIKey: "apiKey"}
	geocodingResponse, httpResponse, err = geocodingClient.Geocoding.PartialAddressInformation(&partialAddressInformationParams)

	reverseGeocodingClient := here.NewReverseGeocodingClient(httpClient)
	locationParameters := reverseGeocodingClient.ReverseGeocoding.CreateAddressFromLocationParameters([2]float32{42.3902, -71.1293}, 250, here.ReverseGeocodingMode.RetrieveAddresses, 1, 9, "apiKey")
	geocodingResponse, httpResponse, err = reverseGeocodingClient.ReverseGeocoding.AddressFromLocation(&locationParameters)

	landmarkParameters := reverseGeocodingClient.ReverseGeocoding.CreateLandmarksParameters([2]float32{42.3902, -71.1293}, 1, 9, "apiKey")
	geocodingResponse, httpResponse, err = reverseGeocodingClient.ReverseGeocoding.Landmarks(&landmarkParameters)

	autocompleteGeocodingClient := here.NewAutocompleteGeocodingClient(httpClient)
	suggestionsParameters := autocompleteGeocodingClient.AutocompleteGeocoding.CreateDetailsForSuggestionParameters("Pariser 1 Berl", "apiKey")
	autocompleteGeocodingResponse, httpResponse, err := autocompleteGeocodingClient.AutocompleteGeocoding.DetailsForSuggestion(&suggestionsParameters)
	fmt.Println(autocompleteGeocodingResponse)
	fmt.Println(httpResponse)
	fmt.Println(err)

	fleetTelematicsClient := here.NewFleetTelematicsClient(httpClient)
	start := here.DestinationParams{Coordinates: [2]float32{-25.643787, -49.158607}, Text: "Start"}
	end := here.DestinationParams{Coordinates: [2]float32{-20.778591, -51.591198}, Text: "End"}
	destinations := []here.DestinationParams{{Coordinates: [2]float32{-25.644764, -49.158558}, Text: "Destination1"}, {Coordinates: [2]float32{-22.98319, -49.903282}, Text: "Destination2"}, {Coordinates: [2]float32{-20.778555, -51.591164}, Text: "Destination3"}}
	fleetTelematicsParams := fleetTelematicsClient.FleetTelematics.CreateFleetTelematicsParams(start, end, destinations, "apiKey", []here.Enum{here.RouteMode.Fastest, here.RouteMode.Truck})
	routeSequence, httpResponse, err := fleetTelematicsClient.FleetTelematics.FindSequence(&fleetTelematicsParams)
	fmt.Println(routeSequence)
	fmt.Println(httpResponse)
	fmt.Println(err)
}
