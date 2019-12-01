package here

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWeatherService(t *testing.T) {
	var httpClient = &http.Client{
		Timeout: time.Second * 15,
	}
	client := NewWeatherClient(httpClient)
	assert.NotNil(t, client)
}
