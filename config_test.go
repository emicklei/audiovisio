package main

import (
	"testing"
)

func Test_loadConfig(t *testing.T) {
	f := "test/test.yaml"
	oConfig = &f
	cfg := loadConfig()
	if got, want := cfg.Title, "Test - Episode One"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := cfg.Slides[0].ID, "1"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := cfg.PauseBeforeAudio, 500; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := cfg.Slides[0].PauseAfterAudio, 500; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := cfg.Slides[1].PauseBeforeAudio, 1000; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}
