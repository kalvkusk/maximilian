package main

import (
	"github.com/go-vgo/robotgo"
	"flag"
  "fmt"
	"os"
)

func main() {

textPtr := flag.String("name", "none", "Process name (Required)")
flag.Parse()

if *textPtr == "" {
	        flag.PrintDefaults()
	        os.Exit(1)
		    }

fpid, _:= robotgo.FindIds(*textPtr)
 fmt.Printf("%d ", fpid)
  if len(fpid) > 0 {
		fmt.Printf("Maximizing")
  	robotgo.MaxWindow(fpid[0])
  } else {
		fmt.Printf("%s process was not found!", *textPtr)
	}
}
