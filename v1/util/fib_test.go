package util_test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"testing"

	. "github.com/reedobrien/euler/util"
)

func TestFibgen(t *testing.T) {
	values := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21}
	fib := Fibgen()
	for i := range values {
		s := <-fib
		got, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			t.Errorf("Fibgen failed to generate a valid number: %s", s)
		}
		equals(t, got, values[i])
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
