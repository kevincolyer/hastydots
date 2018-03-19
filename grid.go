package main

import "fmt"

//import "strings"

type Grid struct {
	cells  []Piece // starts topleft to top right moving to bottom left bottom right in i=y*width+x
	Width  int
	Height int
	Ice    []Piece // to hold ice and water overlays
	Mark   []int   // to make a list of cells to mark as affected by explosions or removals etc. Clean each round
}

//func (g Grid) Mark(x,y int)
//func (g Grid) ClearMarks()
//func (g Grid) NextMark() (x,y int)

//func (g Grid) Empty(x,y) // removes a square and does the essential counting or marking of surrounding squares - used from update grid

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
		return NULL
	}
	return g.cells[x+y*g.Width]
}
func (g Grid) SetGrid(x, y int, p Piece) {
	if g.OnGrid(x, y) == false {
		panic(fmt.Sprintf("Tried to set an off an off grid element %v,%v", x, y))
	}
	g.cells[x+y*g.Width] = p
}

func (g Grid) detectSquare(m []Move) Piece {
	pc := m[0].Colour
	x := m[0].X
	y := m[0].Y
	// 2x2 simple square: if len 4 and from m[0] looking up, down, left or right should give  m[3] to be a square
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

	// 3x3 m[0] +2,+2 = m[4] up down left right = m[8]
	if len(m) == 8 {
		if x+1 == m[7].X && y == m[7].Y && x+2 == m[4].X && y-2 == m[4].Y {
			return pc
		}
		if x+1 == m[7].X && y == m[7].Y && x+2 == m[4].X && y+2 == m[4].Y {
			return pc
		}

		if x-1 == m[7].X && y == m[7].Y && x-2 == m[4].X && y-2 == m[4].Y {
			return pc
		}
		if x-1 == m[7].X && y == m[7].Y && x-2 == m[4].X && y+2 == m[4].Y {
			return pc
		}

		if x == m[7].X && y+1 == m[7].Y && x+2 == m[4].X && y+2 == m[4].Y {
			return pc
		}
		if x == m[7].X && y+1 == m[7].Y && x-2 == m[4].X && y+2 == m[4].Y {
			return pc
		}

		if x == m[7].X && y-1 == m[7].Y && x-2 == m[4].X && y-2 == m[4].Y {
			return pc
		}
		if x == m[7].X && y-1 == m[7].Y && x+2 == m[4].X && y-2 == m[4].Y {
			return pc
		}
	}
	return NULL
}
