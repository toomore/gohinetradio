// Package main is to play radio with VLC(mac).
//
/*
Install:

	go install github.com/toomore/gohinetradio/gohinetradio

Usage:

	gohinetradio [RadioNo]

*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/toomore/gohinetradio"
)

func play(radioNo string) {
	if radioData, err := gohinetradio.GetURL(radioNo); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s %s\n%s\n",
			radioData.ChannelTitle, radioData.ProgramName, radioData.PlayRadio)
		if runtime.GOOS == "darwin" {
			if err := exec.Command("/Applications/VLC.app/Contents/MacOS/VLC", radioData.PlayRadio).Start(); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	flag.Parse()
	var args = flag.Args()

	if len(args) == 1 {
		play(args[0])
	} else {
		//PrintChannel()
		gohinetradio.GenList()
		//fmt.Println(GetRadioList(LISTPAGE))
		fmt.Println("輸入廣播頻道編號：")
		in := bufio.NewReader(os.Stdin)
		stdString, _ := in.ReadString('\n')
		radioNo := strings.Split(stdString, "\n")[0]
		if len(radioNo) > 0 {
			play(radioNo)
		}
	}
}
