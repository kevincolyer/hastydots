package main

import "fmt"

type Cell int
type Grid struct {
	cells  []Cell
	Width  int
	Height int
}

type GameOpt struct {
	MaxGridWidth  int
	MaxGridHeight int
}

type GameState struct {
	Points      int
	Name        string
	HighScore   int
	Level       int
}

type LevelState struct {
	MoveCounter int
	Pick        []Cell
	Goal        map[piece]int
	GoalCounter map[piece]int
	Level       int
}

type piece int

const  (
	NULL piece = iota + 1 // #
	EMPTY          // _ (means fill with random)
	DOTRED         // r
	DOTGREEN       // g
	DOTBLUE        // b
	DOTYELLOW      // y
	DOTPURPLE      // p
	DOTWHITE       // w
	DOTWILDCARD    // *
	DOTANCHOR      // a
	DOTBOMB        // o
	ICE0 = 32
	ICE1 = 64
	ICE2 = 128
)

type PlayerEvents int

const  (
    PLAYERQUITS  PlayerEvents = iota + 1
    PLAYERWINSLEVEL 
    PLAYERLOSESLEVEL 
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
	glvl=[]string{"width 4; height 4; grid rrrr bbbb #yy# ____; pick r3 b3 y g p a10; count 10; goal a3"}
}

// mainloop
func main() {
	fmt.Println("Welcome to HastyDots")
	// prepare display
	// initialise gamestate
	// ADVANCE restore a game state
	// levelloop
        // prepare a level
	// display level
	// GameLoop
	//  accept player input
	//  validate player input
	//  loop UpdateGrid until GridNeedsUpdating == False
	//  decrease turn counter
	// GameLoop until player quits or out of moves or acheives goal
        // if acheives goal then move to next level levelloop
        // if failed restart levelloop
        // quit
        

}

func PrepareLevel(level int) (l *LevelState) {
    // parse a level string
    // stuff it into grid variable
    return
    
}

func (l *LevelState) LevelLoop() (pe PlayerEvents) {
    return
}

func GameLoop () {
    
}

// update display
func Render() {

}

// user input - returns nil or valid user input
func UserInput() {

}

func SetupGrid() {

}

func UpdateGrid() {

}

func GridNeedsUpdating() {

}

// helper functions (may not need)
// picks a number between 0 and n-1. Returns n
func Pick() {

}

func Shuffle() {

}

func (g Grid) OnGrid(x, y int) bool {
	if x < 0 || y < 0 || x >= g.Width || y >= g.Height {
		return false
	}
	return true
}

func (g Grid) GetGrid(x, y int) Cell {
	if g.OnGrid(x, y) == false {
		panic("Tried to get an off grid element")
	}
	return g.cells[x+y*g.Width]
}
func (g Grid) SetGrid(x, y int, c Cell) {
	if g.OnGrid(x, y) == false {
		panic("Tried to set an off grid element")
	}
	g.cells[x+y*g.Width] = c
}

func piece2symbol(p piece) (s string) {
    	m:=map[piece]string{NULL: " ",
	EMPTY: "_",          // _ (means fill with random)
	DOTRED: "r",        // r
	DOTGREEN: "g",       // g
	DOTBLUE: "b",        // b
	DOTYELLOW: "y",      // y
	DOTPURPLE: "p",      // p
	DOTWHITE: "w",       // w
	DOTWILDCARD: "*",    // *
	DOTANCHOR: "a",      // a
	DOTBOMB: "o",        // o
	ICE0: "0",
	ICE1: "1",
	ICE2: "2",
        }
        if _,ok:=m[p]; ok==false { warn(fmt.Sprintf("piece2symbol recieved an unknown piece |%v|\n",p))
            return }
        s=m[p]
        return
}

func symbol2piece(s string) (p piece) {
    	m:=map[string]piece{" ": NULL,
	"_": EMPTY,          // _ (means fill with random)
	"r": DOTRED,        // r
	"g": DOTGREEN,       // g
	"b": DOTBLUE,        // b
	"y": DOTYELLOW,      // y
	"p": DOTPURPLE,      // p
	"w": DOTWHITE,       // w
	"*": DOTWILDCARD,    // *
	"a": DOTANCHOR,      // a
	"o": DOTBOMB,        // o
	"0": ICE0,
	"1": ICE1,
	"2": ICE2,
        }
        if _,ok:=m[s]; ok==false { warn(fmt.Sprintf("symbol2piece recieved an unknown symbol |%v|\n",s))
            return }
        p=m[s]
        return
    
}

func warn(s string) {
    fmt.Println("WARNING:"+s)
}
