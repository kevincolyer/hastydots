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

func TestmakeGridScanner(t *testing.T) {
	l := PrepareLevel(0)
	l.Render()
	got := 0
	expect := 4
	scan := makeGridScanner(DOTRED)
	for x, y := scan(); x >= 0; {
		if grid.GetGrid(x, y) == DOTRED {
			got++
		}
	}
	if got != expect {
		t.Error(fmt.Sprintf("makeGridScanner fail. Expected DOTRED %v got %v", expect, got))
	}

	got = 0
	expect = 2
	scan = makeGridScanner(DOTYELLOW)
	for x, y := scan(); x >= 0; {
		if grid.GetGrid(x, y) == DOTYELLOW {
			got++
		}
	}
	if got != expect {
		t.Error(fmt.Sprintf("makeGridScanner fail. Expected DOTYELLOW %v got %v", expect, got))
	}

	got = 0
	expect = 0
	scan = makeGridScanner(DOTGREEN)
	for x, y := scan(); x >= 0; {
		if grid.GetGrid(x, y) == DOTGREEN {
			got++
		}
	}
	if got != expect {
		t.Error(fmt.Sprintf("makeGridScanner fail. Expected DOTGREEN %v got %v", expect, got))
	}

}

func TestPlayerInputOk(t *testing.T) {
	l := PrepareLevel(0)
	l.Render()
	var move []Move
	var err bool
	var input string
	input = "a"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	input = "a1"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	input = "h5l"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	input = "a1l"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	input = "a1u"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	input = "a4d"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	input = "d4r"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}

	input = "a1r"
	move, err = PlayerInputOk(input)
	if len(move) != 2 {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	input = "a1rrrd"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	//cant select anchor
	input = "b4r"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	input = "a4rr"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	//cant select null
	input = "a3r"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	input = "b3rrr"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}

	//wildcard
	input = "a2r"
	move, err = PlayerInputOk(input)
	if err == false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	input = "a2u"
	move, err = PlayerInputOk(input)
	if err == false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	input = "a2ru"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	input = "b3d"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}

}
