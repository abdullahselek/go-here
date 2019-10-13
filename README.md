# go-here [![codecov](https://codecov.io/gh/abdullahselek/go-here/branch/master/graph/badge.svg)](https://codecov.io/gh/abdullahselek/go-here) [![GoDoc](https://godoc.org/github.com/abdullahselek/go-here/here?status.svg)](https://godoc.org/github.com/abdullahselek/go-here/here) [![Go Report Card](https://goreportcard.com/badge/abdullahselek/go-here)](https://goreportcard.com/report/abdullahselek/go-here)

| Build Type | Status  |
| ---        | ---     |
| Linux | [![Build Status](https://travis-ci.org/abdullahselek/go-here.svg?branch=master)](https://travis-ci.org/abdullahselek/go-here) |
| Windows | [![Build status](https://ci.appveyor.com/api/projects/status/d9g1ehueqau9s57h?svg=true)](https://ci.appveyor.com/project/abdullahselek/go-here) |

**go-here** is a Go client library for the [HERE API](https://developer.here.com). [HERE](https://www.here.com) provides location based services. HERE exposes [rest APIs](https://developer.here.com/develop/rest-apis) and this library is intended to make it even easier for Go programmers to use. Check the usage section or try the examples to see how to access the HERE API.

### Features

* HERE REST API:
    * Routing
    * Geocoding
    * Reverse Geocoding
    * Geocoding Autocomplete

Will add rest of the apis in time and all contributions are welcome.

## Install

    go get github.com/abdullahselek/go-here/here

## Documentation

Read [GoDoc](https://godoc.org/github.com/abdullahselek/go-here/here)

## Usage

The `here` package provides a `Client` for accessing the HERE API and each API service requires an AppID and AppKey. Here are some example requests.

```go
var httpClient = &http.Client{
    Timeout: time.Second * 15,
}
// Routing client
routingClient := here.NewRoutingClient(httpClient)
routingParams := routingClient.Routing.CreateRoutingParams([2]float32{52.5160, 13.3779}, [2]float32{52.5206, 13.3862}, "appID", "appCode", []here.Enum{here.RouteMode.Fastest, here.RouteMode.Car, here.RouteMode.TrafficDefault})
routes, httpResponse, err := routingClient.Routing.Route(&routingParams)

// Finding Address in boundingbox
geocodingClient := here.NewGeocodingClient(httpClient)
addressBoundingBoxParams := here.AddressInBoundingBoxParameters{SearchText: "1 main", MapView: geocodingClient.Geocoding.CreateMapView([2]float32{42.3902, -71.1293}, [2]float32{42.3312, -71.0228}), Gen: 9, AppID: "appID", AppCode: "appCode"}
geocodingResponse, httpResponse, err := geocodingClient.Geocoding.AddressInBoundingBox(&addressBoundingBoxParams)

// Partial address information
partialAddressInformationParams := here.PartialAddressInformationParameters{HouseNumber: 425, Street: "randolph", City: "chicago", Country: "usa", Gen: 9, AppID: "appID", AppCode: "appCode"}
geocodingResponse, httpResponse, err = geocodingClient.Geocoding.PartialAddressInformation(&partialAddressInformationParams)

// Reverse geocoding for address details
reverseGeocodingClient := here.NewReverseGeocodingClient(httpClient)
locationParameters := reverseGeocodingClient.ReverseGeocoding.CreateAddressFromLocationParameters([2]float32{42.3902, -71.1293}, 250, here.ReverseGeocodingMode.RetrieveAddresses, 1, 9, "appID", "appCode")
geocodingResponse, httpResponse, err := reverseGeocodingClient.ReverseGeocoding.AddressFromLocation(&locationParameters)

// Reverse geocoding for landmark details
landmarkParameters := reverseGeocodingClient.ReverseGeocoding.CreateLandmarksParameters([2]float32{42.3902, -71.1293}, 1, 9, "appID", "appCode")
geocodingResponse, httpResponse, err := reverseGeocodingClient.ReverseGeocoding.Landmarks(&landmarkParameters)

// Complete location details
autocompleteGeocodingClient := here.NewAutocompleteGeocodingClient(httpClient)
suggestionsParameters := autocompleteGeocodingClient.AutocompleteGeocoding.CreateDetailsForSuggestionParameters("Pariser 1 Berl", "appID", "appCode")
autocompleteGeocodingResponse, httpResponse, err := autocompleteGeocodingClient.AutocompleteGeocoding.DetailsForSuggestion(&suggestionsParameters)
```

## Roadmap

* Add new clients for other endpoints.
* ~~Use parameter structs on functions.~~

## License

[MIT License](https://github.com/abdullahselek/go-here/blob/master/LICENSE)
