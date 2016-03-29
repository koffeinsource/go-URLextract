package plugins

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/asaskevich/govalidator"
	"github.com/koffeinsource/go-URLextract/webpage"
	"github.com/koffeinsource/go-klogger"
)

// Littlegamers extract a comic from a littlegamers page
func Littlegamers(i *webpage.Info, sourceURL string, doc *goquery.Document, log klogger.KLogger) {
	if !strings.Contains(sourceURL, "www.little-gamers.com") {
		return
	}

	log.Infof("Running little-gamers plugin.")

	selection := doc.Find("img#comic")

	if len(selection.Nodes) == 0 {
		log.Infof("little-gamers plugin found no img#comic. " + sourceURL)
	} else {
		if len(selection.Nodes) > 1 {
			// that should actually never happen
			log.Errorf("little-gamers plugin found >1 img#comic. ??? " + sourceURL)
		}
		m := htmlAttributeToMap(selection.Nodes[0].Attr)

		if govalidator.IsRequestURL(m["src"]) {
			i.Description = createImageHTMLTag(m["src"])
			i.ImageURL = ""
		} else {
			log.Errorf("little-gamers plugin invalid url. " + m["src"])
		}
	}

}
