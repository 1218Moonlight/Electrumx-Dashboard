package _andlabsUI

import (
	"github.com/andlabs/ui"
	"time"
	"encoding/json"
	"strconv"
)

func logTab() ui.Control {
	boardLog.write("Setting logTab")
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	status := ui.NewMultilineEntry()
	status.SetReadOnly(true)
	vbox.Append(status, true)

	return vbox
}

func serverTab() ui.Control {
	boardLog.write("Setting serverTab")
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	ipText := ui.NewEntry()
	ipBtn := ui.NewButton("Connet")
	hbox.Append(ipText, true)
	hbox.Append(ipBtn, false)

	inputGroup := ui.NewGroup("Input")
	inputGroup.SetMargined(true)
	inputGroup.SetChild(hbox)
	vbox.Append(inputGroup, false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	pingStatus := ui.NewLabel("")

	pingGroup := ui.NewGroup("ping")
	pingGroup.SetMargined(true)
	pingGroup.SetChild(pingStatus)
	vbox.Append(pingGroup, false)

	var pingU = pingUtil{"", pingStatus, pingMutex, 1}
	ipBtn.OnClicked(func(button *ui.Button) {
		pingU.url = ipText.Text()
		if !pingBool {
			go func() {
				for {
					if pingU.exit == 0 {
						pingU.laber.SetText("")
						pingU.exit = 1
						break
					}
					pingChan <- pingU
					time.Sleep(time.Second * 3)
				}
			}()
			ipBtn.SetText("Close")
			pingBool = true
		} else {
			pingU.exit = 0
			ipBtn.SetText("Connet")
			pingBool = false
		}
	})

	return vbox
}

func electrumxTab() ui.Control {
	boardLog.write("Setting electrumxTab")
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	status := ui.NewButton("getElectrumxInfo")
	vbox.Append(status, false)

	groupVbox := ui.NewVerticalBox()
	groupVbox.SetPadded(true)

	elexStatus := ui.NewLabel("")
	groupVbox.Append(elexStatus, true)

	daemonGroup := ui.NewGroup("daemon_height")
	daemonGroup.SetMargined(true)
	daemon_height := ui.NewLabel("...")
	daemonGroup.SetChild(daemon_height)
	vbox.Append(daemonGroup, false)

	dbGroup := ui.NewGroup("db_height")
	daemonGroup.SetMargined(true)
	db_height := ui.NewLabel("...")
	dbGroup.SetChild(db_height)
	vbox.Append(dbGroup, false)

	elexInfo := elexGetinfo{}
	status.OnClicked(func(button *ui.Button) {
		if pingBool {
			status.SetText("pingBool true")

			json.Unmarshal(Getinfo(), &elexInfo)

			daemon_height.SetText(strconv.Itoa(elexInfo.Daemon_height))
			db_height.SetText(strconv.Itoa(elexInfo.Db_height))

		} else {
			status.SetText("pingBool false")
			elexStatus.SetText("")
		}
	})



	return vbox
}
