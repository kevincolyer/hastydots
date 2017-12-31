package main

import "fmt"
import "strings"
import "strconv"
import "math/rand"
import "time"
import "bufio"
import "os"

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

type LevelState struct {
	MoveCounter int
	Pick        []Piece
	Goal        map[Piece]int
	GoalCounter map[Piece]int
	Level       int
	Seed        int
}


type Move struct {
    Colour  Piece
    X   int
    Y   int
}

type Piece int

const (
	NULL        Piece = iota + 1 // #
	EMPTY                        // _ (means fill with random)

        DOTBLUE                      // b
	DOTGREEN                     // g
	DOTPURPLE                    // p
	DOTRED                       // r
	DOTWHITE                     // w
	DOTYELLOW                    // y
        DOTWILDCARD                  // *

        DOTANCHOR                    // a
	DOTBOMB                      // o

	ICE0        = 32
	ICE1        = 64
	ICE2        = 128
)

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
	glvl = []string{"width 4; height 4; grid rrrr *bbb #*y# ____; pick r3 b3 y g p a10; moves 10; goal a3; seed 1"}
}

// mainloop
func main() {
	fmt.Println("Welcome to HastyDots")
//         scanner := bufio.NewScanner(os.Stdin)
	// prepare display
	// initialise gamestate
	// ADVANCE restore a game state
	// levelloop
	// prepare a level
	l := PrepareLevel(0)
	fmt.Println("starting level ", l.Level)
        
	// GameLoop
        gameloop:=PLAYERPLAYING
	for gameloop==PLAYERPLAYING {
            // display level
            l.Render()
            //  accept player input
            input:=UserInput()
            if input=="quit" { gameloop=PLAYERQUITS 
                fmt.Println("Player quits - bye!")}
            if input=="restart" { gameloop=PLAYERRESTARTSLEVEL
                fmt.Println("Player restarts")}
            //  validate player input
            if moves,err:=PlayerInputOk(input); err!=false {
                // gameloop=PLAYERMAKESILLEGALMOVE 
                // gameloop=PLAYERMAKESMOVE
                // gameloop=PLAYERMAKESSQUARE+PLAYERMAKESMOVE
                fmt.Println(moves)
                //  loop UpdateGrid until GridNeedsUpdating == False
                
                // loop through all goalcounters. Make zero if <0. add and if all = 0 PLAYERWINSLEVEL
                // gameloop=PLAYERWINS
                
                //  decrease turn counter
                l.MoveCounter--
                if gameloop!=PLAYERWINSLEVEL && l.MoveCounter==0 { gameloop=PLAYERLOSESLEVEL }
                gameloop=PLAYERLOSESLEVEL
            }
            if gameloop==PLAYERMAKESILLEGALMOVE || gameloop&PLAYERMAKESMOVE==1 { gameloop= PLAYERPLAYING }
        }   // end gameloop
        // process gameloop results
	// GameLoop until player quits or out of moves or acheives goal
	// if acheives goal then move to next level levelloop
	// if failed restart levelloop
	// quit

}

// parses a level description to LevelState struct
// example "width 4; height 4; grid rrrr bbbb #yy# ____; pick r3 b3 y g p a10; moves 10; goal a3; seed 1"
func PrepareLevel(level int) (l *LevelState) {
	// parse a level string
	if level < 0 || level >= len(glvl) {
		panic("don't have a level definition for level " + string(level))
	}
	s := strings.Replace(glvl[level], "\n", ";", -1) // convenieve=replace all newlines with ;
	s = strings.Replace(s, ";;", ";", -1)            // convenieve=replace all double ;; with ;
	commands := strings.Split(s, ";")
	fmt.Printf("Parsing level %v\n%v commands found\nParsing |%q|\n", level, len(commands), commands)

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
		fmt.Printf("%v) parsing %q\n", i, vn)
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
			fmt.Printf("goal=%v\n", l.Goal)

		} else if vn[0] == "pick" {
			for i := 1; i < len(vn); i++ {
				s, j := a1splitter(vn[i])
				pc := symbol2piece(s)
				for j > 0 {
					l.Pick = append(l.Pick, pc)
					j--
				}
			}
			fmt.Printf("pick %v\n", l.Pick)
		} else if vn[0] == "grid" {
			grid.cells = []Piece{}
			// shift out verb
			for i := 1; i < len(vn); i++ {
				for _, k := range vn[i] {
					grid.cells = append(grid.cells, symbol2piece(string(k)))
				}
			}
			fmt.Printf("%#v\n", grid) // to test
		} else {
			fmt.Println("unrecognised verb " + vn[0])
		}

	}
	fmt.Printf("%#v\n", l) // to test
	// fill in any empty parts of the grid with random picks
	for k, v := range grid.cells {
		if v == EMPTY {
			grid.cells[k] = l.RandPick()
		}
	}
	// refresh random seed for rest of gameplay
	rand.Seed(time.Now().UTC().UnixNano())
	return

}

// splits strings in two of form [.\d*]. Second string is coerced to int value or 1 if missing/err
func a1splitter(s string) (l string, i int) {
	arr := strings.SplitN(s, "", 2) // split into two parts
	l = arr[0]
	i = 1
	if len(arr) == 1 {
		return
	}
	if j, err := strconv.Atoi(arr[1]); err == nil {
		i = j
	}
	return
}

func (l *LevelState) LevelLoop() (pe PlayerEvents) {
	return
}

func GameLoop() {

}

// update display
func (l *LevelState) Render() {
	fmt.Printf("\nHastyDots\tLevel %v\t\tMoves left %v\tGoal ", l.Level+1, l.MoveCounter)
		fmt.Println()
	for k, v := range l.Goal {
		fmt.Printf("%v=%v  ", piece2symbol(k), v-l.GoalCounter[k])
	}
	fmt.Printf("\n     ")
	for i := 0; i < grid.Width; i++ {
		fmt.Printf("%s ", string(i+97))
	}
	for i := 0; i < grid.Height; i++ {

		fmt.Printf("\n   %v ", string(i+48+1))
		for j := 0; j < grid.Width; j++ {
			fmt.Printf("%v ", piece2symbol(grid.GetGrid(j, i)))
		}

	}
		fmt.Println()

}

// user input - returns nil or valid user input
func UserInput() (s string) {
    fmt.Printf("Enter start coordinates then u,d,r,l as many times as you wish. Type quit to leave and restart to begin again\n>")
    reader := bufio.NewReader(os.Stdin)
    s, _ = reader.ReadString('\n')
    s=strings.TrimSpace(s)
    fmt.Println(s)
    
    return
}


func PlayerInputOk(input string) (moves []Move, err bool) {
    err=false // false  until we validate a correct input
    if len(input)<3 { fmt.Println("input too short - need two dots!",input); return } // have to have at least coord and a direction
    x:=int(input[0])-97 // a b
    y:=int(input[1])-49
    if grid.OnGrid(x,y)==false { fmt.Println("not on grid",input);return }
    fmt.Printf("x=%v,y=%v\n",x,y)
    pc:=grid.GetGrid(x,y)
    rawpc:=pc&(ICE0-1)
    fmt.Println(rawpc)
    if rawpc<DOTBLUE || rawpc > DOTWILDCARD {  fmt.Println("choice bad",input) ; return }
    fmt.Println("choice ok",input)
    chosencolour:=rawpc
    // u=117 d=100 l=108 r=114
    moves=append(moves,Move{X:y,Y:y,Colour:pc})
    i:=0
    for j:=2;j<len(input);j++ {
        i++
        ascii:=int(input[j])
        if ascii==117 { y--
            
        } else if ascii==100 { y++
            
        } else if ascii==108 { x--
            
        } else if ascii==114 { x++
            
        } else { fmt.Println("unknown command ",input[i]); return }
        if grid.OnGrid(x,y)==false { fmt.Println("not on grod",input);return }
        pc=grid.GetGrid(x,y)
        rawpc=pc&(ICE0-1)
        if rawpc<DOTBLUE || rawpc > DOTWILDCARD {  fmt.Println("not a choosable dot",input) ; return }
        fmt.Printf("chosencolour=%v,rawpc=%v\n",chosencolour,rawpc)
        if chosencolour==DOTWILDCARD && rawpc != DOTWILDCARD { chosencolour = rawpc}
        if chosencolour!=rawpc { fmt.Println("not of same colour",input); return }
        moves=append(moves,Move{X:x,Y:y,Colour:pc})
    }
    err=true
    moves[0].Colour=chosencolour // first colour indicates all colours of range  as a valid range is only one colour
    fmt.Printf("%v:all ok: moves=%v\n",input,moves);
    return 
}

func SetupGrid() {

}

func UpdateGrid() {

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

func (g Grid) OnGrid(x, y int) bool {
	if x < 0 || y < 0 || x >= g.Width || y >= g.Height {
		return false
	}
	return true
}

func (g Grid) GetGrid(x, y int) Piece {
	if g.OnGrid(x, y) == false {
		panic("Tried to get an off grid element")
	}
	return g.cells[x+y*g.Width]
}
func (g Grid) SetGrid(x, y int, p Piece) {
	if g.OnGrid(x, y) == false {
		panic("Tried to set an off grid element")
	}
	g.cells[x+y*g.Width] = p
}

func piece2symbol(p Piece) (s string) {
	m := map[Piece]string{NULL: "#",
		EMPTY:       "_", // _ (means fill with random)
		DOTRED:      "r", // r
		DOTGREEN:    "g", // g
		DOTBLUE:     "b", // b
		DOTYELLOW:   "y", // y
		DOTPURPLE:   "p", // p
		DOTWHITE:    "w", // w
		DOTWILDCARD: "*", // *
		DOTANCHOR:   "a", // a
		DOTBOMB:     "o", // o
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

func warn(s string) {
	fmt.Println("WARNING:" + s)
}
