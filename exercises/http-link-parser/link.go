package link

import "io"

// Link is a representation of link (<a href="...">) in HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take an HTML document and will return a slice of links parsed
// from it.
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
