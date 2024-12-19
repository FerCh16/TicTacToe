package game

import (
	"fmt"
)
// 22
var Title = []string {
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
}

// P1:PLAY | P2:PLAY
const TURN = "P%d: PLAY           " // MODE: PvP | MODE: PvM
const MODE ="MODE: %s          "

func RenderUiByRows(turn int, mode string) []string {
  return append(Title, fmt.Sprintf(TURN, turn), fmt.Sprintf(MODE, mode))
}
