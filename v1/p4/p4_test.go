package main

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestStrReverse(t *testing.T) {
	values := []struct {
		in, want string
	}{
		{"12345", "54321"},
		{"23456", "65432"},
		{"11011", "11011"},
		{"998001", "100899"},
	}
	for _, tv := range values {
		got := strReverse(tv.in)
		equals(t, got, tv.want)
	}
}

func BenchmarkStrReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strReverse("1234567890")
	}
}

func BenchmarkSortReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortReverse("1234567890")
	}
}

// equals fails the test if got is not equal to want.
func equals(tb testing.TB, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\tgot: %#v\n\n\twant: %#v\033[39m\n\n", filepath.Base(file), line, got, want)
		tb.FailNow()
	}
}
