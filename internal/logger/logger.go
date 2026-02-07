package logger

import (
	"log"
	"os"
)

var (
	info  = log.New(os.Stdout, "[info] ", log.LstdFlags)
	warn  = log.New(os.Stdout, "[warn] ", log.LstdFlags)
	error = log.New(os.Stderr, "[error] ", log.LstdFlags)
)

func Info(msg string)  { info.Println(msg) }
func Warn(msg string)  { warn.Println(msg) }
func Error(msg string) { error.Println(msg) }
