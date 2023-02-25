package xlog

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Key string

var (
	//Info logger instance for things like starts steps, scheduled tasks
	info *log.Logger
	//Debug logger instance for debug statements to troubleshoot
	debug *log.Logger
	//Warning logger instance for (no one reads warnings)
	warning *log.Logger
	//Error Logger instance for error events, stuff that shouldn't be happening or shouldn't be happening a lot
	errorl *log.Logger

	//audit is for specic lines we want to audit for cloud watch/stats/legder
	audit *log.Logger

	ctxKeys = []Key{}
)

func getvalsfromctx(ctx ...context.Context) string {
	result := ""
	if len(ctx) == 1 {
		for _, v := range ctxKeys {
			val := ctx[0].Value(v)
			if val != nil {
				result += "[" + string(v) + ":" + val.(string) + "] "
			}
		}
	}
	return result
}

func Info(ctx ...context.Context) *log.Logger {
	info.SetPrefix(infoname + getvalsfromctx(ctx...))
	return info
}
func Debug(ctx ...context.Context) *log.Logger {
	debug.SetPrefix(debugname + getvalsfromctx(ctx...))
	return debug
}
func Warning(ctx ...context.Context) *log.Logger {
	warning.SetPrefix(warningname + getvalsfromctx(ctx...))
	return warning
}
func Error(ctx ...context.Context) *log.Logger {
	errorl.SetPrefix(errorname + getvalsfromctx(ctx...))
	return errorl
}

func Audit(ctx ...context.Context) *log.Logger {
	audit.SetPrefix(auditname + getvalsfromctx(ctx...))
	return audit
}

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

	Auditlvl = 5

	errorname   = "ERROR: "
	warningname = "WARNING: "
	debugname   = "DEBUG: "
	infoname    = "INFO: "
	auditname   = "AUDIT: "

	xborbits = log.Ldate | log.Ltime | log.Lshortfile
)

// New creates a new set of loggers for various logging levels
func New(lvl int, ctxkeys []Key, logw ...io.Writer) error {
	if lvl < Silencelvl || lvl > Auditlvl {
		return fmt.Errorf("lvl must be 0,1, 2, 3, 4 or 5")
	}
	ctxKeys = ctxkeys
	errorl = log.New(os.Stderr, errorname, xborbits)
	warning = log.New(os.Stdout, warningname, xborbits)
	debug = log.New(os.Stdout, debugname, xborbits)
	info = log.New(os.Stdout, infoname, xborbits)
	audit = log.New(os.Stdout, auditname, xborbits)

	if lvl >= Silencelvl {
		errorl = log.New(ioutil.Discard, errorname, 0)
		warning = log.New(ioutil.Discard, warningname, 0)
		debug = log.New(ioutil.Discard, debugname, 0)
		info = log.New(ioutil.Discard, infoname, 0)
		audit = log.New(io.Discard, auditname, 0)
	}

	if lvl >= Errorlvl {
		errorl = getLogger(Errorlvl, errorname, logw)
		warning = log.New(ioutil.Discard, warningname, 0)
		debug = log.New(ioutil.Discard, debugname, 0)
		info = log.New(ioutil.Discard, infoname, 0)

		//This needs to be on no matter what the log level
		audit = getLogger(Auditlvl, auditname, logw)
	}
	if lvl >= Warninglvl {
		warning = getLogger(Warninglvl, warningname, logw)
		debug = log.New(ioutil.Discard, debugname, 0)
		info = log.New(ioutil.Discard, infoname, 0)

		//This needs to be on no matter what the log level
		audit = getLogger(Auditlvl, auditname, logw)
	}
	if lvl >= Debuglvl {
		debug = getLogger(Debuglvl, debugname, logw)
		info = log.New(ioutil.Discard, infoname, 0)
		//This needs to be on no matter what the log level
		audit = getLogger(Auditlvl, auditname, logw)
	}
	if lvl == Infolvl {
		info = getLogger(Infolvl, infoname, logw)
		//This needs to be on no matter what the log level
		audit = getLogger(Auditlvl, auditname, logw)
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
