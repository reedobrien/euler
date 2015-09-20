package main

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestSqos(t *testing.T) {
	var want, got int64
	want = 3025
	got = sqos(10)
	equals(t, got, want)
}

func TestSosq(t *testing.T) {
	var want, got int64
	want = 385
	got = sosq(10)
	equals(t, got, want)
}

// equals fails the test if got is not equal to want.
func equals(tb testing.TB, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\tgot: %#v\n\n\twant: %#v\033[39m\n\n", filepath.Base(file), line, got, want)
		tb.FailNow()
	}
}
