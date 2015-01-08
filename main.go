package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/tabwriter"
)

const (
	PlayURL string = "http://hichannel.hinet.net/radio/play.do?id=%s"
)

type RadioData struct {
	ChannelTitle   string
	PlayRadio      string
	ProgramName    string
	ChannelCollect bool
}

func GetUrl(url_no string) (r RadioData) {
	url := fmt.Sprintf(PlayURL, url_no)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	jsonData := json.NewDecoder(bytes.NewReader(data))
	jsonData.Decode(&r)
	return
}

func GetList() {
	resp, _ := http.Get("http://hichannel.hinet.net/radio/mobile/index.do?id=207")
	defer resp.Body.Close()
	html := new(bytes.Buffer)
	html.ReadFrom(resp.Body)
	reg := regexp.MustCompile(`<div class="stationName">(.+)</div>[\s]+<div class="list"><a href="#" onclick="getInfo\('list','([\d]+)'\);return false;"></a></div>`)
	url_string := reg.FindAllStringSubmatch(html.String(), -1)
	url_string = append(url_string, []string{"", "中廣新聞網", "207"})
	//fmt.Println(url_string)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	var output string
	for no, data := range url_string {
		output += fmt.Sprintf("%d. [%s] %s\t", no+1, data[2], data[1])
		if (no+1)%3 == 0 {
			fmt.Fprintln(w, output)
			output = ""
		}
	}
	fmt.Fprintln(w, output)
	w.Flush()
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
	//fmt.Println(GetUrl("207"))
	//PrintChannel()
	//GetList()

	in := bufio.NewReader(os.Stdin)
	std_string, _ := in.ReadString('\n')
	radio_url := GetUrl(strings.Split(std_string, "\n")[0])
	fmt.Println(radio_url)
	//exec.Command("/Applications/VLC.app/Contents/MacOS/VLC", radio_url).Start()
}
