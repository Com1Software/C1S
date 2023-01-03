package c1s

import (
	"fmt"
	"os"
	"runtime"

	term "github.com/nsf/termbox-go"
)

type CommandControl struct {
	Status     string
	ReturnCode string
	Command    string
}

func Cmd(cmdctl CommandControl) CommandControl {
	cbuf := ""
	cmd := cmdctl.Command
	fmt.Printf("C1Script\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	fmt.Println("Command Line Mode")
	fmt.Printf(". %s", cmd)
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyArrowUp:
				fmt.Println("Arrow Up pressed")
			case term.KeyArrowDown:
				fmt.Println("Arrow Down pressed")
			case term.KeyEnter:
				fmt.Printf("\n.")
				cmd = cbuf
				cbuf = ""
			default:
				s := rune(ev.Ch)
				cbuf = cbuf + string(s)
				fmt.Printf("%s", string(s))
			}
		}
		switch {
		case cmd == "end":
			fmt.Println("end C1Script")
			return cmdctl
		case cmd == "exit":
			fmt.Println("\nExit C1Script")
			os.Exit(2)
		case cmd == "?":
			fmt.Println("\nHelp")

			fmt.Printf("\n.")
		}
	}
}
func Run(cmd string) {
	fmt.Println("[Test Hello Run] " + cmd)
}
