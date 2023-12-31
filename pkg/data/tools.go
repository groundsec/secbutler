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
		Name:         "vhostscan",
		Category:     "Recon",
		Subcategory:  "Subdomain Enumeration",
		Requirements: []string{"git", "python3"},
		InstallCmd: `
		git clone https://github.com/codingo/VHostScan.git &> /dev/null
		cd VHostScan
		pip install numpy==1.12.0 &> /dev/null
		python3 setup.py install &> /dev/null
		cd ..
		`,
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
	*/
}
