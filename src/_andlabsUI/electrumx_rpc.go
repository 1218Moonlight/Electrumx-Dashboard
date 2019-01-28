package _andlabsUI

import (
	"encoding/json"
	"log"
)

type elexGetinfo struct {
	Closing      int    `json:"closing"`
	Daemon       string `json:"daemon"`
	Daemonheight int    `json:"daemonheight"`
	Dbheight     int    `json:"dbheight"`
	Errors       int    `json:"errors"`
	Groups       int    `json:"groups"`
	Logged       int    `json:"logged"`
	Paused       int    `json:"paused"`
}

func Getinfo() []byte {
	jsonGetinfo, err := json.Marshal(elexGetinfo{
		Closing:       1,
		Daemon:        "192.168.0.2:8332/",
		Daemonheight: 520527,
		Dbheight:     520527,
		Errors:        0,
		Groups:        7,
		Logged:        0,
		Paused:        0,
	})
	if err != nil {
		log.Println(err)
	}
	return jsonGetinfo
}
