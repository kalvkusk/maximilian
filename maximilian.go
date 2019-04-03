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
			 robotgo.MaxWindow(fpid[0])

		}
	}
}


func main() {
maxim()
}
