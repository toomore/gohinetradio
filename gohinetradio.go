package gohinetradio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"sync"
	"text/tabwriter"
)

const (
	PLAYURL  string = "http://hichannel.hinet.net/radio/play.do?id=%s"
	LISTURL  string = "http://hichannel.hinet.net/radio/channelList.do?radioType=&freqType=&freq=&area=&pN=%s"
	LISTPAGE int    = 4
)

type RadioData struct {
	ChannelTitle   string `json:"channel_title"`
	PlayRadio      string `json:"playRadio"`
	ProgramName    string `json:"programName"`
	ChannelCollect bool   `json:"channel_collect"`
}

func GetURL(url_no string) (r RadioData) {
	resp, _ := http.Get(fmt.Sprintf(PLAYURL, url_no))
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	jsonData := json.NewDecoder(bytes.NewReader(data))
	jsonData.Decode(&r)
	return
}

type RadioListData struct {
	PageNo   int
	PageSize int
	List     []RadioListDatas
}

type RadioListDatas struct {
	ChannelImage string `json:"channel_image"`
	ChannelTitle string `json:"channel_title"`
	RadioType    string `json:"radio_type"`
	IsChannel    bool   `json:"isChannel"`
	ProgramName  string `json:"program_name"`
	ChannelId    string `json:"channel_id"`
}

func getRadioPageList(page int) (r RadioListData) {
	resp, _ := http.Get(fmt.Sprintf(LISTURL, page))
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	jsonData := json.NewDecoder(bytes.NewReader(data))
	jsonData.Decode(&r)
	return
}

func GetRadioList(total int) (result []RadioListDatas) {
	queue := make(chan RadioListData)
	var wg sync.WaitGroup
	wg.Add(int(LISTPAGE))
	for i := 1; i <= total; i++ {
		go func(i int) {
			defer wg.Done()
			runtime.Gosched()
			queue <- getRadioPageList(i)
		}(i)
	}
	go func() {
		defer wg.Done()
		for v := range queue {
			for _, data := range v.List {
				result = append(result, data)
			}
		}
	}()
	wg.Wait()
	return
}

func GenList() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	var output string
	for no, data := range GetRadioList(LISTPAGE) {
		output += fmt.Sprintf("%d. [%s] %s\t", no+1, data.ChannelId, data.ChannelTitle)
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