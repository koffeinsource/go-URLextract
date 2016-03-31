package URLextract

import (
	"net/http"

	"github.com/koffeinsource/go-klogger"
)

// Client used to access go-URLextract
type Client struct {
	HTTPClient *http.Client
	Log        klogger.KLogger
	AmazonAdID string
}
