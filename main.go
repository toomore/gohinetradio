package main

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
)

func main() {
	resp, _ := http.Get("http://hichannel.hinet.net/radio/mobile/index.do?id=232")
	defer resp.Body.Close()
	html := new(bytes.Buffer)
	html.ReadFrom(resp.Body)
	fmt.Println(html.String())
	reg := regexp.MustCompile(`var url = '([\S]+)'`)
	bb := reg.FindAllStringSubmatch(html.String(), -1)
	fmt.Println(bb[0][1])

}
