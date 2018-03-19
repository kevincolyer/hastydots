package main

import "fmt"
import "strings"
import "strconv"
import "math/rand"
import "time"

type LevelState struct {
	MoveCounter int
	Pick        []Piece
	Goal        map[Piece]int
	GoalCounter map[Piece]int
	Level       int
	Seed        int
}

// parses a level description to LevelState struct
// example "width 4; height 4; grid rrrr bbbb #yy# ____; pick r3 b3 y g p a10; moves 10; goal a3; seed 1"
func PrepareLevel(level int) (l *LevelState) {
	// parse a level string
	if level < 0 || level >= len(glvl) {
		panic("don't have a level definition for level " + string(level))
	}
	s := strings.Replace(glvl[level], "\n", ";", -1) // conveniece=replace all newlines with ;
	s = strings.Replace(s, ";;", ";", -1)            // conveniece=replace all double ;; with ;
	commands := strings.Split(s, ";")
	debug("Parsing level %v\n%v commands found\nParsing |%q|\n", level, len(commands), commands)

	// DEFAULTS
	grid = Grid{}
	grid.Width = gopt.MaxGridWidth
	grid.Height = gopt.MaxGridHeight
	l = new(LevelState)
	l.Level = level

	rand.Seed(0) // default for a level
	// l.seed is always 0 unless set below

	for i, phrase := range commands {
		phrase = strings.TrimSpace(phrase)
		vn := strings.Fields(phrase)
		debug("%v) parsing %q\n", i, vn)
		if len(vn) < 2 {
			panic("Too few fields in command to parse")
		}

		if vn[0] == "width" {
			grid.Width, _ = strconv.Atoi(vn[1])
		} else if vn[0] == "height" {
			grid.Height, _ = strconv.Atoi(vn[1])
		} else if vn[0] == "moves" {
			l.MoveCounter, _ = strconv.Atoi(vn[1])
		} else if vn[0] == "seed" {
			l.Seed, _ = strconv.Atoi(vn[1])
			rand.Seed(int64(l.Seed))
		} else if vn[0] == "goal" {
			l.Goal = make(map[Piece]int)
			l.GoalCounter = make(map[Piece]int)
			for i := 1; i < len(vn); i++ {
				s, j := a1splitter(vn[i])
				l.Goal[symbol2piece(s)] = j
				l.GoalCounter[symbol2piece(s)] = 0
			}
			debug("goal=%v\n", l.Goal)

		} else if vn[0] == "pick" {
			for i := 1; i < len(vn); i++ {
				s, j := a1splitter(vn[i])
				pc := symbol2piece(s)
				for j > 0 {
					l.Pick = append(l.Pick, pc)
					j--
				}
			}
			debug("pick %v\n", l.Pick)
		} else if vn[0] == "grid" {
			grid.cells = []Piece{}
			// shift out verb
			for i := 1; i < len(vn); i++ {

				for _, k := range vn[i] {
					grid.cells = append(grid.cells, symbol2piece(string(k)))
				}
			}
			// if grid data is malformed this might stop a panic
			for len(grid.cells) < grid.Width*grid.Height {
				warn("malformed level grid data - padding with empty")
				grid.cells = append(grid.cells, EMPTY)
			}
			debug("%#v\n", grid) // to test
		} else {
			warn("level creation data: unrecognised verb " + vn[0])
		}

	}
	debug("%#v\n", l) // to test
	// fill in any empty parts of the grid with random picks
	for k, v := range grid.cells {
		if v == EMPTY {
			pc := l.RandPick()
			// don't fill empty cells with anchors on bottom row...
			for pc == DOTANCHOR && k >= grid.Width*(grid.Height-1) {
				pc = l.RandPick()
			}
			grid.cells[k] = pc
		}
	}
	// refresh random seed for rest of gameplay
	rand.Seed(time.Now().UTC().UnixNano())
	return

}

// update display
func (l *LevelState) Render() (s string) {
	s += fmt.Sprintf("\nHastyDots\tLevel %v\t\tMoves left %v\nGoal ", l.Level+1, l.MoveCounter)
	for k, v := range l.Goal {
		g := v - l.GoalCounter[k]
		if g < 0 {
			g = 0
		}
		s += fmt.Sprintf("%v=%v  ", piece2symbol(k), g)
	}
	s += fmt.Sprintf("\n     ")
	for i := 0; i < grid.Width; i++ {
		s += fmt.Sprintf("%s ", string(i+97))
	}
	for i := 0; i < grid.Height; i++ {

		s += fmt.Sprintf("\n   %v ", string(i+48+1))
		for j := 0; j < grid.Width; j++ {
			s += fmt.Sprintf("%s ", piece2symbol(grid.GetGrid(j, i)))
		}

	}
	s += fmt.Sprintln()
	return
}

func SetupGrid() {

}

func (l *LevelState) UpdateGrid(m []Move) bool {

	// start  by doing players actions.
	// then do the chain reactions and secondary and then tertiaries until all is quiet again.

	// PLAYER MOVES
	// ------------
	if len(m) > 0 {
		// square detection

		for i, _ := range m {
			pc := grid.GetGrid(m[i].X, m[i].Y)
			// update goal
			if _, ok := l.GoalCounter[pc]; ok && pc != DOTWILDCARD {
				l.GoalCounter[pc]++
			}
			grid.SetGrid(m[i].X, m[i].Y, EMPTY)
			// TODO if there is a space next to a bomb - decrement bomb count

		}
		// if square then scan and remove all colour and add a bomb in larger squares
		pc := grid.detectSquare(m) // returns PC or colour of start of square or NULL
		if pc != NULL {
			debug("square detected!")
			// currently only 2x2squares TODO 3x3 with bomb!
			scan := makeGridScanner(pc)
			for x, y, done := scan(); done == false; x, y, done = scan() { // needs to return x and y
				grid.SetGrid(x, y, EMPTY)
				// update goal
				l.GoalCounter[pc]++
			}
		}
		// empty move list (on return)
		return true
	}

	// DROP DOWN
	// ---------
	// scan for empty spaces
	// from empty space walk up. if not null bring down
	// stop when reached the top
	// Then scan down and fill with new
	// if nothing then fill downwards
	// return true
	scan := makeGridScanner(EMPTY)
	finished := false
	for x, y, done := scan(); done == false; x, y, done = scan() {
		fmt.Printf("found empty %v,%v\n", x, y)
		y2 := y - 1
		y3 := y
		// move down
		// TODO skip Nulls!
		for grid.OnGrid(x, y2) {
			if grid.GetGrid(x, y2) == NULL || grid.GetGrid(x, y2) == EMPTY {
				y2--
				continue
			}
			grid.SetGrid(x, y3, grid.GetGrid(x, y2)) // move everything one down
			grid.SetGrid(x, y2, EMPTY)
			y3--
			for grid.GetGrid(x, y3) == NULL {
				y3--
			}
			y2--
		}

		// fill down
		for y4 := 0; y4 <= y; y4++ {
			if grid.GetGrid(x, y4) == EMPTY {
				grid.SetGrid(x, y4, l.RandPick()) // todo could have anchors on last row here is whole colomn empty
			}
		}
		finished = true // have done some work!
	}
	if finished {
		return true
	}

	// scan and if bomb at zero explode, decrementing bombs
	// return true

	// ANCHORS
	// -------
	// scan and if anchors at bottom descend the colomn
	// increment GoalCounter if present

	bottom := grid.Height - 1
	for x := 0; x < grid.Width; x++ {
		if grid.GetGrid(x, bottom) == DOTANCHOR {
			finished = true
			grid.SetGrid(x, bottom, EMPTY)
			if _, ok := l.GoalCounter[DOTANCHOR]; ok {
				l.GoalCounter[DOTANCHOR]++
			}
		}
	}
	if finished {
		return true
	}

	// TODO ICE
	// TODO ladybirds
	// TODO firesquares
	// TODO gems
	// TODO disapearing blocks, triangles, teleports, moves,

	// is there an allowable move for the player - if not shuffle!
	// return true

	return false // nothing to do so grid is stable! Phew!
}

// returns a closure that scans right to left from bottom to top!
// because counting to -1 makes more sense than counting away from -1
func makeGridScanner(p Piece) func() (int, int, bool) {
	x := grid.Width - 1
	y := grid.Height
	done := false
	//fmt.Println("made new iterator looking for p starting x and y ", p, x, y-1)
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
					y = grid.Height
					x--
					// 					fmt.Printf("y was ==0 so now %v,%v\n",x,y)
					if x == -1 {
						y = -1
						done = true
						// 						fmt.Printf(" iterator exhuasted 2 %v,%v\n", x, y)
						return x, y, done
					}
				}
				y--
				if grid.GetGrid(x, y) == p {
					//                                         fmt.Println("got a dot!", x, y)
					return x, y, done
				}
			}
		}
		// 		fmt.Printf(" iterator exhuasted 1 %v,%v", x, y)
		panic("should not reach this!")
		// 		x = -1; y = -1 ; done=true ;return x, y, done
	}
}

func GridNeedsUpdating() {

}

// helper functions (may not need)
// picks a number between 0 and n-1. Returns n
func (l *LevelState) RandPick() Piece {
	if len(l.Pick) == 0 {
		panic("I am asked to pick from l.Pick and there is nothing to pick from!")
	}
	return l.Pick[rand.Intn(len(l.Pick))]
}

func Shuffle() {

}
