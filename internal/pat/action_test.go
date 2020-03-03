package pat

import (
	"testing"
)

func TestNewPlayer(t *testing.T) {
	sett := Settings{StartingChips: 1000}
	p, err := NewPlayer("jakub", 123.0, sett)
	if err != nil {
		t.Error(err)
	}
	if p.name != "jakub" {
		t.Errorf("The name does not match. Got %s, expected %s\n", p.name, "jakub")
	}
	if p.chips != 123.0 {
		t.Errorf("The chip amount does not match. Got %f, expected %f\n", p.chips, 123.0)
	}
	t.Log(p)

	p, err = NewPlayer("ola", 0, sett)
	if err != nil {
		t.Error(err)
	}
	if p.name != "ola" {
		t.Errorf("The name does not match. Got %s, expected %s\n", p.name, "ola")
	}
	if p.chips != 1000.0 {
		t.Errorf("The chip amount does not match. Got %f, expected %f\n", p.chips, 1000.0)
	}

	t.Log(p)
}
