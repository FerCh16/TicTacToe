package main

import (
	"errors"
	"strings"
)

const (
  EmptySpace = iota
  Player1
  Player2
)

var (
  ErrNotEmpty = errors.New("The space is not empty")
  OutOfRange = errors.New("Coords out of range")
)

type Table [3][3]int

func (t *Table)Put(v ,x , y int) error {
  if (x < 0 || x > 2)  || (y < 0 || y > 2) {
    return ErrNotEmpty
  }

  if t[x][y] != EmptySpace {
    return ErrNotEmpty
  }

  t[x][y] = v
  return nil
}

func (t *Table) RenderTable() string {
  var builder strings.Builder
  

  return builder.String()
}
