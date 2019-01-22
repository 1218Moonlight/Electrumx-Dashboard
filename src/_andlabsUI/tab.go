package _andlabsUI

import (
	"github.com/andlabs/ui"
	"time"
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

	elexGroup := ui.NewGroup("getInfo")
	elexGroup.SetMargined(true)

	groupVbox := ui.NewVerticalBox()
	groupVbox.SetPadded(true)
	elexGroup.SetChild(groupVbox)
	vbox.Append(elexGroup, true)

	elexStatus := ui.NewLabel("")
	groupVbox.Append(elexStatus, true)

	//go electrumxInfo(status, elexStatus)

	return vbox
}
