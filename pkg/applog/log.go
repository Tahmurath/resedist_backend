package applog

import (
	"fmt"
	"log"
	"resedist/pkg/config"
)

const colorRed = "\033[0;31m"

const colorNone = "\033[0m"

func Info(msg string) {
	cfg := config.Get()

	switch cfg.Log.LogLevel {
	case 0:
		return
	case 1:
		//fmt.Fprintf(os.Stdout, "Red: \033[0;31m %s None: \033[0m %s", msg, msg)
		//fmt.Fprintf(os.Stdout, "Red: %s %s None: %s %s", colorRed, msg, colorNone, msg)
		fmt.Println(colorRed, msg, colorNone)
	case 2:
		log.Printf(msg)
		//log to file
	default:
		return
	}
}
