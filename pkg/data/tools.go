package data

import "github.com/groundsec/secbutler/pkg/types"

var SecTools = []types.SecTool{
	{
		Name:         "amass",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/owasp-amass/amass/v4/...@master &> /dev/null`,
	},
	{
		Name:         "chaos",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/chaos-client/cmd/chaos@latest &> /dev/null`,
	},
	{
		Name:         "shuffledns",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/shuffledns/cmd/shuffledns@latest &> /dev/null`,
	},
	{
		Name:         "as3nt",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"python3", "pip"},
		InstallCmd:   `pip install as3nt &> /dev/null`,
	},
	{
		Name:         "altdns",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"python3", "pip3"},
		InstallCmd:   `pip3 install py-altdns==1.0.2 &> /dev/null`,
	},
	{
		Name:         "dnsx",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/dnsx/cmd/dnsx@latest &> /dev/null`,
	},
	{
		Name:         "hakrevdns",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/hakluke/hakrevdns@latest &> /dev/null`,
	},
	{
		Name:         "knockpy",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"git", "python3"},
		InstallCmd: `
		git clone https://github.com/guelfoweb/knock.git &> /dev/null
		cd knock
		python3 setup.py install &> /dev/null
		cd ..
		`,
	},
	{
		Name:         "subfinder",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest &> /dev/null`,
	},
	{
		Name:         "assetfinder",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/tomnomnom/assetfinder@latest &> /dev/null`,
	},
	{
		Name:         "scilla",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"git", "go"},
		InstallCmd: `
		git clone https://github.com/edoardottt/scilla.git &> /dev/null
		cd scilla
		make linux &> /dev/null
		cd ..
		`,
	},
	{
		Name:         "cero",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/glebarez/cero@latest &> /dev/null`,
	},
	{
		Name:         "masscan",
		Category:     "Recon",
		Subcategory:  "Port Scanning",
		Requirements: []string{"make", "git", "gcc"},
		InstallCmd: `
		git clone https://github.com/robertdavidgraham/masscan &> /dev/null
		cd masscan
		make &> /dev/null
		make install &> /dev/null
		cd ..
		`,
	},
	{
		Name:         "rustscan",
		Category:     "Recon",
		Subcategory:  "Port Scanning",
		Requirements: []string{"cargo"},
		InstallCmd: `
		git clone https://github.com/RustScan/RustScan.git &> /dev/null
		cd RustScan
		cargo build --release &> /dev/null
		mv target/release/rustscan /usr/local/bin/
		cd ..
		`,
	},
	{
		Name:         "naabu",
		Category:     "Recon",
		Subcategory:  "Port Scanning",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/naabu/v2/cmd/naabu@latest &> /dev/null`,
	},
	{
		Name:         "nmap",
		Category:     "Recon",
		Subcategory:  "Port Scanning",
		Requirements: []string{"git", "make", "gcc"},
		InstallCmd: `
		git clone https://github.com/nmap/nmap.git &> /dev/null
		cd nmap
		./configure &> /dev/null
		make &> /dev/null
		make install &> /dev/null
		cd ..
		`,
	},
	{
		Name:         "sandmap",
		Category:     "Recon",
		Subcategory:  "Port Scanning",
		Requirements: []string{},
		InstallCmd: `
		git clone --recursive https://github.com/trimstray/sandmap &> /dev/null
		cd sandmap
		./setup.sh install &> /dev/null
		`,
	},
	{
		Name:         "screenshoteer",
		Category:     "Recon",
		Subcategory:  "Screenshots",
		Requirements: []string{"npm"},
		InstallCmd:   `npm i -g screenshoteer`,
	},
	{
		Name:         "gowitness",
		Category:     "Recon",
		Subcategory:  "Screenshots",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/sensepost/gowitness@latest`,
	},
	{
		Name:         "witnessme",
		Category:     "Recon",
		Subcategory:  "Screenshots",
		Requirements: []string{"pip"},
		InstallCmd:   `pip install witnessme`,
	},
	{
		Name:         "webanalyze",
		Category:     "Recon",
		Subcategory:  "Technologies",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/rverton/webanalyze/cmd/webanalyze@latest`,
	},
	{
		Name:         "whatweb",
		Category:     "Recon",
		Subcategory:  "Technologies",
		Requirements: []string{"git", "make"},
		InstallCmd: `
		git clone https://github.com/urbanadventurer/WhatWeb.git
		cd WhatWeb
		make install
		cd ..`,
	},
	{
		Name:         "retire",
		Category:     "Recon",
		Subcategory:  "Technologies",
		Requirements: []string{"npm"},
		InstallCmd:   `npm i -g retire`,
	},
	{
		Name:         "httpx",
		Category:     "Recon",
		Subcategory:  "Technologies",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/httpx/cmd/httpx@latest`,
	},
	{
		Name:         "fingerprintx",
		Category:     "Recon",
		Subcategory:  "Technologies",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/praetorian-inc/fingerprintx/cmd/fingerprintx@latest`,
	},
	{
		Name:         "gobuster",
		Category:     "Recon",
		Subcategory:  "Content Discovery",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/OJ/gobuster/v3@latest`,
	},
	{
		Name:         "feroxbuster",
		Category:     "Recon",
		Subcategory:  "Content Discovery",
		Requirements: []string{"cargo"},
		InstallCmd:   `cargo install feroxbuster`,
	},
	{
		Name:         "dirsearch",
		Category:     "Recon",
		Subcategory:  "Content Discovery",
		Requirements: []string{"pip"},
		InstallCmd:   `pip install dirsearch`,
	},
	{
		Name:         "gospider",
		Category:     "Recon",
		Subcategory:  "Content Discovery",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/jaeles-project/gospider@latest`,
	},
	{
		Name:         "hakrawler",
		Category:     "Recon",
		Subcategory:  "Content Discovery",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/hakluke/hakrawler@latest`,
	},
	/*
		{
			Name:         "",
			Category:     "",
			Subcategory:  "",
			Requirements: []string{},
			InstallCmd:   ``,
		},
	*/
}
