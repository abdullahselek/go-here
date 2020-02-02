package here

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// WeatherService provides for HERE Weather api.
type WeatherService struct {
	sling *sling.Sling
}

// SevereWeatherAlertsParams parameters for severe alerts.
type SevereWeatherAlertsParams struct {
	Product string `url:"product" default0:"alerts"`
	Name    string `url:"name"`
	APIKey  string `url:"apiKey"`
}

// WeatherConditionsParams parameters for specified latitude and longitude
type WeatherConditionsParams struct {
	Product        string  `url:"product" default0:"observation"`
	Latitude       float32 `url:"latitude"`
	Longitude      float32 `url:"longitude"`
	OneObservation bool    `url:"oneobservation"`
	APIKey         string  `url:"apiKey"`
}

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

// WeatherConditionsResponse model for weather conditions specified by latitude and longitude.
type WeatherConditionsResponse struct {
	Observations struct {
		Location []struct {
			Observation []struct {
				Daylight          string    `json:"daylight"`
				Description       string    `json:"description"`
				SkyInfo           string    `json:"skyInfo"`
				SkyDescription    string    `json:"skyDescription"`
				Temperature       string    `json:"temperature"`
				TemperatureDesc   string    `json:"temperatureDesc"`
				Comfort           string    `json:"comfort"`
				HighTemperature   string    `json:"highTemperature"`
				LowTemperature    string    `json:"lowTemperature"`
				Humidity          string    `json:"humidity"`
				DewPoint          string    `json:"dewPoint"`
				Precipitation1H   string    `json:"precipitation1H"`
				Precipitation3H   string    `json:"precipitation3H"`
				Precipitation6H   string    `json:"precipitation6H"`
				Precipitation12H  string    `json:"precipitation12H"`
				Precipitation24H  string    `json:"precipitation24H"`
				PrecipitationDesc string    `json:"precipitationDesc"`
				AirInfo           string    `json:"airInfo"`
				AirDescription    string    `json:"airDescription"`
				WindSpeed         string    `json:"windSpeed"`
				WindDirection     string    `json:"windDirection"`
				WindDesc          string    `json:"windDesc"`
				WindDescShort     string    `json:"windDescShort"`
				BarometerPressure string    `json:"barometerPressure"`
				BarometerTrend    string    `json:"barometerTrend"`
				Visibility        string    `json:"visibility"`
				SnowCover         string    `json:"snowCover"`
				Icon              string    `json:"icon"`
				IconName          string    `json:"iconName"`
				IconLink          string    `json:"iconLink"`
				AgeMinutes        string    `json:"ageMinutes"`
				ActiveAlerts      string    `json:"activeAlerts"`
				Country           string    `json:"country"`
				State             string    `json:"state"`
				City              string    `json:"city"`
				Latitude          float64   `json:"latitude"`
				Longitude         float64   `json:"longitude"`
				Distance          float64   `json:"distance"`
				Elevation         int       `json:"elevation"`
				UtcTime           time.Time `json:"utcTime"`
			} `json:"observation"`
			Country   string  `json:"country"`
			State     string  `json:"state"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			Distance  float64 `json:"distance"`
			Timezone  int     `json:"timezone"`
		} `json:"location"`
	} `json:"observations"`
	FeedCreation time.Time `json:"feedCreation"`
	Metric       bool      `json:"metric"`
}

// newWeatherService returns a new WeatherService.
func newWeatherService(sling *sling.Sling) *WeatherService {
	return &WeatherService{
		sling: sling,
	}
}

// SevereWeatherAlerts fetches severe weather alert results.
func (s *WeatherService) SevereWeatherAlerts(params *SevereWeatherAlertsParams) (*SevereWeatherAlertsResponse, *http.Response, error) {
	weatherAlerts := new(SevereWeatherAlertsResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("report.json").QueryStruct(params).Receive(weatherAlerts, apiError)
	return weatherAlerts, resp, relevantError(err, *apiError)
}

// WeatherConditions fetches conditions for specified latitude and longitude.
func (s *WeatherService) WeatherConditions(params *WeatherConditionsParams) (*WeatherConditionsResponse, *http.Response, error) {
	weatherConditions := new(WeatherConditionsResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("report.json").QueryStruct(params).Receive(weatherConditions, apiError)
	return weatherConditions, resp, relevantError(err, *apiError)
}
