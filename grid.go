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
	pc := m[0].Colour
	x := m[0].X
	y := m[0].Y
	if len(m) == 4 {
		if x+1 == m[3].X && y == m[3].Y {
			return pc
		}
		if x-1 == m[3].X && y == m[3].Y {
			return pc
		}
		if x == m[3].X && y+1 == m[3].Y {
			return pc
		}
		if x == m[3].X && y-1 == m[3].Y {
			return pc
		}
	}
	return NULL
	// 2x2 simple square: if len 4 and from m[0] looking up, down, left or right should give  m[3] to be a square

	// 3x3 m[0] +2,+2 = m[4] up down left right = m[8]
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
