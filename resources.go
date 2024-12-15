package main

var (
  cellSep = "│"
  rowSep = "──────────────┼──────────────┼──────────────"
)

const (
  cellWidht = 14
  cellHeight = 6
  cellArea = cellWidht * cellHeight
)

var PlayerSymbol = map[int] []string {
  EmptySpace: nil,
  Player1: 
{
"   ██╗  ██╗   ",
"   ╚██╗██╔╝   ",
"    ╚███╔╝    ",
"    ██╔██╗    ",
"   ██╔╝ ██╗   ",
"   ╚═╝  ╚═╝   ",
},

Player2:
{
"   ██████╗    ",
"  ██╔═══██╗   ",
"  ██║   ██║   ",
"  ██║   ██║   ",
"  ╚██████╔╝   ",
"   ╚═════╝    ",
},

}

var TableLayout [][][]string
