package main

import (
	"flag"

	"github.com/FerCh16/TicTacToe/game"
)

var f = flag.Bool("m", false, "Play with the machine")

func main() {
  flag.Parse()
  var mode string 
  if !(*f) {
    mode = "PvP"
  }else {
    mode = "PvM"
  }
  g := game.NewGame(mode)

  g.Run()
}
