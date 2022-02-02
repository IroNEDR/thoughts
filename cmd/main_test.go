package main

import "testing"

func TestAppSetup(t *testing.T) {
	err := AppSetup()
	if err != nil {
		t.Errorf("failed AppSetup with error: %v", err)
	}
}
