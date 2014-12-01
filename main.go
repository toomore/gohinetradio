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
	fmt.Println("[205] 中廣流行網 i like")
	fmt.Println("[206] 中廣音樂網i radio")
	fmt.Println("[232] 飛碟電台")
	fmt.Println("[222] HitFm聯播網 Taipei 北部")
	fmt.Println("[156] KISS RADIO 大眾廣播電台")
	fmt.Println("[308] KISS RADIO 網路音樂台")
	fmt.Println("[187] NEWS98新聞網")
	fmt.Println("[370] POP Radio 台北流行廣播電台")
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
