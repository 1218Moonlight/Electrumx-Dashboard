package _andlabsUI

import (
	"github.com/andlabs/ui"
	"log"
	"github.com/sparrc/go-ping"
	"time"
)

var pingBool bool = false
var pingExit = make(chan int)

func electrumxInfo(status *ui.Button, elexStatus *ui.Label) {
	status.OnClicked(func(button *ui.Button) {
		if pingBool {
			status.SetText("pingBool true")
			elexStatus.SetText(Getinfo())
		} else {
			status.SetText("pingBool false")
			elexStatus.SetText("")
		}
	})
}

func serverPing(ip string, status *ui.Label) {
	log.Println("Start Ping", ip)
serverPingFor:
	for {
		select {
		case <-pingExit:
			status.SetText("")
			break serverPingFor
		default:
			pinger, err := ping.NewPinger(ip)
			pinger.SetPrivileged(true)
			if err != nil {
				panic(err)
			}
			pinger.Timeout = time.Duration(time.Second * 2)
			pinger.Count = 1
			pinger.Run()
			stats := pinger.Statistics()
			if len(stats.Rtts) == 0 {
				status.SetText("Server false")
			} else {
				status.SetText("Server true")
			}
			time.Sleep(time.Second * 3)
		}
	}
	log.Println("Stop Ping", ip)
}
