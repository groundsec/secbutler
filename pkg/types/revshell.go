package types

type OperatingSystem string

type RevShell struct {
	Name     string
	OS       []string
	ShellTpl string
}

type RevShellListener struct {
	Name        string
	ListenerTpl string
}
