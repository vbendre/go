package main

import (
	"strings"
	"testing"
)

func TestToString(t *testing.T) {
	sl := sliceString{"a", "b"}

	if !strings.Contains(sl.toString(), ",") {
		t.Errorf("expected , in toString %v", sl.toString())
	}

}
