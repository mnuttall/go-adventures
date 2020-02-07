package main

import (
	"testing"
)

func TestGenTrickyPayload(t *testing.T) {
	result := genTrickyPayload(generateRegularName)
	t.Logf("Test got %s", result)
}

func generateRegularName() string {
	return "alwaysTheSame"
}
