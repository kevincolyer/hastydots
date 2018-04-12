package main

import "fmt"

//import "strings"

type Grid struct {
	Cells  []Piece // starts topleft to top right moving to bottom left bottom right in i=y*width+x
	Width  int
	Height int
}

type GridOffset struct {
	x int
	y int
}

// Clears grid to null
func (g Grid) Clear() {
	//fmt.Printf("In Clear()\ngrid.Clear len=%v\n",len(g.Cells))
	for i, _ := range g.Cells {
		g.Cells[i] = NULL
	}
	//fmt.Printf("grid.Clear()= %v\n",g)
}

// checks if any of the cells are not null
func (g Grid) IsAllNull() bool {
	for _, i := range g.Cells {
		if i != NULL {
			return false
		}
	}
	return true
}

func (g Grid) SetNEWS(x, y int, mark Piece) {
	g.SafeSetGrid(x+1, y, mark)
	g.SafeSetGrid(x-1, y, mark)
	g.SafeSetGrid(x, y+1, mark)
	g.SafeSetGrid(x, y-1, mark)
}

func (g Grid) SetAllNeighbours(x, y int, mark Piece) {
	g.SetNEWS(x, y, mark)
	g.SafeSetGrid(x+1, y+1, mark)
	g.SafeSetGrid(x-1, y+1, mark)
	g.SafeSetGrid(x+1, y-1, mark)
	g.SafeSetGrid(x-1, y-1, mark)
}

func (g Grid) IndexToXY(i int) (x, y int) {
	x = i % g.Width
	y = int(i / g.Height)
	return
}

func (g *Grid) Size(w, h int) {
	g.Width = w
	g.Height = h
	l := w * h
	if len(g.Cells) >= l {
		//debug("grid big enough\n")
		g.Clear()
		return
	}
	g.Cells = append(g.Cells, make([]Piece, l-len(g.Cells))...)
	g.Clear()
	//debug("made bigger grid %v=%v\n",len(g.Cells),g)
	return
}

func (g Grid) String() (s string) {
	s = "|"
	for i, j := range g.Cells {
		s += piece2ascii(j)
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
	return g.Cells[x+y*g.Width]
}

func (g Grid) SetGrid(x, y int, p Piece) {
	if g.OnGrid(x, y) == false {
		panic(fmt.Sprintf("Tried to set an off an off grid element %v,%v", x, y))
	}
	g.Cells[x+y*g.Width] = p
}

// like above but only sets if on grid
func (g Grid) SafeSetGrid(x, y int, p Piece) {
	if g.OnGrid(x, y) {
		g.Cells[x+y*g.Width] = p
	}
}

//copies to self from argument
func (g Grid) Copy(source Grid) {
	for i, _ := range source.Cells {
		g.Cells[i] = source.Cells[i]
	}
	// shouldn't need this bit as set when initialising but for completeness...
	g.Width = source.Width
	g.Height = source.Height
}

// returns a closure that scans right to left from bottom to top!
// because counting to -1 makes more sense than counting away from -1
func (g Grid) makeGridScanner(p Piece, invert bool) func() (int, int, bool) {
	x := g.Width - 1
	y := g.Height
	done := false
	//fmt.Println("made new iterator looking for p starting x and y and invert ", p, x, y-1, invert)
	return func() (int, int, bool) {
		if done {
			debug("iterator exhausted - stop calling me!")
			return -1, -1, done
		}
		// 		fmt.Printf("it x,y=%v,%v\n",x,y)
		for x >= 0 {
			// check here for col exhaust and then row exhaust. both means iterator is exhausted
			for y >= 0 {

				if y == 0 {
					y = g.Height
					x--
					// 					fmt.Printf("y was ==0 so now %v,%v\n",x,y)
					if x == -1 {
						y = -1
						done = true
						//fmt.Printf(" iterator exhuasted 2 %v,%v\n", x, y)
						return x, y, done
					}
				}
				y--

				if invert == false && g.GetGrid(x, y) == p {
					return x, y, done
				}
				if invert == true && g.GetGrid(x, y) != p {
					return x, y, done
				}

			}
		}
		//fmt.Printf(" iterator exhuasted 1 %v,%v\n", x, y)
		//panic("should not reach this!")
		x = -1
		y = -1
		done = true
		return x, y, done
	}
}

func (g Grid) detectSquare(m []Move) Piece {
	pc := m[0].Colour
	x := m[0].X
	y := m[0].Y

	// note to truely square you need to end up on the starting square!!!!!
	if m[0].X != m[len(m)-1].X || m[0].Y != m[len(m)-1].Y {
		return NULL
	}

	// 2x2 simple square: if len  and from m[0] looking up, down, left or right should give  m[3] to be a square
	if len(m) == 5 {
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
	if len(m) == 9 {
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
