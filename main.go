package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func GetUrl(url_no string) (replace_string string) {
	url := fmt.Sprintf("http://hichannel.hinet.net/radio/mobile/index.do?id=%s", url_no)
	//resp, _ := http.Get("http://hichannel.hinet.net/radio/mobile/index.do?id=232")
	//resp, _ := http.Get("http://hichannel.hinet.net/radio/mobile/index.do?id=207")
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	html := new(bytes.Buffer)
	html.ReadFrom(resp.Body)
	//fmt.Println(html.String())
	reg := regexp.MustCompile(`var url = '([\S]+)'`)
	url_string := reg.FindAllStringSubmatch(html.String(), -1)
	//fmt.Println(url_string[0][1])

	replace := regexp.MustCompile(`\\\/`)
	replace_string = replace.ReplaceAllString(url_string[0][1], `/`)
	//fmt.Println(replace_string)
	return
}

func PrintChannel() {
	fmt.Println("[207] 中廣新聞網")
	fmt.Println("[232] 飛碟電台")
}

func main() {
	//fmt.Println("----- test open -----\r\n")
	//exec.Command("open", "-a", "firefox").Run()

	//fmt.Println(GetUrl("207"))

	PrintChannel()

	in := bufio.NewReader(os.Stdin)
	std_string, _ := in.ReadString('\n')
	fmt.Println(GetUrl(strings.Split(std_string, "\n")[0]))
}
