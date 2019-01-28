package _andlabsUI

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func gui() {
	win := newWindow()

	win.main = ui.NewWindow(win.title, win.width, win.height, win.hasMenubar)

	win.main.OnClosing(func(window *ui.Window) bool {
		ui.Quit()
		boardLog.write("OnClosing")
		return true
	})
	ui.OnShouldQuit(func() bool {
		win.main.Destroy()
		boardLog.write("OnShouldQuit")
		return true
	})

	tab := ui.NewTab()
	win.main.SetChild(tab)
	win.main.SetMargined(true)

	tab.Append("Server", serverTab())
	tab.SetMargined(0, true)

	//tab.Append("Log", logTab())
	//tab.SetMargined(2, true)

	win.main.Show()
}