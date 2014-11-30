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
	//fmt.Println(html.String())
	reg := regexp.MustCompile(`var url = '([\S]+)'`)
	url_string := reg.FindAllStringSubmatch(html.String(), -1)
	//fmt.Println(url_string[0][1])

	replace := regexp.MustCompile(`\\\/`)
	replace_string := replace.ReplaceAllString(url_string[0][1], `/`)
	fmt.Println(replace_string)

}
