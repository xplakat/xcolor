package xcolor

import "fmt"

const Reset = "\033[0m"

const Black = "\033[30m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Magenta = "\033[35m"
const Cyan = "\033[36m"
const White = "\033[37m"

const BlackLgt = "\033[90m"
const RedLgt = "\033[91m"
const GreenLgt = "\033[92m"
const YellowLgt = "\033[93m"
const BlueLgt = "\033[94m"
const MagentaLgt = "\033[95m"
const CyanLgt = "\033[96m"
const WhiteLgt = "\033[97m"

func Paint(color, text string) string {
	return fmt.Sprint(color + text + Reset)
}

func PrintPaint(color, text string) {
	fmt.Print(Paint(color, text))
}

func PrintlnPaint(color, text string) {
	fmt.Println(Paint(color, text))
}

func RedString(text string) string {
	return fmt.Sprint(Paint(Red, text))
}

func GreenString(text string) string {
	return fmt.Sprint(Paint(Green, text))
}

func YellowString(text string) string {
	return fmt.Sprint(Paint(Yellow, text))
}

func BlueString(text string) string {
	return fmt.Sprint(Paint(Blue, text))
}

func MagentaString(text string) string {
	return fmt.Sprint(Paint(Magenta, text))
}

func CyanString(text string) string {
	return fmt.Sprint(Paint(Cyan, text))
}

func WhiteString(text string) string {
	return fmt.Sprint(Paint(White, text))
}

func BlackLgtString(text string) string {
	return fmt.Sprint(Paint(BlackLgt, text))
}
