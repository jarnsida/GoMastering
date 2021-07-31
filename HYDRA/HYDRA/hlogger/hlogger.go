package hlogger

import (
	"log"
	"os"
	"sync"
)

type hydraLogger struct {
	*log.Logger
	filename string
}

var (
	hlogger *hydraLogger
	once    sync.Once
)

//GetInstance creates logger file once
func GetInstance() *hydraLogger {
	once.Do(func() {
		hlogger = createLogger("hydraLogger.log")
	})
	return hlogger
}

// CreateLogger creates Lger instance
func createLogger(fname string) *hydraLogger {
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal("Error creating logger file")
	}

	return &hydraLogger{
		filename: fname,
		Logger:   log.New(file, "Hydra ", log.Lshortfile),
	}
}
