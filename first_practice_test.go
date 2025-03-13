package main

import (
	"bufio"
	"bytes"
	p "golang/first_practice"
	"strings"
	"testing"
)

var testOK = `1
2
3
4
5
`

var testOkResult = `1
2
3
4
5
`

func TestOK(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOK))
	out := new(bytes.Buffer)
	err := p.Scanner(in, out)
	if err != nil {
		t.Errorf("test for ok was failed")
	}
	result := out.String()
	if result != testOkResult {
		t.Errorf("test for ok was failed, %v %v\n", result, testOkResult)
	}
	// запуск: go test -v first_practice_test.go
}

var testFail = `1
2
1
`

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := p.Scanner(in, out)
	if err == nil {
		t.Errorf("test for ok was failed, %v", err)
	}
}
