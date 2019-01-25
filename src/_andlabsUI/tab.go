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
	topVerticalBox := ui.NewVerticalBox()
	topVerticalBox.SetPadded(true)

	inputHorizontalBox := ui.NewHorizontalBox()
	inputHorizontalBox.SetPadded(true)

	ipText := ui.NewEntry()
	ipBtn := ui.NewButton("Connet")
	inputHorizontalBox.Append(ipText, true)
	inputHorizontalBox.Append(ipBtn, false)

	inputGroup := ui.NewGroup("Input")
	inputGroup.SetMargined(true)
	inputGroup.SetChild(inputHorizontalBox)
	topVerticalBox.Append(inputGroup, false)

	///

	topVerticalBox.Append(ui.NewHorizontalSeparator(), false)

	///

	pingStatus := ui.NewLabel("")

	pingGroup := ui.NewGroup("ping")
	pingGroup.SetMargined(true)
	pingGroup.SetChild(pingStatus)
	topVerticalBox.Append(pingGroup, false)

	///

	getInfoVerticalBox := ui.NewVerticalBox()
	getInfoVerticalBox.SetPadded(true)
	getInfoHorizontalBox := ui.NewHorizontalBox()
	getInfoHorizontalBox.SetPadded(true)
	getInfoHorizontalBox1 := ui.NewHorizontalBox()
	getInfoHorizontalBox1.SetPadded(true)
	getInfoHorizontalBox2 := ui.NewHorizontalBox()
	getInfoHorizontalBox2.SetPadded(true)
	getInfoHorizontalBox3 := ui.NewHorizontalBox()
	getInfoHorizontalBox3.SetPadded(true)
	getInfoVerticalBox.Append(getInfoHorizontalBox, true)
	getInfoVerticalBox.Append(getInfoHorizontalBox1, true)
	getInfoVerticalBox.Append(getInfoHorizontalBox2, true)
	getInfoVerticalBox.Append(getInfoHorizontalBox3, true)

	closing := ui.NewLabel("...")
	daemon := ui.NewLabel("...")
	getInfoHorizontalBox.Append(closing, true)
	getInfoHorizontalBox.Append(daemon, true)

	daemon_height := ui.NewLabel("...")
	db_height := ui.NewLabel("...")
	getInfoHorizontalBox1.Append(daemon_height, true)
	getInfoHorizontalBox1.Append(db_height, true)

	errors := ui.NewLabel("...")
	groups := ui.NewLabel("...")
	getInfoHorizontalBox2.Append(errors, true)
	getInfoHorizontalBox2.Append(groups, true)

	logged := ui.NewLabel("...")
	paused := ui.NewLabel("...")
	getInfoHorizontalBox3.Append(logged, true)
	getInfoHorizontalBox3.Append(paused, true)

	daemonGroup := ui.NewGroup("getInfo")
	daemonGroup.SetMargined(true)
	daemonGroup.SetChild(getInfoVerticalBox)
	topVerticalBox.Append(daemonGroup, false)

	elexInfo := elexGetinfo{}

	var pingU = pingUtil{"", pingStatus, pingMutex, 1}
	ipBtn.OnClicked(func(button *ui.Button) {
		pingU.url = ipText.Text()
		if !pingBool {

			json.Unmarshal(Getinfo(), &elexInfo)

			closing.SetText("closing : " + strconv.Itoa(elexInfo.Closing))
			daemon.SetText("daemon : " + elexInfo.Daemon)
			daemon_height.SetText("daemon_height : " + strconv.Itoa(elexInfo.Daemon_height))
			db_height.SetText("db_height : " + strconv.Itoa(elexInfo.Db_height))
			errors.SetText("errors : " + strconv.Itoa(elexInfo.Errors))
			groups.SetText("groups : " + strconv.Itoa(elexInfo.Groups))
			logged.SetText("logged : " + strconv.Itoa(elexInfo.Logged))
			paused.SetText("paused : " + strconv.Itoa(elexInfo.Paused))

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
			closing.SetText("...")
			daemon.SetText("...")
			daemon_height.SetText("...")
			db_height.SetText("...")
			errors.SetText("...")
			groups.SetText("...")
			logged.SetText("...")
			paused.SetText("...")
		}
	})

	return topVerticalBox
}
