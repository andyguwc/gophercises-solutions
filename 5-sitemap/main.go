/* 
Sitemap is a map of all the pages within a sepcific domain. Used by search engines

Need to implement a BFS

First get the html 
Get links on the page and add domains and filter for the right ones
If starts with / vs. if starts with http 

go build . && ./5-sitemap --input "https://www.github.com" 
*/

package main 

import (
	// "encoding/xml"
	"fmt"
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/andyguwc/go-course/gophercises/5-sitemap/link"

)

// take an input url
// var URLs []string

// declare flags 
var InputURL = flag.String("input", "https://gophercises.com", "input domain url to parse")
var maxDepth = flag.Int("depth", 3, "the maximum number of links") // can avoid cycles 

func init() {
	flag.Parse()
}

func main() {

	res := bfs(*InputURL, *maxDepth)
	for _, v := range res {
		fmt.Println(v)
	}

}

func getCleanedLinks(input string) []string {
	resp, err := http.Get(input)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	res := getLinks(resp.Body, input)
	return res
}

// get links from a res.Body
func getLinks(reader io.Reader, base string) []string {
	var URLs []string 
	parsedLinks, _ := link.Parse(reader)

	for _, link := range parsedLinks {
		if strings.HasPrefix(link.Href, "/") {
			newlink := base + link.Href
			URLs = append(URLs, newlink)
		} else {
			URLs = append(URLs, link.Href)
		}
	}
	return filter(base, URLs)
}

// filter the returned URLs 
func filter(base string, links []string) []string {
	parsedBase , _ := url.Parse(base)
	parsedHost := parsedBase.Host
	var filteredURLs []string
	for _, link := range links {
		u, _ := url.Parse(link)
		if u.Host == parsedHost {
			filteredURLs = append(filteredURLs, link)
		} else {
			continue	
		}
	}

	return filteredURLs
}

// implement BFS to traverse the urls
func bfs(urlStr string, maxDepth int) []string {
	// keep track of all urls visited (struct{} so no value, sort of like a set)
	seen := make(map[string]struct{})

	// q will be the url we need to call getCleanedLinks on
	q := make(map[string]struct{})

	// nq will be links we haven't seen yet. Once we have nq we'll assign q as nq and keep looping
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}

	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		if len(q) == 0 {
			break
		} 
		for currURL, _ := range q {
			if _, ok := seen[currURL]; ok {
				continue 
			} else {
				seen[currURL] = struct{}{}
				for _, link := range getCleanedLinks(currURL) {
					nq[link] = struct{}{}
				}
			}
		}
	}

	ret := make([]string,0,len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}

	return ret

}

