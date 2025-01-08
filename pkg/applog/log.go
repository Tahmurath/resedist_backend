package applog

import (
	"fmt"
	"log"
	"resedist/pkg/config"
)

const Reset = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Magenta = "\033[35m"
const Cyan = "\033[36m"
const Gray = "\033[37m"
const White = "\033[97m"

func Info(msg string) {
	cfg := config.Get()

	switch cfg.Log.LogLevel {
	case 0:
		return
	case 1:
		//fmt.Fprintf(os.Stdout, "Red: \033[0;31m %s None: \033[0m %s", msg, msg)
		//fmt.Fprintf(os.Stdout, "Red: %s %s None: %s %s", colorRed, msg, colorNone, msg)
		//fmt.Println("[App-Log]", Green, msg, Reset)
		fmt.Println(Cyan + "[App-Log] " + Green + msg + Reset)
	case 2:
		log.Printf(msg)
		//log to file
	default:
		return
	}
}
