package main

import "testing"
import "fmt"

func TestGrid(t *testing.T) {
	g := Grid{Width: 2, Height: 2, cells: []Piece{DOTRED, DOTRED, DOTRED, DOTRED}}
	expect := "|rr|rr|"
	got := fmt.Sprintf("%v", g)
	if got != expect {
		t.Error("Grid stringy fail. Expected "+expect+" got ", got)
	}

	g = Grid{Width: 3, Height: 2, cells: []Piece{DOTRED, DOTANCHOR, DOTRED, DOTRED, DOTANCHOR, DOTRED}}
	expect = "|rar|rar|"
	got = fmt.Sprintf("%v", g)
	if got != expect {
		t.Error("Grid stringy fail. Expected "+expect+" got ", got)
	}
}

func TestA1splitter(t *testing.T) {
	got1, got2 := a1splitter("a")
	expect1 := "a"
	expect2 := 1
	if got1 != expect1 || got2 != expect2 {
		t.Error(fmt.Sprintf("a1splitter fail. Expected %v%v got %v%v\n", expect1, expect2, got1, got2))
	}
	got1, got2 = a1splitter("b400")
	expect1 = "b"
	expect2 = 400
	if got1 != expect1 || got2 != expect2 {
		t.Error(fmt.Sprintf("a1splitter fail. Expected %v%v got %v%v\n", expect1, expect2, got1, got2))
	}
}
func TestSymbol2piece(t *testing.T) {
	gots2p := symbol2piece("r")
	expects2p := DOTRED
	if gots2p != expects2p {
		t.Error(fmt.Sprintf("symbol2piece fail. Expected %v got %v", expects2p, gots2p))
	}
}
func TestPrepareLevel(t *testing.T) {
	l := PrepareLevel(0)
	fmt.Println(l)
}
