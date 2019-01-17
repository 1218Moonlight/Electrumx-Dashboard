package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"log"
	"github.com/sparrc/go-ping"
	"time"
)

type window struct {
	main       *ui.Window
	title      string
	width      int
	height     int
	hasMenubar bool
}

func main() {
	ui.Main(GUI)
}

func GUI() {
	win := window{
		main:       nil,
		title:      "Electrumx-Dashboard",
		width:      600,
		height:     600,
		hasMenubar: false,
	}

	win.main = ui.NewWindow(win.title, win.width, win.height, win.hasMenubar)

	win.main.OnClosing(func(window *ui.Window) bool {
		ui.Quit()
		log.Println("OnClosing")
		return true
	})
	ui.OnShouldQuit(func() bool {
		win.main.Destroy()
		log.Println("OnShouldQuit")
		return true
	})

	tab := ui.NewTab()
	win.main.SetChild(tab)
	win.main.SetMargined(true)

	tab.Append("Server", mainTab())

	win.main.Show()
}

func mainTab() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(ui.NewLabel("Hello, World!"), false)

	textBox := ui.NewMultilineEntry()
	textBox.SetReadOnly(true)

	go serverPing(textBox)

	vbox.Append(textBox, true)

	return vbox
}

func serverPing(textBox *ui.MultilineEntry) {
	for {
		pinger, err := ping.NewPinger("URL")
		pinger.SetPrivileged(true)
		if err != nil {
			panic(err)
		}
		pinger.Timeout = time.Duration(time.Second * 2)
		pinger.Count = 1
		pinger.Run()
		stats := pinger.Statistics()
		if len(stats.Rtts) == 0 {
			textBox.SetText("Server false")
		} else {
			textBox.SetText("Server true")
		}
		time.Sleep(time.Second * 3)
	}
}