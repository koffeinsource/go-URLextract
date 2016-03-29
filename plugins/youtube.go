package plugins

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/koffeinsource/go-URLextract/webpage"
	"github.com/koffeinsource/go-klogger"
)

// Youtube extracts the video from an youtube url
func Youtube(i *webpage.Info, sourceURL string, doc *goquery.Document, log klogger.KLogger) {
	if !strings.Contains(sourceURL, "www.youtube.com") {
		return
	}

	log.Infof("Running Youtube plugin.")

	// update title

	videoIDstart := strings.Index(i.URL, "v=")
	if videoIDstart == -1 {
		log.Infof("Youtube plugin found no video ID. " + sourceURL)
		return
	}
	videoIDstart += 2 // ID is after 'v='
	videoID := i.URL[videoIDstart:]
	i.Description += "<br/><br/><br/><iframe width=\"560\" height=\"315\" src=\"https://www.youtube.com/embed/"
	i.Description += videoID
	i.Description += "\" frameborder=\"0\" allowfullscreen></iframe>"

	i.ImageURL = ""
}
