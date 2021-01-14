package Cprint

import (
	"fmt"
	"github.com/gookit/color"
	"time"
)

func ToString(msg ...interface{}) string {
	return fmt.Sprint(msg...)
}

func Println(level string, msg ...interface{}) {
	msgStr := ToString(msg...)
	switch level {
	case "debug", "DEBUG", "Debug", "0":
		color.Green.Println(time.Now().Format("[2006-01-02 15:04:05]  "), msgStr)
	case "info", "INFO", "Info", "1":
		color.Cyan.Println(time.Now().Format("[2006-01-02 15:04:05]  "), msgStr)
	case "Warning", "WARNING", "warning", "warn", "Warn", "WARN", "2":
		color.Magenta.Println(time.Now().Format("[2006-01-02 15:04:05]  "), msgStr)
	case "error", "Error", "ERROR", "3":
		color.Red.Println(time.Now().Format("[2006-01-02 15:04:05]  "), msgStr)
	default:
		color.Println(time.Now().Format("[2006-01-02 15:04:05]  "), msgStr)
	}
}
