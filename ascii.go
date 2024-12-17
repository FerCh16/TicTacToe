package main

var (
	cellWidth  = 14
	cellHeight = 6

	cellSep = "│"
	rowSep  = "──────────────┼──────────────┼──────────────\n"
)

var PlayerSymbol = map[int][]string{
	EmptySpace: {
		"              ",
		"              ",
		"              ",
		"              ",
		"              ",
		"              ",
	},
	Player1: {
		"   ██╗  ██╗   ",
		"   ╚██╗██╔╝   ",
		"    ╚███╔╝    ",
		"    ██╔██╗    ",
		"   ██╔╝ ██╗   ",
		"   ╚═╝  ╚═╝   ",
	},

	Player2: {
		"   ██████╗    ",
		"  ██╔═══██╗   ",
		"  ██║   ██║   ",
		"  ██║   ██║   ",
		"  ╚██████╔╝   ",
		"   ╚═════╝    ",
	},
}
