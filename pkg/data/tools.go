package data

import "github.com/groundsec/secbutler/pkg/types"

var SecTools = []types.SecTool{
	{
		Name:         "amass",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/owasp-amass/amass/v4/...@master`,
	},
	{
		Name:         "chaos",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/chaos-client/cmd/chaos@latest`,
	},
	{
		Name:         "shuffledns",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/shuffledns/cmd/shuffledns@latest`,
	},
	{
		Name:         "as3nt",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"python3", "pip"},
		InstallCmd:   `pip install as3nt`,
	},
	{
		Name:         "altdns",
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
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"git", "python3"},
		InstallCmd: `
		git clone https://github.com/guelfoweb/knock.git
		cd knock
		python3 setup.py install
		cd ..
		`,
	},
	{
		Name:         "subfinder",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest`,
	},
	{
		Name:         "assetfinder",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/tomnomnom/assetfinder@latest`,
	},
	{
		Name:         "vhostscan",
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
		Name:         "scilla",
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
		Name:         "rustscan",
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
		Name:         "naabu",
		Category:     "Recon",
		Subcategory:  "Port Scanning",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/naabu/v2/cmd/naabu@latest`,
	},
	{
		Name:         "nmap",
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
		Name:         "sandmap",
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
	/*
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
	*/
}
