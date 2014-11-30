package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://hichannel.hinet.net/radio/mobile/index.do?id=232")
	defer resp.Body.Close()
	html := new(bytes.Buffer)
	html.ReadFrom(resp.Body)
	fmt.Println(html)
}
