package main

import "testing"

func TestEnvVarNotDefined(t *testing.T) {
	_, err := GlobalEnv.Find("a")

	if err == nil {
		t.Errorf("got %v want nil", err)
	}
}
