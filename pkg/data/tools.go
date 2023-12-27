package data

import "github.com/groundsec/secbutler/pkg/types"

var SecTools = []types.SecTool{
	{
		Name:         "Amass",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/owasp-amass/amass/v4/...@master`,
	},
	{
		Name:         "Chaos Client",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/chaos-client/cmd/chaos@latest`,
	},
	{
		Name:         "shuffleDNS",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/shuffledns/cmd/shuffledns@latest`,
	},
	{
		Name:         "As3nt",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"python3", "pip"},
		InstallCmd:   `pip install as3nt`,
	},
	{
		Name:         "AltDNS",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"python3", "pip3"},
		InstallCmd:   `pip3 install py-altdns==1.0.2`,
	},
	{
		Name:         "dnsx",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/dnsx/cmd/dnsx@latest`,
	},
	{
		Name:         "hackrevdns",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/hakluke/hakrevdns@latest`,
	},
	{
		Name:         "knock",
		Category:     "",
		Subcategory:  "",
		Requirements: []string{"git", "python3"},
		InstallCmd: `
		git clone https://github.com/guelfoweb/knock.git
		cd knock
		python3 setup.py install
		cd ..
		`,
	},
	{
		Name:         "Subfinder",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest`,
	},
	{
		Name:         "Assetfinder",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/tomnomnom/assetfinder@latest`,
	},
	{
		Name:         "VHostScan",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"git", "python3"},
		InstallCmd: `
		git clone https://github.com/codingo/VHostScan.git
		cd VHostScan
		python3 setup.py install
		cd ..
		`,
	},
	{
		Name:         "Scilla",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"git", "go"},
		InstallCmd: `
		git clone https://github.com/edoardottt/scilla.git
		cd scilla
		make linux
		cd ..
		`,
	},
	{
		Name:         "cero",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/glebarez/cero@latest`,
	},
	{
		Name:         "masscan",
		Category:     "Recon",
		Subcategory:  "Port Scanning",
		Requirements: []string{"make", "git", "gcc"},
		InstallCmd: `
		git clone https://github.com/robertdavidgraham/masscan
		cd masscan
		make
		make install
		cd ..
		`,
	},
	{
		Name:         "RustScan",
		Category:     "Recon",
		Subcategory:  "Port Scanning",
		Requirements: []string{"cargo"},
		InstallCmd: `
		git clone https://github.com/RustScan/RustScan.git
		cd RustScan
		cargo build --release
		cd ..
		`,
	},
	{
		Name:         "Naabu",
		Category:     "Recon",
		Subcategory:  "Port Scanning",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/naabu/v2/cmd/naabu@latest`,
	},
	{
		Name:         "Nmap",
		Category:     "Recon",
		Subcategory:  "Port Scanning",
		Requirements: []string{"git", "make", "gcc"},
		InstallCmd: `
		git clone https://github.com/nmap/nmap.git
		cd nmap
		./configure
		make
		make install
		cd ..
		`,
	},
	{
		Name:         "Sandmap",
		Category:     "Recon",
		Subcategory:  "Port Scanning",
		Requirements: []string{},
		InstallCmd: `
		git clone --recursive https://github.com/trimstray/sandmap
		cd sandmap
		./setup.sh install
		sandmap
		`,
	},
	{
		Name:         "",
		Category:     "",
		Subcategory:  "",
		Requirements: []string{},
		InstallCmd:   ``,
	},
	{
		Name:         "",
		Category:     "",
		Subcategory:  "",
		Requirements: []string{},
		InstallCmd:   ``,
	},
	{
		Name:         "",
		Category:     "",
		Subcategory:  "",
		Requirements: []string{},
		InstallCmd:   ``,
	},
	{
		Name:         "",
		Category:     "",
		Subcategory:  "",
		Requirements: []string{},
		InstallCmd:   ``,
	},
	{
		Name:         "",
		Category:     "",
		Subcategory:  "",
		Requirements: []string{},
		InstallCmd:   ``,
	},
	{
		Name:         "",
		Category:     "",
		Subcategory:  "",
		Requirements: []string{},
		InstallCmd:   ``,
	},
	{
		Name:         "",
		Category:     "",
		Subcategory:  "",
		Requirements: []string{},
		InstallCmd:   ``,
	},
}
