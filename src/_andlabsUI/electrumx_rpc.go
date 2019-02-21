package _andlabsUI

import (
	"math/big"
	"github.com/andlabs/ui"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const (
	urlHttp  = "http://"
	endPoint = "/getinfo"
)

type getinfo struct {
	JsonRpc string  `json:"jsonrpc"`
	Id      big.Int `json:"id"`
	Result struct {
		Daemon        string  `json:"daemon"`
		Daemon_height big.Int `json:"daemon_height"`
		Db_height     big.Int `json:"db_height"`
		Closing       big.Int `json:"closing"`
		Errors        big.Int `json:"errors"`
		Groups        big.Int `json:"groups"`
		Logged        big.Int `json:"logged"`
		Paused        big.Int `json:"paused"`
		Pid           big.Int `json:"pid"`
		Peers struct {
			Bad   big.Int `json:"bad"`
			Good  big.Int `json:"good"`
			Never big.Int `json:"never"`
			Stale big.Int `json:"stale"`
			Total big.Int `json:"total"`
		} `json:"peers"`
		Requests big.Int `json:"requests"`
		Sessions big.Int `json:"sessions"`
		Subs     big.Int `json:"subs"`
		Txs_sent big.Int `json:"txs_sent"`
		Uptime   string  `json:"uptime"`
	} `json:"result"`
}

type electrumxLaber struct {
	closing      *ui.Label
	daemon       *ui.Label
	daemonHeight *ui.Label
	dbHeight     *ui.Label
	errors       *ui.Label
	groups       *ui.Label
	logged       *ui.Label
	paused       *ui.Label
}

func electrumxGetinfo(url string, elexLaber electrumxLaber) {
	resp, err := http.Get(urlHttp + url + endPoint)
	if checkError(err,false) {return}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if checkError(err, false) {return}
	var info getinfo
	err = json.Unmarshal(data, &info)
	if checkError(err, false) {return}

	elexLaber.closing.SetText("Closing : " + info.Result.Closing.String())
	elexLaber.daemon.SetText("Daemon : " + info.Result.Daemon)
	elexLaber.daemonHeight.SetText("DaemonHeight : " + info.Result.Daemon_height.String())
	elexLaber.dbHeight.SetText("DB Height : " + info.Result.Db_height.String())
	elexLaber.errors.SetText("Errors : " + info.Result.Errors.String())
	elexLaber.groups.SetText("Groups : " + info.Result.Groups.String())
	elexLaber.logged.SetText("Logged : " + info.Result.Logged.String())
	elexLaber.paused.SetText("Paused : " + info.Result.Paused.String())
}
