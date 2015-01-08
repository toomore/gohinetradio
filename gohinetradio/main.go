package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/toomore/gohinetradio"
)

func main() {
	//PrintChannel()
	gohinetradio.GenList()

	//fmt.Println(GetRadioList(LISTPAGE))

	in := bufio.NewReader(os.Stdin)
	stdString, _ := in.ReadString('\n')
	radioData := gohinetradio.GetURL(strings.Split(stdString, "\n")[0])
	fmt.Printf("%s %s\n%s\n",
		radioData.ChannelTitle, radioData.ProgramName, radioData.PlayRadio)
	exec.Command("/Applications/VLC.app/Contents/MacOS/VLC", radioData.PlayRadio).Start()
}
