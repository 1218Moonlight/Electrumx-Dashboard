package _andlabsUI

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func Start() {
	logFile := initLogger()
	defer logFile.Close()

	boardLog.writeInfo("Starting to chan")

	go chanHandler()

	ui.Main(gui)
}
