package game

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	. "github.com/FerCh16/TicTacToe/board"
)

type Game struct {
  turn int // P1 or P2
  mode string // PvP or PvM
 
  board *Board
}

func NewGame(mode string) *Game {
  return &Game{turn: 1, mode: mode, board: new(Board)}
}

func (g *Game) Run() {
 
  count := 0
  // game loop
  for {
    // Clear term and render game
    fmt.Println("\033[H\033[2J")
    fmt.Println(g.RenderGame())

    if w := g.board.Status() ; w != -1 {
      fmt.Printf("P%d WINS!!!!!!\n", w)
      break; 
    }else if count == 9 {
      fmt.Printf("TIE GAME!\n")
      break; 
    }

    // Get Input
    var input string
    fmt.Println("ENTER COORDS: x,y")
    fmt.Print("→ ")
    fmt.Scanf("%s", &input) 
  
    s := strings.Split(input, ",")
    if len(s) != 2 {
      HandleError(errors.New("Invalid input"))
      continue
    }

    x, errParse := strconv.ParseInt(s[0], 10, 64)

    if errParse != nil {
      HandleError(errParse)
      continue 
    }

    y, errParse := strconv.ParseInt(s[1],10, 64)

    if errParse != nil {
      HandleError(errParse)
      continue 
    }

    errPut := g.board.Put(g.turn, int(x-1), int(y-1))
  
    if errPut != nil {
      HandleError(errPut)
      continue 
    }

    g.turn ^= 3
    count++ 
  }
}

func HandleError(err error) {
  fmt.Println(err)
  // wait to the user press enter
  fmt.Scanf("%s", nil)
}

func (g *Game) RenderGame() string {
  ui := RenderUiByRows(g.turn, g.mode)
  board := g.board.RenderByRows()

  var builder strings.Builder

  for i := 0 ; i < len(board) ; i++ {
    builder.WriteString(strings.Join([]string{ui[i], board[i]}, " ║ ") + "\n")
  }

  return builder.String()
}
