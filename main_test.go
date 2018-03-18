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

func TestGetGrid(t *testing.T) {
	l := PrepareLevel(0)
	debug("%v\n", l)
        got:=grid.GetGrid(1,1)
	expect := DOTBLUE
	if got != expect {
		t.Error(fmt.Sprintf("getgrid fail. Expected %v got %v", expect, got))
	}
}


func TestSetGrid(t *testing.T) {
	l := PrepareLevel(0)
	debug("%v\n", l)
        grid.SetGrid(1,1,DOTRED)
        got:=grid.GetGrid(1,1)
	expect := DOTRED
	if got != expect {
		t.Error(fmt.Sprintf("getgrid fail. Expected %v got %v", expect, got))
	}
}
func TestPrepareLevel(t *testing.T) {
	l := PrepareLevel(0)
	debug("%v\n", l)
}

func TestMakeGridScanner(t *testing.T) {
	l := PrepareLevel(0)
	debug("%v\n", l)
	//fmt.Println(l.Render())

	// check for no dots
	expectd := true
	scan := makeGridScanner(DOTGREEN)
	_, _, gotd := scan()
	if gotd != expectd {
		t.Error(fmt.Sprintf("makeGridScanner fail. Expected done=true but got  %v",  gotd))
	}

	// check for proper counting
	// 	fmt.Println("----looking for Dotred")
	got := 0
	expect := 4
	scan = makeGridScanner(DOTRED)
	for x, y, done := scan(); done == false; x, y, done = scan() {

		if grid.GetGrid(x, y) == DOTRED {
			got++
		}
	}
	if got != expect {
		t.Error(fmt.Sprintf("makeGridScanner fail. Expected DOTRED %v got %v", expect, got))
	}

	// 	fmt.Println("----looking for Dotyellow")
	got = 0
	expect = 1
	scan = makeGridScanner(DOTYELLOW)
	for x, y, done := scan(); done == false; x, y, done = scan() {
		if grid.GetGrid(x, y) == DOTYELLOW {
			got++
		}
	}
	if got != expect {
		t.Error(fmt.Sprintf("makeGridScanner fail. Expected DOTYELLOW %v got %v", expect, got))
	}

	// check for no dots
	// 	fmt.Println("----looking for Dotgreen (expecting none)")
	got = 0
	expect = 0
	scan = makeGridScanner(DOTGREEN)
	for x, y, done := scan(); done == false; x, y, done = scan() {
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
	// input to short
	input = "a"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
	// input too short
	input = "a1"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
	// input out of bounds
	input = "h5l"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
	// select out of  bounds
	input = "a1l"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
	// select out of  bounds
	input = "a1u"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
	// select out of  bounds
	input = "a4d"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
	// select out of  bounds
	input = "d4r"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}

	// correct move made
	input = "a1r"
	move, err = PlayerInputOk(input)
	if len(move) != 2 {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
	// incorrect selection
	input = "a1rrrd"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
	// good move
	input = "b4r"
	move, err = PlayerInputOk(input)
	if err == true {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
	// bad selection - two colours
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
	
	//cant select null
	input = "b3l"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
	// cant select null
	input = "b3rrr"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}

	// wildcard good
	input = "a2r"
	move, err = PlayerInputOk(input)
	if err == false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
	// wildcard good
	input = "b2l"
	move, err = PlayerInputOk(input)
	if err == false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
	// wildcard good
	input = "a1d"
	move, err = PlayerInputOk(input)
	if err == false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
	// wildcard bad - two colours
	input = "a1dr"
	move, err = PlayerInputOk(input)
	if err == true {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
	// wildcard bad - two colours
	input = "a2ru"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
	// good wildcard
	input = "b3d"
	move, err = PlayerInputOk(input)
	if err != false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}
	
        // good wildcard
        input = "c3l"
	move, err = PlayerInputOk(input)
	if err == false {
		t.Error("invalid input recognised as false failed with input of:", input)
	}


	
}
