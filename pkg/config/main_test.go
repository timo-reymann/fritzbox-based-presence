package config

import "testing"

func TestRead(t *testing.T) {
	err := Read()
	if err == nil {
		t.Fatal("Expected reading to happen with an error as no env var is specified")
	}
}

func TestPrintUsage(t *testing.T) {
	PrintUsage()
}

func TestGet(t *testing.T) {
	Get()
}
