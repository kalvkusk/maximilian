package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
  fpid, _:= robotgo.FindIds("MicrosoftEdge")
  if len(fpid) > 0 {
  	robotgo.ActivePID(fpid[0])
  	robotgo.MaxWindow(fpid[0])
  }
}
