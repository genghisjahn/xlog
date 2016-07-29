# xlog
Simple logging package based off the article https://www.goinggo.net/2013/11/using-log-package-in-go.html

Examples:

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/genghisjahn/xlog"
)

func main() {

	infoLevel()
	errorLevel()
	infoLevel2()
}

func infoLevel() {
	fmt.Println("Info Level")
	errLog := xlog.New(xlog.Infolvl)
	if errLog != nil {
		fmt.Println(errLog)
		return
	}
	xlog.Error.Println("This is an error!")
	xlog.Warning.Println("WARNING! Be careful!")
	xlog.Debug.Println("This is a piece of Debug info")
	xlog.Info.Println("Here's some info")
	fmt.Println("----------")
}

func errorLevel() {
	fmt.Println("Error Level")
	errLog := xlog.New(xlog.Errorlvl)
	if errLog != nil {
		fmt.Println(errLog)
		return
	}
	xlog.Error.Println("This is an error!")

	//These won't print because the level is set to Errorlvl
	xlog.Warning.Println("WARNING! Be careful!")
	xlog.Debug.Println("This is a piece of Debug info")
	xlog.Info.Println("Here's some info")
	fmt.Println("----------")
}

func infoLevel2() {
	fmt.Println("Info Level 2")
	buf := new(bytes.Buffer)
	t := &thing{}
	p := &person{}
	errLog := xlog.New(xlog.Infolvl, os.Stdout, buf, t, p)
	if errLog != nil {
		fmt.Println(errLog)
		return
	}
	xlog.Error.Println("This is an error!")
	xlog.Warning.Println("WARNING! Be careful!")
	xlog.Debug.Println("This is a piece of Debug info")
	xlog.Info.Println("Here's some info")

	fmt.Println("Output from Warning Buffer: ", buf.String())
	fmt.Println("----------")
}

type thing struct {
}

func (t *thing) Write(p []byte) (n int, err error) {
	//do something
	return
}

type person struct {
	ID   int
	Name string
	Age  int
}

func (p *person) Write(a []byte) (n int, err error) {
	b, _ := json.Marshal(*p)
	return len(b), nil
}
```
