package _andlabsUI

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func Start() {
	logFile := initLogger()
	defer logFile.Close()

	boardLog.write("Starting to set andlabsUI")

	ui.Main(gui)

	boardLog.write("Exit andlabsUi")
}