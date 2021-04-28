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

	reqidname string
)

const (
	//Silencelvl turns off all logging, primarily used during tests
	Silencelvl = 0
	//Errorlvl sets logging level to Error
	Errorlvl = 1
	//Warninglvl sets logging level to Warning and above
	Warninglvl = 2
	//Debuglvl sets logging level to Debug and above
	Debuglvl = 3
	//Infolvl sets logging to Info
	Infolvl = 4

	errorname   = "ERROR: "
	warningname = "WARNING: "
	debugname   = "DEBUG: "
	infoname    = "INFO: "
	xborbits    = log.Ldate | log.Ltime | log.Lshortfile
)

func NewWithReqID(lvl int, reqid string) error {
	reqidname = "ReqID: " + reqid
	return New(lvl)
}

//New creates a new set of loggers for various logging levels
func New(lvl int, logw ...io.Writer) error {
	if lvl < Silencelvl || lvl > Infolvl {
		return fmt.Errorf("lvl must be 0,1, 2, 3, or 4")
	}

	Error = log.New(os.Stderr, reqidname+errorname, xborbits)
	Warning = log.New(os.Stdout, reqidname+warningname, xborbits)
	Debug = log.New(os.Stdout, reqidname+debugname, xborbits)
	Info = log.New(os.Stdout, reqidname+infoname, xborbits)

	if lvl >= Silencelvl {
		Error = log.New(ioutil.Discard, reqidname+errorname, 0)
		Warning = log.New(ioutil.Discard, reqidname+warningname, 0)
		Debug = log.New(ioutil.Discard, reqidname+debugname, 0)
		Info = log.New(ioutil.Discard, reqidname+infoname, 0)
	}

	if lvl >= Errorlvl {
		Error = getLogger(Errorlvl, reqidname+errorname, logw)
		Warning = log.New(ioutil.Discard, reqidname+warningname, 0)
		Debug = log.New(ioutil.Discard, reqidname+debugname, 0)
		Info = log.New(ioutil.Discard, reqidname+infoname, 0)
	}
	if lvl >= Warninglvl {
		Warning = getLogger(Warninglvl, reqidname+warningname, logw)
		Debug = log.New(ioutil.Discard, reqidname+debugname, 0)
		Info = log.New(ioutil.Discard, reqidname+infoname, 0)
	}
	if lvl >= Debuglvl {
		Debug = getLogger(Debuglvl, reqidname+debugname, logw)
		Info = log.New(ioutil.Discard, reqidname+infoname, 0)
	}
	if lvl == Infolvl {
		Info = getLogger(Infolvl, reqidname+infoname, logw)
	}
	return nil
}

func getLogger(lvl int, name string, logw []io.Writer) *log.Logger {
	if len(logw) >= lvl {
		return log.New(logw[lvl-1],
			name,
			xborbits)
	}
	if lvl == 1 {
		return log.New(os.Stderr, name, xborbits)
	}
	return log.New(os.Stdout, name, xborbits)
}
