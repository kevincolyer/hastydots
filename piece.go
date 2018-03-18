package main

import "fmt"
import "github.com/fatih/color"

//import "strings"

type Move struct {
	Colour Piece
	X      int
	Y      int
}

type Piece int

const (
	NULL  Piece = iota + 1 // #
	EMPTY                  // _ (means fill with random)

	DOTBLUE     // b
	DOTGREEN    // g
	DOTPURPLE   // p
	DOTRED      // r
	DOTWHITE    // w
	DOTYELLOW   // y
	DOTWILDCARD // *

	DOTANCHOR // a
	DOTBOMB   // o

	// we mask off 0xF here...

	ICE0 = 32
	ICE1 = 64
	ICE2 = 128
)

func piece2symbol(p Piece) (s string) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgHiYellow).SprintFunc()
	purple := color.New(color.FgMagenta).SprintFunc()
	white := color.New(color.FgHiWhite).SprintFunc()
	m := map[Piece]string{
		NULL:        "█",
		EMPTY:       "_",                            // _ (means fill with random)
		DOTRED:      fmt.Sprintf("%v", red("●")),    // r
		DOTGREEN:    fmt.Sprintf("%v", green("●")),  // g
		DOTBLUE:     fmt.Sprintf("%v", blue("●")),   // b
		DOTYELLOW:   fmt.Sprintf("%v", yellow("●")), // y
		DOTPURPLE:   fmt.Sprintf("%v", purple("●")), // p
		DOTWHITE:    fmt.Sprintf("%v", white("●")),  // w
		DOTWILDCARD: "◯",                            // *
		DOTANCHOR:   "⚓",                            // a
		DOTBOMB:     "💣",                            // o
		ICE0:        "0",
		ICE1:        "1",
		ICE2:        "2",
	}
	if _, ok := m[p]; ok == false {
		warn(fmt.Sprintf("piece2symbol recieved an unknown piece |%v|\n", p))
		return
	}
	s = m[p]
	return
}

func symbol2piece(s string) (p Piece) {
	m := map[string]Piece{"#": NULL,
		"_": EMPTY,       // _ (means fill with random)
		"r": DOTRED,      // r
		"g": DOTGREEN,    // g
		"b": DOTBLUE,     // b
		"y": DOTYELLOW,   // y
		"p": DOTPURPLE,   // p
		"w": DOTWHITE,    // w
		"*": DOTWILDCARD, // *
		"a": DOTANCHOR,   // a
		"o": DOTBOMB,     // o
		"0": ICE0,
		"1": ICE1,
		"2": ICE2,
	}
	if _, ok := m[s]; ok == false {
		warn(fmt.Sprintf("symbol2piece recieved an unknown symbol |%v|\n", s))
		return
	}
	p = m[s]
	return

}

func DotTypeIsSelectable(p Piece) bool {
	if p >= DOTRED && p <= DOTWILDCARD {
		return true
	}
	return false
}
