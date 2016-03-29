package plugins

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/koffeinsource/go-URLextract/webpage"
	"github.com/koffeinsource/go-klogger"
)

// Pastebin extracts the content from a pastbin page
func Pastebin(i *webpage.Info, sourceURL string, doc *goquery.Document, log klogger.KLogger) {
	if !strings.Contains(sourceURL, "pastebin.com/") {
		return
	}

	// TODO replace the logic below with a query to http://pastebin.com/raw.php?i=
	log.Infof("Running pastebin plugin.")

	selection := doc.Find("#paste_code")

	if len(selection.Nodes) == 0 {
		log.Infof("Pastebin plugin found no #paste_code. " + sourceURL)
	} else {
		if len(selection.Nodes) > 1 {
			log.Infof("Pastebin plugin found >1 #paste_code. " + sourceURL)
		}

		str, err := selection.Html()
		if err != nil {
			log.Infof("Error when creating html in pastebin plugin: ")
			return
		}
		str = strings.Replace(str, "\n", "<br />\n", -1)
		i.Description = str
		i.ImageURL = ""
	}

}
