package util

/*
   Helper for logging
*/

import (
	"log"
	"os"
)

// loggin
var Log *log.Logger

func NewLog(logfile string) {
	CleanUp(logfile)
	file, err := os.Create(logfile)
	if err != nil {
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
