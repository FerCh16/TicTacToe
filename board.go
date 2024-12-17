package main

import (
	"errors"
	"strings"
)

type Board [3][3]Cell

func (b *Board) Put(v, x, y int) error {
	if (x < 0 || x > 2) || (y < 0 || y > 2) {
		return errors.New("Coords out of range")
	}

	if (*b)[x][y].Val != EmptySpace {
		return errors.New("The space is not empty")
	}

	(*b)[x][y].Val = v
	return nil
}

func (b Board) RenderBoard() string {
	var builder strings.Builder

	for _, row := range b {
		for h := range cellHeight {
			var tmp []string
			for _, v := range row {
				tmp = append(tmp, v.ToAsciiArt()[h])
			}
			builder.WriteString(strings.Join(tmp, cellSep) + "\n")
		}
		builder.WriteString(rowSep)
	}

	return builder.String()
}
