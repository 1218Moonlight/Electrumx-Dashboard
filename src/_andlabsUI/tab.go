package _andlabsUI

import (
	"github.com/andlabs/ui"
	"time"
	"strings"
)

func logTab() ui.Control {
	forAppend := func(status *ui.MultilineEntry, logs []string) {
		for i := range logs {
			status.Append(logs[i] + "\n")
		}
	}

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	logBtn := ui.NewButton("refresh")
	vbox.Append(logBtn, false)

	status := ui.NewMultilineEntry()
	status.SetReadOnly(true)
	vbox.Append(status, true)

	_, revert, err := readFile("log.txt")
	if !checkError(err, false) {
		forAppend(status, revert)
	}

	logBtn.OnClicked(func(button *ui.Button) {
		origin, _, err := readFile("log.txt")
		if checkError(err, false) {
			return
		}
		status.SetText("")
		forAppend(status, origin)
	})

	return vbox
}

func serverTab() ui.Control {
	topVbox := ui.NewVerticalBox()
	topVbox.SetPadded(true)

	///

	ipHbox := ui.NewHorizontalBox()
	ipHbox.SetPadded(true)

	ipEntry := ui.NewEntry()
	ipBtn := ui.NewButton("Connet")
	ipHbox.Append(ipEntry, true)
	ipHbox.Append(ipBtn, false)

	ipGroup := ui.NewGroup("Input")
	ipGroup.SetMargined(true)
	ipGroup.SetChild(ipHbox)
	topVbox.Append(ipGroup, false)

	///

	topVbox.Append(ui.NewHorizontalSeparator(), false)

	///

	pingStatus := ui.NewLabel("")

	pingGroup := ui.NewGroup("ping")
	pingGroup.SetMargined(true)
	pingGroup.SetChild(pingStatus)
	topVbox.Append(pingGroup, false)

	///

	getinfoLabel := ui.NewLabel("")

	getInfoGroup := ui.NewGroup("getInfo")
	getInfoGroup.SetMargined(true)
	getInfoGroup.SetChild(getinfoLabel)
	topVbox.Append(getInfoGroup, false)

	///

	sessionsLavel := ui.NewLabel("")

	sessionsGroup := ui.NewGroup("sessions")
	sessionsGroup.SetMargined(true)
	sessionsGroup.SetChild(sessionsLavel)
	topVbox.Append(sessionsGroup, false)

	///

	var pingU = pingUtil{"", pingStatus, pingMutex, 1,
		electrumxLaber{
			getinfo: getinfoLabel, sessions: sessionsLavel}}

	ipBtn.OnClicked(func(button *ui.Button) {
		if strings.Contains(ipEntry.Text(), urlHttp) {
			boardLog.writeError("Please do not enter http://")
			return
		} else if ipEntry.Text() == "" {
			boardLog.writeError("Enter url")
			return
		}
		pingU.url = ipEntry.Text()
		if !pingBool {
			boardLog.writeInfo("Ping " + pingU.url)

			go func() {
				for {
					if pingU.exit == 0 {
						pingU.laber.SetText("")
						pingU.exit = 1
						break
					}
					pingChan <- pingU
					time.Sleep(time.Second * 10)
				}
			}()
			ipBtn.SetText("Close")
			pingBool = true
		} else {
			pingU.exit = 0
			ipBtn.SetText("Connet")
			pingBool = false

			pingU.elexLaber.getinfo.SetText("")
			pingU.elexLaber.sessions.SetText("")
		}
	})

	return topVbox
}
