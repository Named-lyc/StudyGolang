package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1:%v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attribute := range n.Attr {
			if attribute.Key == "href" {
				links = append(links, attribute.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

// 递归版本
func visitRecursive(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	return visitRecursion(links, n.FirstChild)
}

func visitRecursion(links []string, child *html.Node) []string {
	//结束条件
	if child == nil {
		return links
	}
	links = visitRecursive(links, child)
	child = child.NextSibling
	return visitRecursion(links, child)
}
