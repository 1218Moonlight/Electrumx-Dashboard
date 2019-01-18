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

	tab.Append("Server", serverTab())
	tab.SetMargined(0, true)

	tab.Append("Electrumx", electrumxTab())
	tab.SetMargined(1, true)

	win.main.Show()
}

var pingBool bool = false

func serverTab() ui.Control {
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
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	status := ui.NewButton("getElectrumxInfo")
	vbox.Append(status, false)

	go electrumxInfo(status)

	return vbox
}

func electrumxInfo(status *ui.Button) {
	status.OnClicked(func(button *ui.Button) {
		if pingBool {
			status.SetText("pingBool true")
		} else {
			status.SetText("pingBool false")
		}
	})
}

func serverPing(ip string, status *ui.Entry) {
	log.Println("Start Ping", ip)
	for {
		if !pingBool {
			status.SetText("")
			break
		}
		pinger, err := ping.NewPinger(ip)
		pinger.SetPrivileged(true)
		if err != nil {
			panic(err)
		}
		pinger.Timeout = time.Duration(time.Second * 2)
		pinger.Count = 1
		pinger.Run()
		stats := pinger.Statistics()
		if len(stats.Rtts) == 0 {
			status.SetText("Server false")
		} else {
			status.SetText("Server true")
		}
		time.Sleep(time.Second * 3)
	}
	log.Println("Stop Ping", ip)
}
