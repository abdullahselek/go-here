package here

import (
	"github.com/dghubble/sling"
)

// WeatherService provides for HERE Weather api.
type WeatherService struct {
	sling *sling.Sling
}

// newWeatherService returns a new WeatherService.
func newWeatherService(sling *sling.Sling) *WeatherService {
	return &WeatherService{
		sling: sling,
	}
}
