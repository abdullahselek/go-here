package here

import (
	"time"
)

// APIError represents a HERE API Error response
type APIError struct {
	ErrorType string `json:"_type"`
	Type      string `json:"type"`
	Subtype   string `json:"subtype"`
	Details   string `json:"details"`
	MetaInfo  struct {
		Timestamp           time.Time `json:"timestamp"`
		MapVersion          string    `json:"mapVersion"`
		ModuleVersion       string    `json:"moduleVersion"`
		InterfaceVersion    string    `json:"interfaceVersion"`
		AvailableMapVersion []string  `json:"availableMapVersion"`
	} `json:"metaInfo"`
}
