package plugins

import (
	"strings"

	"github.com/koffeinsource/go-URLextract/webpage"
	"github.com/koffeinsource/go-klogger"
)

// Image is called for links directly to images
func Image(i *webpage.Info, sourceURL string, contentType string, log klogger.KLogger) {
	if !(strings.Index(contentType, "image/") == 0) {
		return
	}

	log.Infof("Running Image plugin.")

	i.ImageURL = ""
	i.Caption = sourceURL[strings.LastIndex(sourceURL, "/")+1:]
	i.Description = "<img src=\"" + sourceURL + "\">"
}
