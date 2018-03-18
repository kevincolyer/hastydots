package main

import "fmt"

//import "strings"

type Grid struct {
	cells  []Piece // starts topleft to top right moving to bottom left bottom right in i=y*width+x
	Width  int
	Height int
}

func (g Grid) String() (s string) {
	s = "|"
	for i, j := range g.cells {
		s += piece2symbol(j)
		if (i+1)%g.Width == 0 {
			s += "|"
		}
	}
	return s
}

func (g Grid) OnGrid(x, y int) bool {
	if x < 0 || y < 0 || x >= g.Width || y >= g.Height {
		return false
	}
	return true
}

func (g Grid) GetGrid(x, y int) Piece {
	if g.OnGrid(x, y) == false {
		panic(fmt.Sprintf("Tried to get an off grid element %v,%v", x, y))
	}
	return g.cells[x+y*g.Width]
}
func (g Grid) SetGrid(x, y int, p Piece) {
	if g.OnGrid(x, y) == false {
		panic(fmt.Sprintf("Tried to set an off an off grid element %v,%v", x, y))
	}
	g.cells[x+y*g.Width] = p
}

// func
func (g Grid) detectSquare(m []Move) Piece {
    return NULL
    /* .urd
     * .uld
     * .dlu
     * .dru
     * .ldr
     * .lur
     * .rdr
     * .rur
    */
}
