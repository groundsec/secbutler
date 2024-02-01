package utils

var ANSICodes = map[string]string{
	"Reset":             "\033[0m",
	"Bold":              "\033[1m",
	"Dim":               "\033[2m",
	"Italic":            "\033[3m",
	"Underline":         "\033[4m",
	"Blink":             "\033[5m",
	"ReverseVideo":      "\033[7m",
	"Concealed":         "\033[8m",
	"Black":             "\033[30m",
	"Red":               "\033[31m",
	"Green":             "\033[32m",
	"Yellow":            "\033[33m",
	"Blue":              "\033[34m",
	"Magenta":           "\033[35m",
	"Cyan":              "\033[36m",
	"White":             "\033[37m",
	"BlackBackground":   "\033[40m",
	"RedBackground":     "\033[41m",
	"GreenBackground":   "\033[42m",
	"YellowBackground":  "\033[43m",
	"BlueBackground":    "\033[44m",
	"MagentaBackground": "\033[45m",
	"CyanBackground":    "\033[46m",
	"WhiteBackground":   "\033[47m",
}

const Version = "0.1.9"
const MainDirName = ".secbutler"
const PayloadsDirName = "payloads"
const CheatsheetsDirName = "cheatsheets"
const ToolsDirName = "tools"
