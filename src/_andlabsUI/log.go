package _andlabsUI

import (
	"log"
	"os"
)

type boardLogger struct {
	log *log.Logger
}

var boardLog = boardLogger{}

func initLogger() (*os.File) {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panic(err)
	}
	boardLog.log = log.New(logFile, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
	return logFile
}

func (dLog boardLogger) write(sLog string) {
	dLog.log.Println(sLog)
	log.Println(sLog)
}