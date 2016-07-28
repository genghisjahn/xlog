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
	errorName   = "ERROR: "
	warningName = "WARNING: "
	debugname   = "DEBUG: "
	infoname    = "INFO: "
)

//New creates a new set of loggers for various logging levels
func New(lvl int, logw ...io.Writer) error {
	lg := len(logw)
	if lg > 4 {
		return fmt.Errorf("You supplied %d flag writers.  4 is the max", lg)
	}

	Error = log.New(os.Stderr, errorName, log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, warningName, log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(os.Stdout, debugname, log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, infoname, log.Ldate|log.Ltime|log.Lshortfile)

	if len(logw) > 0 {
		Error = log.New(ioutil.Discard, errorName, 0)
		Warning = log.New(ioutil.Discard, warningName, 0)
		Debug = log.New(ioutil.Discard, debugname, 0)
		Info = log.New(ioutil.Discard, infoname, 0)
	}

	switch lvl {
	case 1:
		Error = log.New(logw[0],
			errorName,
			log.Ldate|log.Ltime|log.Lshortfile)
	case 2: //Warning & Error
		Error = log.New(logw[0],
			errorName,
			log.Ldate|log.Ltime|log.Lshortfile)
		Warning = log.New(logw[1],
			warningName,
			log.Ldate|log.Ltime|log.Lshortfile)
	case 3: //Debug, Warning & Error
		Error = log.New(logw[0],
			errorName,
			log.Ldate|log.Ltime|log.Lshortfile)
		Warning = log.New(logw[1],
			warningName,
			log.Ldate|log.Ltime|log.Lshortfile)
		Debug = log.New(logw[2],
			debugname,
			log.Ldate|log.Ltime|log.Lshortfile)
	case 4: //All
		Error = log.New(logw[0],
			errorName,
			log.Ldate|log.Ltime|log.Lshortfile)
		Warning = log.New(logw[1],
			warningName,
			log.Ldate|log.Ltime|log.Lshortfile)
		Debug = log.New(logw[2],
			debugname,
			log.Ldate|log.Ltime|log.Lshortfile)
		Info = log.New(logw[3],
			infoname,
			log.Ldate|log.Ltime|log.Lshortfile)
	}
	return nil
}
