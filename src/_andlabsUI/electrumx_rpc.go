package _andlabsUI

import (
	"math/big"
	"github.com/andlabs/ui"
	"net/http"
	"io/ioutil"
)

const (
	urlHttp          = "http://"
	endPointGetinfo  = "/getinfo"
	endPointSessions = "/sessions"
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
	getinfo  *ui.Label
	sessions *ui.Label
}

func electrumxGetinfo(url string, getinfoLabel *ui.Label) {
	resp, err := http.Get(urlHttp + url + endPointGetinfo)
	if checkError(err, false) {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if checkError(err, false) {
		return
	}
	if checkError(err, false) {
		return
	}

	getinfoLabel.SetText(string(data))
}

func electrumxSessions(url string, sessionsLabel *ui.Label) {
	resp, err := http.Get(urlHttp + url + endPointSessions)
	if checkError(err, false) {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if checkError(err, false) {
		return
	}
	if checkError(err, false) {
		return
	}

	sessionsLabel.SetText(string(data))
}
