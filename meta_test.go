package ircities_test

import (
	"testing"

	"github.com/gomig/ircities"
)

func TestQuery(t *testing.T) {
	t.Log(ircities.States())
	t.Log(ircities.State(3))
	t.Log(ircities.Cities(20))
	if ircities.City(3222) != nil {
		t.Fatalf("Failed!")
	}
}
