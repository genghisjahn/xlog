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

func TestInvalidWritersGivenLevel(t *testing.T) {
	errbuf := new(bytes.Buffer)
	errLog := New(2, errbuf)
	if errLog == nil {
		t.Error("This should have returned an error.  Level 2 was specified, but only 1 logger was given.")
		t.Fail()
	}
}

func TestErrorNoWarningOutput(t *testing.T) {
	errbuf := new(bytes.Buffer)
	wbuf := new(bytes.Buffer)
	errLog := New(1, errbuf, wbuf)
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
	Warning.Println("This should not print.")
	wo := wbuf.String()
	if wo != "" {
		t.Errorf("Warning.Println shouldn't print anything, but it printed %s\n", wo)
	}
}

func TestNoWriters(t *testing.T) {
	errLog := New(0)
	if errLog != nil {
		t.Error(errLog)
		t.Fail()
	}
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
