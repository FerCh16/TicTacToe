package game

import (
	"fmt"
)

// 22
var Title = []string{
	`                   `,
	`                   `,
	` _____ ___ ____    `,
	`|_   _|_ _/ ___|   `,
	`  | |  | | |       `,
	`  | |  | | |___    `,
	` _|_|_|___\____|_  `,
	`|_   _|/ \  / ___| `,
	`  | | / _ \| |     `,
	`  | |/ ___ \ |___  `,
	` _|_/_/___\_\____| `,
	`|_   _/ _ \| ____| `,
	`  | || | | |  _|   `,
	`  | || |_| | |___  `,
	`  |_| \___/|_____| `,
	`                   `,
	`                   `,
	`                   `,
	`                   `,
	`                   `,
}

// P1:PLAY | P2:PLAY
const TURN = "P%d: PLAY           " // MODE: PvP | MODE: PvM

func RenderUiByRows(turn int) []string {
	return append(Title, fmt.Sprintf(TURN, turn))
}
