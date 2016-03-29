package webpage

// An Info is contains all data we got from a webpage
type Info struct {
	// Most likely a headline
	Caption string

	// The final URL of the requested page. May be different from the input
	URL string

	// Typically a small picture
	ImageURL string

	// Typically an abstract. But may also contain an image if the source was e.g. a comic.
	Description string

	// The HTML body of the queried webpage
	HTML string
}
