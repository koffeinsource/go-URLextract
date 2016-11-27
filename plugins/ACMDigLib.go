package plugins

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/koffeinsource/go-URLextract/webpage"
	"github.com/koffeinsource/go-klogger"
)

// ACMDigLib parses ACMs digital library
func ACMDigLib(i *webpage.Info, sourceURL string, doc *goquery.Document, log klogger.KLogger) {
	if !strings.Contains(sourceURL, "dl.acm.org/citation.cfm") {
		return
	}
	log.Infof("Running ACMDigLib plugin.")

	/*
	   Example meta tags
	   <meta name="citation_title" content="HermitCore: A Unikernel for Extreme Scale Computing"/>
	   <meta name="citation_doi" content="10.1145/2931088.2931093"/>
	   <meta name="citation_conference" content="Proceedings of the 6th International Workshop on Runtime and Operating Systems for Supercomputers"/>
	   <meta name="citation_publisher" content="ACM"/>
	   <meta name="citation_authors" content="Lankes, Stefan; Pickartz, Simon; Breitbart, Jens"/>
	   <meta name="citation_date" content="06/01/2016"/>
	   <meta name="citation_isbn" content="978-1-4503-4387-9"/>
	   <meta name="citation_firstpage" content="4"/>
	   <meta name="citation_abstract_html_url" content="http://dl.acm.org/citation.cfm?id=2931088.2931093"/>
	   <meta name="citation_pdf_url" content="http://dl.acm.org/ft_gateway.cfm?id=2931093&amp;type=pdf"/>
	*/

	selection := doc.Find("head meta[name=citation_title]")
	if selection.Size() == 0 {
		log.Errorf("ACMDigLib plugin found not citation_title. " + sourceURL)
	} else {
		m := htmlAttributeToMap(selection.Nodes[0].Attr)
		i.Caption = m["content"]
	}

	selection = doc.Find("head meta[name=citation_doi]")
	if selection.Size() == 0 {
		log.Errorf("ACMDigLib plugin found not citation_doi. " + sourceURL)
	} else {
		m := htmlAttributeToMap(selection.Nodes[0].Attr)
		i.URL = "http://dx.doi.org/" + m["content"]
	}

	selection = doc.Find("span#toHide6 div p")
	if len(selection.Nodes) == 0 {
		log.Errorf("ACMDigLib plugin found not abstract. " + sourceURL)
	} else {
		h, err := selection.Html()
		if err != nil {
			log.Errorf("ACMDigLib plugin error convertig abstract to html. " + sourceURL)
		} else {
			i.Description = h
		}
	}
}
