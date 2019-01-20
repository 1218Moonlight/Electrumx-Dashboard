package _andlabsUI

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"_Log"
)

func Start(){
	_Log.Write("Starting to set andlabsUI")
	ui.Main(gui)
}

func Exit(){
	_Log.Write("Exit andlabsUi")
	_Log.Exit()
}