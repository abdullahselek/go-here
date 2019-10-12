package here

import (
	"fmt"
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

func (e APIError) Error() string {
	if len(e.Details) > 0 {
		return fmt.Sprintf("go-here: %v %v", e.ErrorType, e.Details)
	}
	return ""
}

// relevantError returns any non-nil http-related error (creating the request,
// getting the response, decoding) if any. If the decoded apiError is non-zero
// the apiError is returned. Otherwise, no errors occurred, returns nil.
func relevantError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}
	if len(apiError.ErrorType) > 0 {
		return apiError
	}
	return nil
}
