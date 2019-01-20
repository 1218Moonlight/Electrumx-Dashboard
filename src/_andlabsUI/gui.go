package _andlabsUI

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"_Log"
)

func gui() {
	_Log.Write("Setting window GUI")
	win := newWindow()

	win.main = ui.NewWindow(win.title, win.width, win.height, win.hasMenubar)

	win.main.OnClosing(func(window *ui.Window) bool {
		ui.Quit()
		_Log.Write("OnClosing")
		return true
	})
	ui.OnShouldQuit(func() bool {
		win.main.Destroy()
		_Log.Write("OnShouldQuit")
		return true
	})

	tab := ui.NewTab()
	win.main.SetChild(tab)
	win.main.SetMargined(true)

	tab.Append("Server", serverTab())
	tab.SetMargined(0, true)

	tab.Append("Electrumx", electrumxTab())
	tab.SetMargined(1, true)

	tab.Append("Log", logTab())
	tab.SetMargined(2, true)

	_Log.Write("Start window show")
	win.main.Show()
}
