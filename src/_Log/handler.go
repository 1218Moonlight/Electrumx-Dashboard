package _Log

import (
	"github.com/andlabs/ui"
)

var writeChan = make(chan string)
var exitChan = make(chan int)

func Handler(logStatus *ui.MultilineEntry) {
	for {
		select {
		case wright := <-writeChan:
			logStatus.Append(wright+"\n")
		case <-exitChan:
			break
		}
	}
}
