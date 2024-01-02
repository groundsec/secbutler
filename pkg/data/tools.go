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
		Requirements: []string{"python3", "pipx"},
		InstallCmd:   `pipx install as3nt &> /dev/null`,
	},
	{
		Name:         "altdns",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"python3", "pipx"},
		InstallCmd:   `pipx install py-altdns==1.0.2 &> /dev/null`,
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
		InstallCmd:   `npm i -g screenshoteer &> /dev/null`,
	},
	{
		Name:         "gowitness",
		Category:     "Recon",
		Subcategory:  "Screenshots",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/sensepost/gowitness@latest &> /dev/null`,
	},
	{
		Name:         "webanalyze",
		Category:     "Recon",
		Subcategory:  "Technologies",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/rverton/webanalyze/cmd/webanalyze@latest &> /dev/null`,
	},
	{
		Name:         "whatweb",
		Category:     "Recon",
		Subcategory:  "Technologies",
		Requirements: []string{"git", "make"},
		InstallCmd: `
		git clone https://github.com/urbanadventurer/WhatWeb.git &> /dev/null
		cd WhatWeb
		make install &> /dev/null
		cd ..`,
	},
	{
		Name:         "retire",
		Category:     "Recon",
		Subcategory:  "Technologies",
		Requirements: []string{"npm"},
		InstallCmd:   `npm i -g retire &> /dev/null`,
	},
	{
		Name:         "httpx",
		Category:     "Recon",
		Subcategory:  "Technologies",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/projectdiscovery/httpx/cmd/httpx@latest &> /dev/null`,
	},
	{
		Name:         "fingerprintx",
		Category:     "Recon",
		Subcategory:  "Technologies",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/praetorian-inc/fingerprintx/cmd/fingerprintx@latest &> /dev/null`,
	},
	{
		Name:         "gobuster",
		Category:     "Recon",
		Subcategory:  "Content Discovery",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/OJ/gobuster/v3@latest &> /dev/null`,
	},
	{
		Name:         "feroxbuster",
		Category:     "Recon",
		Subcategory:  "Content Discovery",
		Requirements: []string{"cargo"},
		InstallCmd:   `cargo install feroxbuster &> /dev/null`,
	},
	{
		Name:         "dirsearch",
		Category:     "Recon",
		Subcategory:  "Content Discovery",
		Requirements: []string{"pipx"},
		InstallCmd:   `pipx install dirsearch &> /dev/null`,
	},
	{
		Name:         "gospider",
		Category:     "Recon",
		Subcategory:  "Content Discovery",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/jaeles-project/gospider@latest &> /dev/null`,
	},
	{
		Name:         "hakrawler",
		Category:     "Recon",
		Subcategory:  "Content Discovery",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/hakluke/hakrawler@latest &> /dev/null`,
	},
	{
		Name:         "GoLinkFinder",
		Category:     "Recon",
		Subcategory:  "Links",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/0xsha/GoLinkFinder@latest &> /dev/null`,
	},
	{
		Name:         "waybackurls",
		Category:     "Recon",
		Subcategory:  "Links",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/tomnomnom/waybackurls@latest &> /dev/null`,
	},
	{
		Name:         "gau",
		Category:     "Recon",
		Subcategory:  "Links",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/lc/gau/v2/cmd/gau@latest &> /dev/null`,
	},
	{
		Name:         "getJS",
		Category:     "Recon",
		Subcategory:  "Links",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/003random/getJS@latest &> /dev/null`,
	},
	{
		Name:         "linx",
		Category:     "Recon",
		Subcategory:  "Links",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/riza/linx/cmd/linx@latest &> /dev/null`,
	},
	{
		Name:         "arjun",
		Category:     "Recon",
		Subcategory:  "Parameters",
		Requirements: []string{"pipx"},
		InstallCmd:   `pipx install arjun &> /dev/null`,
	},
	{
		Name:         "paramspider",
		Category:     "Recon",
		Subcategory:  "Parameters",
		Requirements: []string{"git", "pipx"},
		InstallCmd: `
		git clone https://github.com/devanshbatham/paramspider &> /dev/null
		cd paramspider
		pipx install . &> /dev/null
		cd ..
		`,
	},
	{
		Name:         "x8",
		Category:     "Recon",
		Subcategory:  "Parameters",
		Requirements: []string{"cargo"},
		InstallCmd:   `cargo install x8 &> /dev/null`,
	},
	{
		Name:         "ffuf",
		Category:     "Recon",
		Subcategory:  "Fuzzing",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/ffuf/ffuf/v2@latest &> /dev/null`,
	},
	{
		Name:         "qsfuzz",
		Category:     "Recon",
		Subcategory:  "Fuzzing",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/ameenmaali/qsfuzz@latest &> /dev/null`,
	},
	{
		Name:         "vaf",
		Category:     "Recon",
		Subcategory:  "Fuzzing",
		Requirements: []string{"curl", "sudo", "bash"},
		InstallCmd:   `curl https://raw.githubusercontent.com/d4rckh/vaf/main/install.sh &> /dev/null | sudo bash &> /dev/null`,
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
