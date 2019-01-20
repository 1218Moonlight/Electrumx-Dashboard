package _andlabsUI

import "github.com/andlabs/ui"

type window struct {
	main       *ui.Window
	title      string
	width      int
	height     int
	hasMenubar bool
}

func newWindow() window{
	return window{
		main:       nil,
		title:      "Electrumx-Dashboard",
		width:      600,
		height:     600,
		hasMenubar: false,
	}
}