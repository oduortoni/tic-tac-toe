package ttt

const (
	ORIGINAL = "\033[0m"
	CYAN     = "\033[36m"
	BLACK    = "\033[30m"
	RED      = "\033[31m"
	GREEN    = "\033[32m"
	YELLOW   = "\033[33m"
	BLUE     = "\033[34m"
	MAGENTA  = "\033[35m"
	WHITE    = "\033[37m"
	ORANGE   = "\033[38;5;208m"
    GRAY = "\033[1;30m"
	BCYAN    = "\033[1;36m"
	BRED     = "\033[1;31m"
	BGWHITE   = "\033[47m"
)

func GetColor(name string) (color string) {
	switch name {
	case "cyan":
		color = CYAN
	case "black":
		color = BLACK
	case "red":
		color = RED
	case "green":
		color = GREEN
	case "yellow":
		color = YELLOW
	case "blue":
		color = BLUE
	case "magenta":
		color = MAGENTA
	case "white":
		color = WHITE
	case "orange":
		color = ORANGE
	case "bcyan" :
		color = BCYAN
	case "bred":
		color = BRED
	}
	return
}
