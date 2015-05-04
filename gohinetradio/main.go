// Package main is to play radio with VLC(mac).
//
/*
Install:

	go install github.com/toomore/gohinetradio/gohinetradio

Usage:

	gohinetradio

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/toomore/gohinetradio"
)

func main() {
	//PrintChannel()
	gohinetradio.GenList()

	//fmt.Println(GetRadioList(LISTPAGE))

	fmt.Println("輸入廣播頻道編號：")
	in := bufio.NewReader(os.Stdin)
	stdString, _ := in.ReadString('\n')
	radioNo := strings.Split(stdString, "\n")[0]
	if len(radioNo) > 0 {
		radioData, err := gohinetradio.GetURL(radioNo)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s %s\n%s\n",
				radioData.ChannelTitle, radioData.ProgramName, radioData.PlayRadio)
			if runtime.GOOS == "darwin" {
				exec.Command("/Applications/VLC.app/Contents/MacOS/VLC", radioData.PlayRadio).Start()
			}
		}
	}
}
