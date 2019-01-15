package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"log"
)

var mainWin *ui.Window

func main() {
	ui.Main(GUI)
}

func GUI() {
	mainWin = ui.NewWindow("Electrumx-Dashboard", 600, 600, false)

	mainWin.OnClosing(func(window *ui.Window) bool {
		ui.Quit()
		log.Println("OnClosing")
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainWin.Destroy()
		log.Println("OnShouldQuit")
		return true
	})

	tab := ui.NewTab()
	mainWin.SetChild(tab)
	mainWin.SetMargined(true)

	tab.Append("Server", mainTab())

	mainWin.Show()
}

func mainTab() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(ui.NewLabel("Hello, World!"), false)
	return vbox
}
