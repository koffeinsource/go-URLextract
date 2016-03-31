package plugins

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/koffeinsource/go-URLextract/webpage"
	"github.com/koffeinsource/go-imgur"
	"github.com/koffeinsource/go-klogger"
)

func createIMGTag(link string, gifv string, height int, width int) string {
	var ret string
	ret = "<img "
	ret += "height=\"" + strconv.Itoa(height) + "\" width=\"" + strconv.Itoa(width) + "\""
	ret += "src =\""
	if gifv != "" {
		ret += gifv
	} else {
		ret += link
	}
	// todo height
	ret += "\" /><br/>"
	return ret
}

func image(i *webpage.Info, v interface{}) {
	i.ImageURL = ""

	switch s := v.(type) {
	case *imgur.ImageInfo:
		i.Caption = s.Title
		i.Description = createIMGTag(s.Link, s.Gifv, s.Height, s.Width)
		if s.Section != "" {
			i.Description += "#" + s.Section + "<br/>"
		}
		if s.Description != "" {
			i.Description += s.Description
		}

	case *imgur.GalleryImageInfo:
		i.Caption = s.Title
		i.Description = createIMGTag(s.Link, s.Gifv, s.Height, s.Width)
		if s.Section != "" {
			i.Description += "#" + s.Section + "<br/>"
		}
		if s.Description != "" {
			i.Description += s.Description
		}
	default:
		panic("Passed invalid type to image function in go-URLextract imgur plugin")
	}
}

func album(i *webpage.Info, v interface{}) {
	i.ImageURL = ""

	switch s := v.(type) {
	case *imgur.AlbumInfo:
		i.Caption = "[ALBUM] " + s.Title
		i.Description = ""
		fmt.Println(s.ImagesCount)
		if s.ImagesCount > 0 {
			i.Description = createIMGTag(s.Images[0].Link, s.Images[0].Gifv, s.Images[0].Height, s.Images[0].Width)
			i.Description += "…"
		}
		if s.Description != "" {
			i.Description += s.Description
		}
		i.URL = s.Link

	case *imgur.GalleryAlbumInfo:
		i.Caption = "[ALBUM] " + s.Title
		i.Description = ""
		if s.ImagesCount > 0 {
			i.Description = createIMGTag(s.Images[0].Link, s.Images[0].Gifv, s.Images[0].Height, s.Images[0].Width)
			i.Description += "…"
		}
		if s.Description != "" {
			i.Description += s.Description
		}
		i.URL = s.Link

	default:
		panic("Passed invalid type to album function in go-URLextract imgur plugin")
	}
}

// Imgurl extract all images from an imgurl album
func Imgurl(i *webpage.Info, sourceURL string, httpClient *http.Client, log klogger.KLogger, imgurClientID string) {
	if !strings.Contains(sourceURL, "imgur.com/") {
		return
	}

	log.Infof("Running imgurl plugin.")

	client := new(imgur.Client)
	client.HTTPClient = httpClient
	client.Log = log
	client.ImgurClientID = imgurClientID

	gi, status, err := client.GetInfoFromURL(sourceURL)

	if err != nil {
		log.Errorf("Error using go-imgur. Status: %v. Error: %v", status, err)
		return
	}
	if status > 399 {
		log.Errorf("Error using go-imgur. Status: %v. Error: nil", status)
		return
	}

	if gi.Image != nil {
		image(i, gi.Image)
		return
	}
	if gi.GImage != nil {
		image(i, gi.GImage)
		return
	}

	if gi.Album != nil {
		log.Debugf("1\n")
		album(i, gi.Album)
		return
	}
	if gi.GAlbum != nil {
		log.Debugf("2\n")
		album(i, gi.GAlbum)
		return
	}

	panic("Unknown type used in go-URLextract imgur plugin")
}
