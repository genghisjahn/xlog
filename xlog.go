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
	//Debug logger intsance for debug statements to troubleshoot
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
	//Infolvl sets logging to Info and above
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
	lg := len(logw)
	if lvl < Errorlvl || lvl > Infolvl {
		return fmt.Errorf("lvl must be 1, 2, 3, or 4")
	}

	Error = log.New(os.Stderr, errorname, log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, warningname, log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(os.Stdout, debugname, log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, infoname, log.Ldate|log.Ltime|log.Lshortfile)

	if lvl >= Errorlvl {
		if lg >= lvl {
			Error = log.New(logw[0],
				errorname,
				log.Ldate|log.Ltime|log.Lshortfile)
		} else {
			Error = log.New(os.Stderr, errorname, log.Ldate|log.Ltime|log.Lshortfile)
		}
		Warning = log.New(ioutil.Discard, warningname, 0)
		Debug = log.New(ioutil.Discard, debugname, 0)
		Info = log.New(ioutil.Discard, infoname, 0)
	}
	if lvl >= Warninglvl {
		if lg >= lvl {
			Warning = log.New(logw[1],
				warningname,
				log.Ldate|log.Ltime|log.Lshortfile)
		} else {
			Warning = log.New(os.Stdout, warningname, log.Ldate|log.Ltime|log.Lshortfile)
		}
		Debug = log.New(ioutil.Discard, debugname, 0)
		Info = log.New(ioutil.Discard, infoname, 0)
	}
	if lvl >= Debuglvl {
		if lg >= lvl {
			Debug = log.New(logw[2],
				debugname,
				log.Ldate|log.Ltime|log.Lshortfile)
		} else {
			Debug = log.New(os.Stdout, debugname, log.Ldate|log.Ltime|log.Lshortfile)
		}
		Info = log.New(ioutil.Discard, infoname, 0)
	}
	if lvl == Infolvl {
		if lg >= lvl {
			Info = log.New(logw[3],
				infoname,
				log.Ldate|log.Ltime|log.Lshortfile)
		} else {
			Info = log.New(os.Stdout, infoname, log.Ldate|log.Ltime|log.Lshortfile)

		}
	}
	return nil
}
