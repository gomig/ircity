package ircity_test

import (
	"testing"

	"github.com/gomig/ircity"
)

func TestQuery(t *testing.T) {
	t.Log(ircity.States())
	t.Log(ircity.State(3))
	t.Log(ircity.Cities(20))
	if ircity.City(3222) != nil {
		t.Fatalf("Failed!")
	}
}
