package _andlabsUI

import (
	"github.com/andlabs/ui"
	"sync"
)

type pingUtil struct {
	url   string
	laber *ui.Label
	mutex *sync.Mutex
	exit  int
}

var pingChan = make(chan pingUtil)
var pingMutex = new(sync.Mutex)

func chanHandler() {
	for {
		select {
		case pingU := <-pingChan:
			pingU.mutex.Lock()
			serverPing(pingU)
			pingU.mutex.Unlock()
		}
	}
}
