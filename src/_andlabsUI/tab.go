package _andlabsUI

import (
	"github.com/andlabs/ui"
	"time"
	"encoding/json"
	"strconv"
)

func logTab() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	status := ui.NewMultilineEntry()
	status.SetReadOnly(true)
	vbox.Append(status, true)

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

	getInfoVBox := ui.NewVerticalBox()
	getInfoHBox1 := ui.NewHorizontalBox()
	getInfoHBox2 := ui.NewHorizontalBox()
	getInfoVBox.Append(getInfoHBox1, false)
	getInfoVBox.Append(getInfoHBox2, false)

	infoClosing := ui.NewLabel("...")
	infoDaemon := ui.NewLabel("...")
	infoDaemonHeight := ui.NewLabel("...")
	infoDbHeight := ui.NewLabel("...")
	infoErrors := ui.NewLabel("...")
	infoGroups := ui.NewLabel("...")
	infoLogged := ui.NewLabel("...")
	infoPaused := ui.NewLabel("...")

	getInfoHBox1.Append(infoClosing, true)
	getInfoHBox1.Append(infoDaemon, true)
	getInfoHBox1.Append(infoDaemonHeight, true)
	getInfoHBox1.Append(infoDbHeight, true)
	getInfoHBox2.Append(infoErrors, true)
	getInfoHBox2.Append(infoGroups, true)
	getInfoHBox2.Append(infoLogged, true)
	getInfoHBox2.Append(infoPaused, true)

	getInfoGroup := ui.NewGroup("getInfo")
	getInfoGroup.SetMargined(true)
	getInfoGroup.SetChild(getInfoVBox)
	topVbox.Append(getInfoGroup, false)

	///

	elexInfo := elexGetinfo{}

	var pingU = pingUtil{"", pingStatus, pingMutex, 1}
	ipBtn.OnClicked(func(button *ui.Button) {
		pingU.url = ipEntry.Text()
		if !pingBool {
			boardLog.write("Ping " + pingU.url)

			json.Unmarshal(Getinfo(), &elexInfo)

			infoClosing.SetText("Closing : " + strconv.Itoa(elexInfo.Closing))
			infoDaemon.SetText("Daemon : " + elexInfo.Daemon)
			infoDaemonHeight.SetText("DaemonHeight : " + strconv.Itoa(elexInfo.Daemonheight))
			infoDbHeight.SetText("DB Height : " + strconv.Itoa(elexInfo.Dbheight))
			infoErrors.SetText("Errors : " + strconv.Itoa(elexInfo.Errors))
			infoGroups.SetText("Groups : " + strconv.Itoa(elexInfo.Groups))
			infoLogged.SetText("Logged : " + strconv.Itoa(elexInfo.Logged))
			infoPaused.SetText("Paused : " + strconv.Itoa(elexInfo.Paused))

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

			infoClosing.SetText("...")
			infoDaemon.SetText("...")
			infoDaemonHeight.SetText("...")
			infoDbHeight.SetText("...")
			infoErrors.SetText("...")
			infoGroups.SetText("...")
			infoLogged.SetText("...")
			infoPaused.SetText("...")
		}
	})

	return topVbox
}
