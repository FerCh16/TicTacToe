package board

import (
	"errors"
	"strings"
)

const (
	EmptySpace = iota
	P1
	P2
)

type Board [3][3]int

func (b Board) RenderByRows() (s []string) {
	for _, row := range b {
		for i := 0; i < SymbolHeight; i++ {
			var tmp []string
			for _, v := range row {
				tmp = append(tmp, PlayerSymbol[v][i])
			}
			s = append(s, strings.Join(tmp, colSep))
		}
		s = append(s, RowSep)
	}
	return s
}

func (b *Board) Put(v, x, y int) error {
	if (x < 0 || x > 2) || (y < 0 || y > 2) {
		return errors.New("Coords out of range")
	}
	if (*b)[x][y] != EmptySpace {
		return errors.New("The space is not empty")
	}
	(*b)[x][y] = v
	return nil
}

func (b *Board) Status() int {
	// Vertical Case
vertical:
	for i := 0; i < len(b); i++ {
		curr := b[0][i]
		for j := 1; j < len(b[i]); j++ {
			if b[j][i] != curr {
				continue vertical
			}
		}
		if curr != EmptySpace {
			return curr
		}
	}

	// Horizontal Case
horizontal:
	for i := 0; i < len(b); i++ {
		curr := b[i][0]
		for j := 1; j < len(b[i]); j++ {
			if b[i][j] != curr {
				continue horizontal
			}
		}
		if curr != EmptySpace {
			return curr
		}
	}

	curr := b[0][0]
	winner := true
	for i := 0; i < len(b); i++ {
		if curr != b[i][i] {
			winner = false
			break
		}
	}

	if winner && curr != EmptySpace {
		return curr
	}

	winner = true
	for i := len(b) - 1; i > 0; i-- {
		if curr != b[i][i] {
			winner = false
			break
		}
	}

	if winner && curr != EmptySpace {
		return curr
	}

	return -1
}
