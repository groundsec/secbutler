package data

import "github.com/groundsec/secbutler/pkg/types"

var Wordlists = []types.Wordlist{
	{
		Name:         "SecLists",
		Requirements: []string{"git"},
		InstallCmd:   `git clone --depth 1 https://github.com/danielmiessler/SecLists.git /usr/share/wordlists/SecLists &> /dev/null`,
	},
	{
		Name:         "samlists",
		Requirements: []string{"git"},
		InstallCmd:   `git clone --depth 1 https://github.com/the-xentropy/samlists.git /usr/share/wordlists/samlists &> /dev/null`,
	},
	{
		Name:         "PayloadsAllTheThings",
		Requirements: []string{"git"},
		InstallCmd:   `git clone --depth 1 https://github.com/swisskyrepo/PayloadsAllTheThings.git /usr/share/wordlists/PayloadsAllTheThings &> /dev/null`,
	},
	{
		Name:         "fuzzdb",
		Requirements: []string{"git"},
		InstallCmd:   `git clone --depth 1 https://github.com/fuzzdb-project/fuzzdb.git /usr/share/wordlists/fuzzdb &> /dev/null`,
	},
	{
		Name:         "fuzz.txt",
		Requirements: []string{"git"},
		InstallCmd:   `git clone --depth 1 https://github.com/Bo0oM/fuzz.txt.git /usr/share/wordlists/fuzz.txt &> /dev/null`,
	},
	{
		Name:         "Assetnote",
		Requirements: []string{"wget"},
		InstallCmd:   `wget -r --no-parent -R "index.html*" https://wordlists-cdn.assetnote.io/data/ -nH -e robots=off -P /usr/share/wordlists/Assetnote/ -nd &> /dev/null`,
	},
}
