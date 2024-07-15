package main

import (
	"fmt"

	colors "github.com/ksrof/gocolors"
)

const (
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

func ColorizeRed(a string) string {
	return Red + a + colors.Reset
}

func ColorizeGreen(a string) string {
	return Green + a + colors.Reset
}

func ColorizeBlue(a string) string {
	return Blue + a + colors.Reset
}

func ColorizeYellow(a string) string {
	return Yellow + a + colors.Reset
}

func ColorizeMagenta(a string) string {
	return Magenta + a + colors.Reset
}

func ColorizeCyan(a string) string {
	return Cyan + a + colors.Reset
}

func ColorizeWhite(a string) string {
	return White + a + colors.Reset
}

func RGB(r, g, b int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

func ColorizeCustom(a string, r, g, b uint8) string {
	return RGB(int(r), int(g), int(b)) + a + colors.Reset
}

func main() {
	fmt.Println(ColorizeRed("GOLANG"))
	fmt.Println(ColorizeGreen("GOLANG"))
	fmt.Println(ColorizeBlue("GOLANG"))
	fmt.Println(ColorizeYellow("GOLANG"))
	fmt.Println(ColorizeMagenta("GOLANG"))
	fmt.Println(ColorizeCyan("GOLANG"))
	fmt.Println(ColorizeWhite("GOLANG"))
	fmt.Println(ColorizeCustom("GOLANG", 150, 150, 100))
}
