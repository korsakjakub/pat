package tests // import github.com/korsakjakub/pat/tests

import (
	"github.com/korsakjakub/pat/pkg"
	"testing"
)

func TestNewTable(t *testing.T) {
	p1, _ := NewPlayer("jakub", 1000.0, Settings{})
	p2, _ := NewPlayer("ola", 1001.20, Settings{})
	p := []Player{*p1, *p2}

	if table, err := NewTable(p, Settings{}); err != nil {
		t.Error(err)
	} else {
		t.Log(table)
	}
}
