package todo

import "fmt"

const (
	ColorDefault = "\x1b[39m"

	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"
	ColorGray  = "\x1b[90m"
)

func makeRed(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func makeGreen(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}
func makeBlue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}
func makeGray(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGray, s, ColorDefault)
}
