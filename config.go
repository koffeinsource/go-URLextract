package URLextract

import (
	"net/http"

	"github.com/koffeinsource/go-klogger"
)

// Config used to configure go-URLextract
type Config struct {
	HTTPClient *http.Client
	Log        klogger.KLogger
	AmazonAdID string
}
