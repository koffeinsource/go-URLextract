package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/koffeinsource/go-URLextract"
	"github.com/koffeinsource/go-klogger"
)

func main() {
	imgurClientID := flag.String("imgurid", "", "Your imgur client id. REQUIRED!")
	webaddr := flag.String("url", "", "The URL to be parsed.")
	flag.Parse()

	// Check if URL was passed at command line
	if *webaddr == "" || *imgurClientID == "" {
		flag.PrintDefaults()
		return
	}

	fmt.Print("Validating URL...")
	// Is it a valid URL?
	if !govalidator.IsRequestURL(*webaddr) {
		fmt.Println("invalid!")
		return
	}
	fmt.Println(" success!")
	var c URLextract.Client
	c.HTTPClient = new(http.Client)
	c.Log = new(klogger.CLILogger)
	c.ImgurClientID = *imgurClientID

	wi, err := c.Extract(*webaddr)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}
	fmt.Printf("Caption: %v \nDescription: %v \nImage URL: %v \nURL: %v\n", wi.Caption, wi.Description, wi.ImageURL, wi.URL)
}
