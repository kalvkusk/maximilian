package main

import (
	"github.com/go-vgo/robotgo"
  "fmt"

)

func maxim() {
	// find the process id by the process name
	fpid, err := robotgo.FindIds("Rainforest VM")
	if err == nil {
		fmt.Println("pids...", fpid)
		if len(fpid) > 0 {
			 robotgo.ActivePID(fpid[0])
			 hwnd := robotgo.GetHWND()
			 robotgo.MaxWindow()

		}
	}
}


func main() {
maxim()
}
