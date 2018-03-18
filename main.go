package main

import "fmt"
import "strings"

// import "strconv"
// import "math/rand"
// import "time"
import "bufio"
import "os"

var dbg bool // mostly allows printing to stderr for debug() func (a wrapper around fmt.Printf)

// Global structs

type GameOpt struct {
	MaxGridWidth  int
	MaxGridHeight int
}

type GameState struct {
	Points    int
	Name      string
	HighScore int
	Level     int
}

type PlayerEvents int

const (
	PLAYERPLAYING PlayerEvents = iota + 1
	PLAYERQUITS
	PLAYERWINSLEVEL
	PLAYERLOSESLEVEL
	PLAYERRESTARTSLEVEL
	PLAYERMAKESMOVE
	PLAYERMAKESILLEGALMOVE
	PLAYERMAKESSQUARE = 32
)

// GLOBALS
var gopt GameOpt
var gste GameState
var glvl []string
var grid Grid

// End GLOBALS

func init() {
	gste = GameState{Name: "HastyDots", HighScore: 10}
	gopt = GameOpt{MaxGridHeight: 8, MaxGridWidth: 7}
	glvl = []string{
            "width 4; height 4; grid rrrr *bbb #*y# ____; pick r3 b3 y g p a10; moves 10; goal a3; seed 1",
            
            "width 8; height 8; grid ________ #______# ________ #______# ___##___ #__##__# ________ #______# ________ #______#; pick r4 b4 g4 p4 w4 g4 * a; moves 20; goal r10 b10 g10 a5;seed 2 ",
            
        }
}

// mainloop
func main() {
	//
	dbg = true
	//
	// prepare a level
	// l := PrepareLevel(0)  // level 0 is test level
        l := PrepareLevel(1)
	// GameLoop
	gameloop := PLAYERPLAYING
	for gameloop == PLAYERPLAYING {
		// display level
		output("Welcome to HastyDots\n")
		output("Level %v\n", l.Level)
		output(l.Render())
		//  accept player input
		input := UserInput()
		if input == "quit" {
			gameloop = PLAYERQUITS
			output("Player quits - bye!\n")
		}
		if input == "restart" {
			gameloop = PLAYERRESTARTSLEVEL
			output("Player restarts\n")
		}
		//  validate player input
		if moves, err := PlayerInputOk(input); err != false {
			// gameloop=PLAYERMAKESILLEGALMOVE
			// gameloop=PLAYERMAKESMOVE
			// gameloop=PLAYERMAKESSQUARE+PLAYERMAKESMOVE
			//output("%v\n", moves)
			//  loop UpdateGrid until GridNeedsUpdating == False
			for l.UpdateGrid(moves) == true{
				if len(moves) > 0 {
					moves = []Move{}
				} // empty moves list as first pass to Updategrid will use them.
				 output(l.Render())

				// pause a while?
				// why not!
			}
			// loop through all goalcounters. Make zero if <0. add and if all = 0 PLAYERWINSLEVEL
			// gameloop=PLAYERWINS

			//  decrease turn counter
			l.MoveCounter--
			if gameloop != PLAYERWINSLEVEL && l.MoveCounter == 0 {
				gameloop = PLAYERLOSESLEVEL
			}
			
		}
		if gameloop == PLAYERMAKESILLEGALMOVE || gameloop&PLAYERMAKESMOVE == 1 {
			gameloop = PLAYERPLAYING
		}
	} // end gameloop
	// process gameloop results
	// GameLoop until player quits or out of moves or acheives goal
	// if acheives goal then move to next level levelloop
	// if failed restart levelloop
	// quit

}

func (l *LevelState) LevelLoop() (pe PlayerEvents) {
	return
}

func GameLoop() {

}

// user input - returns nil or valid user input
func UserInput() (s string) {
	fmt.Printf("Enter start coordinates then u,d,r,l as many times as you wish. Type quit to leave and restart to begin again\n>")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	s = strings.TrimSpace(s)
	debug(s)

	return
}

func PlayerInputOk(input string) (moves []Move, err bool) {
	err = false // false  until we validate a correct input
	if len(input) < 3 {
		debug("input too short - need two dots! %v \n", input)
		return
	} // have to have at least coord and a direction
	x := int(input[0]) - 97 // a b
	y := int(input[1]) - 49
	if grid.OnGrid(x, y) == false {
		debug("not on grid %v \n", input)
		return
	}
	debug("x=%v,y=%v\n", x, y)
	pc := grid.GetGrid(x, y)
	rawpc := pc & 0xf //mask off lower bits of piece
	debug("%v\n", rawpc)
	if rawpc < DOTBLUE || rawpc > DOTWILDCARD {
		debug("choice bad %v \n", input)
		return
	}
	debug("choice ok %v\n", input)
	chosencolour := rawpc
	// u=117 d=100 l=108 r=114

	moves = append(moves, Move{X: x, Y: y, Colour: pc})
        debug("adding move x,y,colour %v,%v,%v\n",x,y,pc)
		
	i := 0
	for j := 2; j < len(input); j++ {
		i++
		ascii := int(input[j])
		if ascii == 117 {
			y--
		} else if ascii == 100 {
			y++
		} else if ascii == 108 {
			x--
		} else if ascii == 114 {
			x++
		} else {
			debug("unknown command %v\n", input[i])
			return
		}
		if grid.OnGrid(x, y) == false {
			debug("not on grid %v\n", input)
			return
		}
		pc = grid.GetGrid(x, y)
                
		rawpc = pc & 0xf // mask off lower bits of piece
		if rawpc < DOTBLUE || rawpc > DOTWILDCARD {
			debug("not a choosable dot %v \n", input)
			return
		}
		debug("chosencolour=%v,rawpc=%v\n", chosencolour, rawpc)
		if chosencolour == DOTWILDCARD && rawpc != DOTWILDCARD {
			chosencolour = rawpc
		}
		if  rawpc != chosencolour && rawpc!=DOTWILDCARD  {
			debug("not of same colour %v \n", input)
			return
		}
		debug("adding move x,y,colour %v,%v,%v\n",x,y,pc)
		moves = append(moves, Move{X: x, Y: y, Colour: pc})
	}
	err = true
	moves[0].Colour = chosencolour // first colour indicates all colours of range  as a valid range is only one colour
	debug("%v:all ok: moves=%v\n", input, moves)
	return
}
