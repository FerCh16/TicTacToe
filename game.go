package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	board Board
}

func (g *Game) Run() {
	// Game loop
	out, player1 := os.Stdout, true
	for ; ; player1 = !player1 {
		// Clear screen
		out.WriteString("\033[H\033[2J")

		out.WriteString(g.board.RenderBoard())

		var input string
		fmt.Println("Enter coord: x,y")
		fmt.Scanf("%s", &input)
		tmp := strings.Split(input, ",")
		x, errParse := strconv.ParseInt(tmp[0], 10, 64)
		y, errParse := strconv.ParseInt(tmp[1], 10, 64)
		if errParse != nil {
			fmt.Println("INVALID INPUT")
			fmt.Println(errParse)
			fmt.Scanf("%s", &input)
			continue
		}

		if player1 {
			g.board.Put(Player1, int(x-1), int(y-1))
		}
		g.board.Put(Player2, int(x-1), int(y-1))
	}
}
