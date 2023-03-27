package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Fetch prints the content found at a URL.
// 获取url中的内容
func main() {

	for _, url := range os.Args[1:] {
		//如果输入的url前缀没有http://
		if !strings.HasPrefix(url, "http://") {
			// url = "http://"+url
			url = strings.Join([]string{"http://", url}, "")
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		//b, err := io.ReadAll(resp.Body)
		b, err := io.Copy(os.Stdout, resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s :%v\n", url, err)
		}
		fmt.Printf("%s", b)
	}
}
