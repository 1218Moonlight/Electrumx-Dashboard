package _andlabsUI

import (
	"log"
	"os"
)

type boardLogger struct {
	log *log.Logger
}

var boardLog = boardLogger{}

func initLogger() (*os.File, error) {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if checkError(err, true){return nil, err}
	boardLog.log = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	return logFile, nil
}

func (dLog boardLogger) writeInfo(sLog interface{}) {
	dLog.log.SetPrefix("[ INFO ] ")
	dLog.log.Println(sLog)
	log.Println(sLog)
}

func (dLog boardLogger) writeError(sLog interface{}) {
	dLog.log.SetPrefix("[ ERROR ] ")
	dLog.log.Println(sLog)
	log.Println(sLog)
}

func checkError(err error, notLog bool) bool {
	if err != nil {
		if !notLog {boardLog.writeError(err.Error())}
		return true
	}
	return false
}