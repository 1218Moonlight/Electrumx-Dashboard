package _andlabsUI

import (
	"github.com/andlabs/ui"
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

	ipBtn.OnClicked(func(button *ui.Button) {
		if !pingBool {
			//go serverPing(ipText.Text(), pingStatus)
			ipBtn.SetText("Close")
			pingBool = true
		} else {
			ipBtn.SetText("Connet")
			pingBool = false
			//pingExit <- 0
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