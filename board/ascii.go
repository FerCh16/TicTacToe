package board

const SymbolHeight = 6

var PlayerSymbol = map[int] []string {
  EmptySpace: {
"          ",
"          ",
"          ",
"          ",
"          ",
"          ",
},
  P1: {
" ██╗  ██╗ ",
" ╚██╗██╔╝ ",
"  ╚███╔╝  ",
"  ██╔██╗  ",
" ██╔╝ ██╗ ",
" ╚═╝  ╚═╝ ",
},
  P2: {
"  ██████╗ ",
" ██╔═══██╗",
" ██║   ██║",
" ██║   ██║",
" ╚██████╔╝",
"  ╚═════╝ ",
},
}

var colSep = "│"
var RowSep = "──────────┼──────────┼──────────"
