package types

type OperatingSystem string

const (
	Linux   OperatingSystem = "Linux"
	Windows OperatingSystem = "Windows"
	Mac     OperatingSystem = "Mac"
)

type RevShell struct {
	Name     string
	OS       []OperatingSystem
	ShellTpl string
}
