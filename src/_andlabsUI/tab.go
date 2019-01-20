package _andlabsUI

import (
	"github.com/andlabs/ui"
	"_Log"
)

func logTab() ui.Control {
	_Log.Write("Setting logTab")
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	status := ui.NewMultilineEntry()
	status.SetReadOnly(true)
	vbox.Append(status, true)


	go _Log.Handler(status)

	return vbox
}

func serverTab() ui.Control {
	_Log.Write("Setting serverTab")
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)

	ipText := ui.NewEntry()
	ipBtn := ui.NewButton("Connet")
	hbox.Append(ipText, true)
	hbox.Append(ipBtn, false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	pingStatus := ui.NewEntry()
	pingStatus.SetReadOnly(true)
	vbox.Append(pingStatus, true)

	ipBtn.OnClicked(func(button *ui.Button) {
		if !pingBool {
			go serverPing(ipText.Text(), pingStatus)
			ipBtn.SetText("Close")
			pingBool = true
		} else {
			ipBtn.SetText("Connet")
			pingBool = false
		}
	})

	return vbox
}

func electrumxTab() ui.Control {
	_Log.Write("Setting electrumxTab")
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	status := ui.NewButton("getElectrumxInfo")
	vbox.Append(status, false)

	go electrumxInfo(status)

	return vbox
}