/* parses html links into structs 

Need to check out the solution on this one

*/
package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
	"strings"
)

// Link represents a link (<a href="...">)

type Link struct {
	Href string
	Text string 
}

// takes in html document and returns a slice of link parsed from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err 
	}
	nodes := linkNodes(doc)	// find link Nodes <a> in doc

	var links []Link
	// for each link node, build a link struct (get href and text)
	for _, node := range nodes {
		links = append(links, buildLink(node))
		// fmt.Println(node)
	}
	return links, nil 
}

// for each <a> tag Node, build a link struct 
func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = text(n)
	return ret
}

// get all text concatenated 
func text(n *html.Node) string {
	var ret string 
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	for c:= n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}
// DFS example 
// Get <a> tags 
func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c:= n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...) // append a slice 
	}
	return ret
}

func main() {
	// read in html
	htmlFile, _ := os.Open("ex4.html")
	// parse to nodes
	links, _ := Parse(htmlFile)
	fmt.Printf("%+v\n", links)

}