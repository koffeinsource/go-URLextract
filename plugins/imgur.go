package plugins

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/asaskevich/govalidator"
	"github.com/koffeinsource/go-URLextract/webpage"
	"github.com/koffeinsource/go-klogger"
)

// Imgurl extract all images from an imgurl album
func Imgurl(i *webpage.Info, sourceURL string, doc *goquery.Document, log klogger.KLogger) {
	if !strings.Contains(sourceURL, "imgur.com/") {
		return
	}

	log.Infof("Running imgurl plugin.")

	selection := doc.Find("meta[property*='og']")

	if selection.Length() != 0 {
		set := make(map[string]bool)

		i.Description = ""
		i.ImageURL = ""

		for _, e := range selection.Nodes {
			m := htmlAttributeToMap(e.Attr)

			if m["property"] == "og:image" {
				if !govalidator.IsRequestURL(m["content"]) {
					log.Infof("Invalid url in og:image. " + sourceURL)
					continue
				}
				if _, in := set[m["content"]]; !in {
					i.Description += "<img src =\""
					temp := strings.Replace(m["content"], "http://", "https://", 1)
					i.Description += temp
					i.Description += "\" /><br/>"
					set[m["content"]] = true
				}
			}
		}

	}
}
