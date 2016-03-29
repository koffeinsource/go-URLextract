package plugins

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/asaskevich/govalidator"
	"github.com/koffeinsource/go-URLextract/webpage"
	"github.com/koffeinsource/go-klogger"

	"golang.org/x/net/html"
)

// Amazon webpage plugin
func Amazon(i *webpage.Info, sourceURL string, doc *goquery.Document, log klogger.KLogger, amazonAdID string) {
	if !strings.Contains(sourceURL, "www.amazon.") {
		return
	}

	log.Infof("Running Amazon plugin.")

	// find picture
	{
		selection := doc.Find("#landingImage")
		if len(selection.Nodes) == 0 {
			log.Infof("Amazon plugin found no #landingImage. " + sourceURL)
		} else {
			if len(selection.Nodes) > 1 {
				log.Infof("Amazon plugin found >1 #landingImage. " + sourceURL)
			}
			for _, e := range selection.Nodes {
				if e.Type == html.ElementNode && e.Data == "img" {
					m := htmlAttributeToMap(e.Attr)
					if govalidator.IsRequestURL(m["data-old-hires"]) {
						i.ImageURL = m["data-old-hires"]
					} else {
						log.Infof("Amazon plugin imgURL invalid. " + m["data-old-hires"])
					}
				}
			}
		}
	}

	// update url to contain tag
	{
		// This is our tag. We should make it configurable
		urlExtension := ""
		if amazonAdID != "" {
			urlExtension = "tag=" + amazonAdID
		}
		start := strings.Index(i.URL, "tag=")
		if start != -1 {
			end := strings.Index(i.URL[start+1:], "&") + start + 1
			i.URL = i.URL[:start] + i.URL[end:]
		}

		if strings.Index(i.URL, "?") == -1 {
			i.URL += "?" + urlExtension
		} else {
			i.URL += "&" + urlExtension
		}
	}

	// update title
	{
		selection := doc.Find("#productTitle")
		if len(selection.Nodes) == 0 {
			log.Infof("Amazon plugin found no #productTitle. " + sourceURL)
		} else {
			if len(selection.Nodes) > 1 {
				log.Infof("Amazon plugin found >1 #productTitle. " + sourceURL)
			}
			for _, e := range selection.Nodes {
				if e.Type == html.ElementNode && e.Data == "span" {
					i.Caption = e.FirstChild.Data
				}
			}
		}
	}

	// Store HTML for the search
	{
		if s, err := doc.Find(".a-container").Html(); err != nil {
			log.Errorf("Error finding .a-container in HTML: %v", err)
		} else {
			i.HTML = s
		}
	}
}
