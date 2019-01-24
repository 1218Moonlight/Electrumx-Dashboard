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

	var pingU = pingUtil{"", pingStatus, pingMutex, 1}
	ipBtn.OnClicked(func(button *ui.Button) {
		pingU.url = ipText.Text()
		if !pingBool {

			json.Unmarshal(Getinfo(), &elexInfo)

			daemon_height.SetText(strconv.Itoa(elexInfo.Daemon_height))
			db_height.SetText(strconv.Itoa(elexInfo.Db_height))

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
			daemon_height.SetText("...")
			db_height.SetText("...")
		}
	})

	return vbox
}
