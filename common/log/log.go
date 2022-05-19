package log

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

//TODO: Logs in file

func Error(err interface{}, format ...interface{}) {
	dt := time.Now()
	d := dt.Format("02.01.2006 15:04:05")
	var log string
	switch v := err.(type) {
	case error:
		log = v.(error).Error()
	default:
		log = v.(string)
	}
	color.Red(fmt.Sprintf("[%s] %s", d, log), format...)
}

func Info(log string, format ...interface{}) {
	dt := time.Now()
	d := dt.Format("02.01.2006 15:04:05")
	color.Green(fmt.Sprintf("[%s] %s", d, log), format...)
}

func Warning(log string, format ...interface{}) {
	dt := time.Now()
	d := dt.Format("02.01.2006 15:04:05")
	color.Yellow(fmt.Sprintf("[%s] %s", d, log), format...)
}

// TODO: Print debug only if Debug in config

func Debug(log string, format ...interface{}) {
	dt := time.Now()
	d := dt.Format("02.01.2006 15:04:05")
	if true {
		color.Blue(fmt.Sprintf("[%s] %s", d, log), format...)
	}
}
