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
	counts := map[string]int{}
	for tag, count := range count(counts, doc) {
		fmt.Printf("<%s>\t%d\n", tag, count)
	}
}

// 循环版本
func count(counts map[string]int, doc *html.Node) map[string]int {
	if doc.Type == html.ElementNode {
		counts[doc.Data]++
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		count(counts, c)
	}
	return counts
}

// 递归版本:
//func countRecursive(counts map[string]int, doc *html.Node) map[string]int {
//	if doc.Type == html.ElementNode {
//		counts[doc.Data]++
//	}
//	return countRecursion(counts, doc.FirstChild)
//}
//
//func countRecursion(counts map[string]int, c *html.Node) map[string]int {
//	if c == nil {
//		return counts
//	}
//	counts = countRecursive(counts, c)
//	c = c.NextSibling
//	return countRecursion(counts, c)
//}
