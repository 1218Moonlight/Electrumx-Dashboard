package _andlabsUI

import (
	"github.com/sparrc/go-ping"
	"time"
	"log"
)

var pingBool bool = false

func serverPing(pingU pingUtil) {
	pinger, err := ping.NewPinger(pingU.url)
	pinger.SetPrivileged(true)
	if err != nil {
		panic(err)
	}
	pinger.Timeout = time.Duration(time.Second * 1)
	pinger.Count = 1
	pinger.Run()
	stats := pinger.Statistics()
	if len(stats.Rtts) == 0 {
		log.Println("Server false")
		pingU.laber.SetText("Server false")
	} else {
		log.Println("Server true")
		pingU.laber.SetText("Server true")
	}
}
