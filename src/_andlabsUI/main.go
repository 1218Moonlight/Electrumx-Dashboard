package _andlabsUI

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func Start() {
	logFile, err := initLogger()
	if !checkError(err, false){defer logFile.Close()}

	boardLog.writeInfo("Starting to chan")

	go chanHandler()

	ui.Main(gui)
}
