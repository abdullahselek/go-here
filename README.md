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
routingClient := here.NewRoutingClient(httpClient, "appId", "appCode")
routes, httpResponse, err := routingClient.Routing.Route([2]float32{52.5160, 13.3779}, [2]float32{52.5206, 13.3862}, []here.Enum{here.RouteMode.Fastest, here.RouteMode.Car, here.RouteMode.TrafficDefault})

// Finding Address in boundingbox
geocodingClient := here.NewGeocodingClient(httpClient, "appId", "appCode")
geocodingResponse, httpResponse, err := geocodingClient.Geocoding.AddressInBoundingBox("1 main", [2]float32{42.3902, -71.1293}, [2]float32{42.3312, -71.0228}, 9)

// Partial address information
geocodingResponse, httpResponse, err := geocodingClient.Geocoding.PartialAddressInformation(425, "randolph", "chicago", "usa", 9)

// Reverse geocoding for address details
reverseGeocodingClient := here.NewReverseGeocodingClient(httpClient, "appID", "appCode")
geocodingResponse, httpResponse, err := reverseGeocodingClient.ReverseGeocoding.AddressFromLocation([2]float32{42.3902, -71.1293}, 250, here.ReverseGeocodingMode.RetrieveAddresses, 1, 9)

// Reverse geocoding for landmark details
geocodingResponse, httpResponse, err := reverseGeocodingClient.ReverseGeocoding.Landmarks([2]float32{42.3902, -71.1293}, 1, 9)

// Complete location details
autocompleteGeocodingClient := here.NewAutocompleteGeocodingClient(httpClient, "appID", "appCode")
autocompleteGeocodingResponse, httpResponse, err := autocompleteGeocodingClient.AutocompleteGeocoding.DetailsForSuggestion("Pariser 1 Berl")
```

## Roadmap

* Add new clients for other endpoints.
* Use parameter structs on functions.

## License

[MIT License](https://github.com/abdullahselek/go-here/blob/master/LICENSE)
