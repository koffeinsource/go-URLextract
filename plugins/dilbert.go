package plugins

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/asaskevich/govalidator"
	"github.com/koffeinsource/go-URLextract/webpage"
	"github.com/koffeinsource/go-klogger"

	"golang.org/x/net/html"
)

// Dilbert extracts the comic from a dilbert page
func Dilbert(i *webpage.Info, sourceURL string, doc *goquery.Document, log klogger.KLogger) {
	if !(strings.Contains(sourceURL, "feed.dilbert.com/") ||
		strings.Contains(sourceURL, "dilbert.com/strips/")) {
		return
	}

	log.Infof("Running Dilbert plugin.")

	selection := doc.Find(".STR_Image").Find("img")

	if len(selection.Nodes) == 0 {
		log.Errorf("Dilbert plugin found no .STR_Image. " + sourceURL)
		return
	}

	if len(selection.Nodes) > 1 {
		log.Infof("Dilbert plugin found >1 .STR_Image. " + sourceURL)
	}

	e := selection.Nodes[0]
	if e.Type == html.ElementNode && e.Data == "img" {
		m := htmlAttributeToMap(e.Attr)
		u := ""
		if !strings.Contains(m["src"], "://dilbert.com") {
			u += "https://dilbert.com"
		}

		u += m["src"]
		if govalidator.IsRequestURL(u) {
			i.Description = createImageHTMLTag(u)
		} else {
			log.Errorf("Dilbert plugin invalid url. " + u)
		}

	} else {
		log.Errorf("Dilbert plugin no image tag where we expect one.")
	}

	i.ImageURL = ""
	i.Caption = "Dilbert"
}
