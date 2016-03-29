package plugins

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/koffeinsource/go-URLextract/webpage"
	"github.com/koffeinsource/go-klogger"

	"golang.org/x/net/html"
)

// Gfycat extacts the animation from a gfycat page
func Gfycat(i *webpage.Info, sourceURL string, doc *goquery.Document, log klogger.KLogger) {
	if !strings.Contains(sourceURL, "gfycat.com/") {
		return
	}

	log.Infof("Running Gfycat plugin.")

	i.ImageURL = ""

	selection := doc.Find(".gfyVid")

	if len(selection.Nodes) == 0 {
		log.Errorf("Gfycat plugin found no .gfyVid. " + sourceURL)
		return
	}
	if len(selection.Nodes) > 1 {
		log.Infof("Gfycat plugin found >1 .gfyVid. " + sourceURL)
	}
	buf := new(bytes.Buffer)
	err := html.Render(buf, selection.Nodes[0])
	if err != nil {
		log.Errorf("Gfycat plugin error while rendering. " + sourceURL + "- " + err.Error())
		return
	}

	i.Description = buf.String()

	selection = doc.Find(".gfyTitle")
	if len(selection.Nodes) == 0 {
		log.Infof("Gfycat plugin found no .gfyTitle. " + sourceURL)
		return
	}
	if len(selection.Nodes) > 1 {
		log.Infof("Gfycat plugin found >1 .gfyTitle. " + sourceURL)
	}
	if len(selection.Nodes) != 0 && selection.Nodes[0].FirstChild != nil {
		i.Caption = selection.Nodes[0].FirstChild.Data
	} else {
		i.Caption = "Gfycat"
	}

}
