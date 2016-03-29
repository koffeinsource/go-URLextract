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
	webaddr := flag.String("url", "", "The URL to be parsed.")
	flag.Parse()

	// Check if URL was passed at command line
	if *webaddr == "" {
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
	var c URLextract.Config
	c.HTTPClient = new(http.Client)
	c.Log = new(klogger.CLILogger)
	fmt.Print(URLextract.Extract(*webaddr, c))
}
