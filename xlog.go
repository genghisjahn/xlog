package xlog

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	//Info logger instance for things like starts steps, scheduled tasks
	Info *log.Logger
	//Debug logger instance for debug statements to troubleshoot
	Debug *log.Logger
	//Warning logger instance for (no one reads warnings)
	Warning *log.Logger
	//Error Logger instance for error events, stuff that shouldn't be happening or shouldn't be happening a lot
	Error *log.Logger
)

const (
	//Errorlvl sets logging level to Error
	Errorlvl = 1
	//Warninglvl sets logging level to Warning and above
	Warninglvl = 2
	//Debuglvl sets logging level to Debug and above
	Debuglvl = 3
	//Infolvl sets logging to Info
	Infolvl = 4
)
const (
	errorname   = "ERROR: "
	warningname = "WARNING: "
	debugname   = "DEBUG: "
	infoname    = "INFO: "
)

//New creates a new set of loggers for various logging levels
func New(lvl int, logw ...io.Writer) error {
	if lvl < Errorlvl || lvl > Infolvl {
		return fmt.Errorf("lvl must be 1, 2, 3, or 4")
	}

	Error = log.New(os.Stderr, errorname, log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, warningname, log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(os.Stdout, debugname, log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, infoname, log.Ldate|log.Ltime|log.Lshortfile)

	if lvl >= Errorlvl {
		Error = getLogger(Errorlvl, errorname, logw)
		Warning = log.New(ioutil.Discard, warningname, 0)
		Debug = log.New(ioutil.Discard, debugname, 0)
		Info = log.New(ioutil.Discard, infoname, 0)
	}
	if lvl >= Warninglvl {
		Warning = getLogger(Warninglvl, warningname, logw)
		Debug = log.New(ioutil.Discard, debugname, 0)
		Info = log.New(ioutil.Discard, infoname, 0)
	}
	if lvl >= Debuglvl {
		Debug = getLogger(Debuglvl, debugname, logw)
		Info = log.New(ioutil.Discard, infoname, 0)
	}
	if lvl == Infolvl {
		Info = getLogger(Infolvl, infoname, logw)
	}
	return nil
}

func getLogger(lvl int, name string, logw []io.Writer) *log.Logger {
	if len(logw) >= lvl {
		return log.New(logw[lvl-1],
			name,
			log.Ldate|log.Ltime|log.Lshortfile)
	}
	return log.New(os.Stdout, name, log.Ldate|log.Ltime|log.Lshortfile)
}
