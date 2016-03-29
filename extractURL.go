package URLextract

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/asaskevich/govalidator"
	"github.com/koffeinsource/go-URLextract/plugins"
	"github.com/koffeinsource/go-URLextract/webpage"
	"golang.org/x/net/html/charset"
)

// Extract extracts all information from URL
func Extract(sourceURL string, c Config) (webpage.Info, error) {

	// Create return value with default values
	returnee := webpage.Info{
		Caption: sourceURL,
		URL:     sourceURL,
	}

	// Check if the URL is valid
	if !govalidator.IsRequestURL(sourceURL) {
		errReturn := fmt.Errorf("Invalid URL: %v", sourceURL)
		c.Log.Errorf(errReturn.Error())
		return returnee, errReturn
	}

	contentType, body, err := getURL(sourceURL, c)
	if err != nil {
		return returnee, err
	}

	//  log.Infof(contentType)
	switch {
	case strings.Contains(contentType, "image/"):
		plugins.Image(&returnee, sourceURL, contentType, c.Log)
	case strings.Contains(contentType, "text/html"):

		var doc *goquery.Document

		charsetReader, err := charset.NewReader(bytes.NewReader(body), contentType)
		if err == nil {
			doc, err = goquery.NewDocumentFromReader(charsetReader)
		} else {
			doc, err = goquery.NewDocumentFromReader(bytes.NewReader(body))
		}

		if err != nil {
			c.Log.Errorf("Problem parsing body. " + sourceURL + " - " + err.Error())
			return returnee, err
		}

		// Make sure to call this one first
		plugins.DefaultHTML(&returnee, sourceURL, doc, c.Log)

		plugins.Amazon(&returnee, sourceURL, doc, c.Log, c.AmazonAdID)

		plugins.Imgurl(&returnee, sourceURL, doc, c.Log)
		plugins.Gfycat(&returnee, sourceURL, doc, c.Log)

		plugins.Fefe(&returnee, sourceURL, doc, c.Log)

		plugins.Youtube(&returnee, sourceURL, doc, c.Log)
		plugins.Vimeo(&returnee, sourceURL, doc, c.Log)

		plugins.Dilbert(&returnee, sourceURL, doc, c.Log)
		plugins.Garfield(&returnee, sourceURL, doc, c.Log)
		plugins.Xkcd(&returnee, sourceURL, doc, c.Log)
		plugins.Littlegamers(&returnee, sourceURL, doc, c.Log)

		plugins.Pastebin(&returnee, sourceURL, doc, c.Log)
	default:
	}

	return returnee, nil
}
