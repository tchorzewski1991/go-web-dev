package link

import (
	"golang.org/x/net/html"
	"io"
)

// Link is a representation of link (<a href="...">) in HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take an HTML document and will return a slice of links parsed
// from it.
func Parse(r io.Reader) ([]Link, error) {

	doc, err := html.Parse(r)

	if err != nil {
		panic(err)
	}

	nodes := findLinks(doc)

	var links []Link

	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

// findLinks() function will take parsed HTML and will
// use Depth First Search (DFS) algorithm to find and return
// slice of all link nodes.
func findLinks(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, findLinks(c)...)
	}

	return ret
}

// findText() function will take parsed link node and will
// use Depth First Search (DFS) algorithm to find and return
// string with all texts from link node.
func findText(n *html.Node) string {
	var ret string

	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += findText(c) + " "
	}

	return ret
}

// buildLink() function will extract value of href attribute
// and all texts from link node and will return Link struct.
func buildLink(n *html.Node) Link {
	var ret Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}

	ret.Text = findText(n)

	return  ret
}
