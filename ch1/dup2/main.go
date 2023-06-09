package main

import (
	"bufio"
	"fmt"
	"os"
)

type theFile struct {
	Count     int
	Filenames []string
}

func main() {
	counts := make(map[string]*theFile)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.Count > 1 {
			fmt.Printf("%d %s  %v\n", n.Count, n.Filenames, line)
		}
	}
}

func countLines(f *os.File, counts map[string]*theFile) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()
		_, ok := counts[key]
		if ok {
			counts[key].Count++
			counts[key].Filenames = append(counts[key].Filenames, f.Name())
		} else {
			counts[key] = new(theFile)
			counts[key].Count = 1
			counts[key].Filenames = append(counts[key].Filenames, f.Name())
		}
	}
}
