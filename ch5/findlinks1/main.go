package main

import "golang.org/x/net/html"

//The visit function traverses an HTML node tree, extracts the link /
//from the href attribute of each anchor element <a href='...'>,

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val) //getting the url
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
