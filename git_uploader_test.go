package main

import "testing"

func TestRandomAlphaNumericString(t *testing.T) {
	got := randomAlphaNumericString()
	length := len(got)
	if length != 10 {
		t.Errorf("randomAlphaNumericString() = %s; want a length of 10", got)
	}
}
