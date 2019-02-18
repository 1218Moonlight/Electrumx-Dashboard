package _andlabsUI

import (
	"github.com/andlabs/ui"
	"sync"
)

type pingUtil struct {
	url       string
	laber     *ui.Label
	mutex     *sync.Mutex
	exit      int
	elexLaber electrumxLaber
}

var pingChan = make(chan pingUtil)
var pingMutex = new(sync.Mutex)

func chanHandler() {
	for {
		select {
		case pingU := <-pingChan:
			pingU.mutex.Lock()
			serverPing(pingU.url, pingU.laber)
			electrumxGetinfo(pingU.url, pingU.elexLaber)
			pingU.mutex.Unlock()
		}
	}
}
