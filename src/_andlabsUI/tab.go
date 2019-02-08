package _andlabsUI

import (
	"github.com/andlabs/ui"
	"time"
	"strings"
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

	var pingU = pingUtil{"", pingStatus, pingMutex, 1,
		electrumxLaber{
			closing: infoClosing, daemon: infoDaemon, daemonHeight: infoDaemonHeight, dbHeight: infoDbHeight,
			errors:  infoErrors, groups:infoGroups, logged: infoLogged, paused: infoPaused}}

	ipBtn.OnClicked(func(button *ui.Button) {
		if strings.Contains(ipEntry.Text(), urlHttp){
			boardLog.write("[err] Please do not enter http://")
			return
		}
		pingU.url = ipEntry.Text()
		if !pingBool {
			boardLog.write("Ping " + pingU.url)

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

			pingU.elexLaber.closing.SetText("...")
			pingU.elexLaber.daemon.SetText("...")
			pingU.elexLaber.daemonHeight.SetText("...")
			pingU.elexLaber.dbHeight.SetText("...")
			pingU.elexLaber.errors.SetText("...")
			pingU.elexLaber.groups.SetText("...")
			pingU.elexLaber.logged.SetText("...")
			pingU.elexLaber.paused.SetText("...")
		}
	})

	return topVbox
}
