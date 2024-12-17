package main

const (
	EmptySpace = iota
	Player1
	Player2
)

type Cell struct {
	Val int
}

func (c Cell) ToAsciiArt() []string {
	return PlayerSymbol[c.Val]
}
