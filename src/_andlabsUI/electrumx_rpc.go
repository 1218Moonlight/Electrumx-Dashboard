package _andlabsUI

import (
	"encoding/json"
	"log"
)

type elexGetinfo struct {
	Closing       int    `json:"closing"`
	Daemon        string `json:"daemon"`
	Daemon_height int    `json:"daemon_height"`
	Db_height     int    `json:"db_height"`
	Errors        int    `json:"errors"`
	Groups        int    `json:"groups"`
	Logged        int    `json:"logged"`
	Paused        int    `json:"paused"`
}

func Getinfo() string {
	jsonGetinfo, err := json.Marshal(elexGetinfo{
		Closing:       1,
		Daemon:        "192.168.0.2:8332/",
		Daemon_height: 520527,
		Db_height:     520527,
		Errors:        0,
		Groups:        7,
		Logged:        0,
		Paused:        0,
	})
	if err != nil {
		log.Println(err)
	}
	return string(jsonGetinfo)
}
