package xlog

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestErrorWithContext(t *testing.T) {
	ctx := context.TODO()
	ctx = context.WithValue(ctx, "reqid", "abcd1234")
	errbuf := new(bytes.Buffer)
	errLog := New(1, nil, errbuf)
	if errLog != nil {
		t.Error(errLog)
	}
	Error(ctx).Println("A message that contains an error with reqid")
	result := errbuf.String()
	p := "ERROR: [reqid:abcd1234]"
	s := "A message that contains an error with reqid\n"
	if !strings.HasPrefix(result, p) || !strings.HasSuffix(result, s) {
		t.Errorf("\nExpected prefix %s & suffix %s\nReceived %s\n", p, s, result)
	}
}

func TestErrorOutput(t *testing.T) {
	errbuf := new(bytes.Buffer)
	errLog := New(1, nil, errbuf)
	if errLog != nil {
		t.Error(errLog)
	}
	Error().Println("A message that contains an error")
	result := errbuf.String()
	p := "ERROR:"
	s := "A message that contains an error\n"
	if !strings.HasPrefix(result, p) || !strings.HasSuffix(result, s) {
		t.Errorf("\nExpected prefix %s & suffix %s\nReceived %s\n", p, s, result)
	}
}

func TestErrorNoWarningOutput(t *testing.T) {
	errbuf := new(bytes.Buffer)
	wbuf := new(bytes.Buffer)
	errLog := New(1, nil, errbuf, wbuf)
	if errLog != nil {
		t.Error(errLog)
	}
	Error().Println("A message that contains an error")
	result := errbuf.String()
	p := "ERROR:"
	s := "A message that contains an error\n"
	if !strings.HasPrefix(result, p) || !strings.HasSuffix(result, s) {
		t.Errorf("\nExpected prefix %s & suffix %s\nReceived %s\n", p, s, result)
	}
	Warning().Println("This should not print.")
	wo := wbuf.String()
	if wo != "" {
		t.Errorf("Warning.Println shouldn't print anything, but it printed %s\n", wo)
	}
}

func TestNoWriters(t *testing.T) {
	errLog := New(1, nil)
	if errLog != nil {
		t.Error(errLog)
	}
}

func TestInvalidLevel(t *testing.T) {
	errLog := New(5, nil)
	if errLog == nil {
		fmt.Println("This should have returned an error, it's an invalid logging level")
	}
}

func TestDebugWithoutWriter(t *testing.T) {
	errLog := New(4, nil)
	if errLog != nil {
		t.Error(errLog)
	}
}

func TestAllOutput(t *testing.T) {
	var eb, wb, db, ib = new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	errLog := New(4, nil, eb, wb, db, ib)
	logmap := map[int]*log.Logger{0: Error(), 1: Warning(), 2: Debug(), 3: Info()}
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
