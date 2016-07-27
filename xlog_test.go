package xlog

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestErrorOutput(t *testing.T) {
	errbuf := new(bytes.Buffer)
	errLog := New(1, errbuf)
	if errLog != nil {
		t.Error(errLog)
		t.Fail()
	}
	Error.Println("A message that contains an error")
	result := errbuf.String()
	p := "ERROR:"
	s := "A message that contains an error\n"
	if !strings.HasPrefix(result, p) || !strings.HasSuffix(result, s) {
		t.Errorf("\nExpected prefix %s & suffix %s\nReceived %s\n", p, s, result)
	}
}
func TestNoWriters(t *testing.T) {
	errLog := New(0)
	if errLog != nil {
		t.Error(errLog)
		t.Fail()
	}
	//Maybe capture stdout?
}

func TestAllOutput(t *testing.T) {
	var eb, wb, db, ib = new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	errLog := New(4, eb, wb, db, ib)
	logmap := map[int]*log.Logger{0: Error, 1: Warning, 2: Debug, 3: Info}
	bufmap := map[int]*bytes.Buffer{0: eb, 1: wb, 2: db, 3: ib}
	if errLog != nil {
		t.Error(errLog)
		t.Fail()
	}
	var tests = []struct {
		pre  string
		body string
	}{
		{"ERROR:", "A message that contains an error"},
		{"WARNING:", "A message that contains a warning"},
		{"DEBUG:", "A message that contains a debug message"},
		{"INFO:", "A message that contains info"},
	}
	for k, v := range tests {
		logitem := logmap[k]
		logitem.Println(v.body)
		result := strings.TrimSpace(bufmap[k].String())
		if !strings.HasPrefix(result, v.pre) || !strings.HasSuffix(result, v.body) {
			t.Errorf("\nExpected Prefix %s | Suffix %s\nReceived messge %s\n", v.pre, v.body, result)
		}
	}
}
