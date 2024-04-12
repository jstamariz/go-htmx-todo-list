package logger

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	Log *log.Logger
)

func init() {
	pwd, getDwErr := os.Getwd()

	if getDwErr != nil {
		fmt.Println(getDwErr)
		os.Exit(1)
	}
	var logPath = pwd + "/info.log"

	flag.Parse()
	var file, err = os.Create(logPath)

	if err != nil {
		panic(err)
	}

	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : %s", logPath)

	file.Close()
}
