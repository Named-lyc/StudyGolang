package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline : %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, doc *html.Node) {
	if doc.Type == html.ElementNode {
		stack = append(stack, doc.Data)
		fmt.Println(stack)
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}

}
