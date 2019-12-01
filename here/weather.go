package here

import (
	"time"

	"github.com/dghubble/sling"
)

// WeatherService provides for HERE Weather api.
type WeatherService struct {
	sling *sling.Sling
}

// SevereWeatherAlertsResponse model for sever alerts.
// SevereWeatherAlertsResponse model for severe alerts.
type SevereWeatherAlertsResponse struct {
	Alerts struct {
		Alerts []struct {
			TimeSegment []struct {
				Value           string `json:"value"`
				Segment         string `json:"segment"`
				OtherAttributes struct {
				} `json:"otherAttributes"`
				DayOfWeek string `json:"day_of_week"`
			} `json:"timeSegment"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"alerts"`
		Country   string  `json:"country"`
		State     string  `json:"state"`
		City      string  `json:"city"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Timezone  int     `json:"timezone"`
	} `json:"alerts"`
	FeedCreation time.Time `json:"feedCreation"`
	Metric       bool      `json:"metric"`
}

// newWeatherService returns a new WeatherService.
func newWeatherService(sling *sling.Sling) *WeatherService {
	return &WeatherService{
		sling: sling,
	}
}
