package xlog

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
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

//New creates a new set of loggers for various logging levels
func New(lvl int, logw ...io.Writer) error {
	lg := len(logw)
	if lg > 4 {
		return fmt.Errorf("You supplied %d flag writers.  4 is the max", lg)
	}
	//default should be to stdout
	Error = log.New(ioutil.Discard, "Error", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(ioutil.Discard, "Warning", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(ioutil.Discard, "Debug", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(ioutil.Discard, "Info", log.Ldate|log.Ltime|log.Lshortfile)
	//loop over from 1 to 4 instead of ranging over logw
	for k := range logw {
		switch k {
		case 0:
			Error = log.New(logw[k],
				"ERROR: ",
				log.Ldate|log.Ltime|log.Lshortfile)
		case 1:
			Warning = log.New(logw[k],
				"WARNING: ",
				log.Ldate|log.Ltime|log.Lshortfile)
		case 2:
			Debug = log.New(logw[k],
				"DEBUG: ",
				log.Ldate|log.Ltime|log.Lshortfile)
		case 3:
			Info = log.New(logw[k],
				"INFO: ",
				log.Ldate|log.Ltime|log.Lshortfile)
		}
	}
	return nil
}
