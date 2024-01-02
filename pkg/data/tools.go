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
		InstallCmd:   `curl -s https://raw.githubusercontent.com/d4rckh/vaf/main/install.sh | sudo bash &> /dev/null`,
	},
	{
		Name:         "commix",
		Category:     "Exploitation",
		Subcategory:  "Command Injection",
		Requirements: []string{"git", "python3"},
		InstallCmd: `
		git clone https://github.com/commixproject/commix.git &> /dev/null
		cd commix
		python3 setup.py install &> /dev/null
		cd ..`,
	},
	{
		Name:         "CorsMe",
		Category:     "Exploitation",
		Subcategory:  "CORS Misconfiguration",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/shivangx01b/CorsMe@latest &> /dev/null`,
	},
	{
		Name:         "crlfsuite",
		Category:     "Exploitation",
		Subcategory:  "CRLF Injection",
		Requirements: []string{"pipx"},
		InstallCmd:   `pipx install crlfsuite &> /dev/null`,
	},
	{
		Name:         "crlfuzz",
		Category:     "Exploitation",
		Subcategory:  "CRLF Injection",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/dwisiswant0/crlfuzz/cmd/crlfuzz@latest &> /dev/null`,
	},
	{
		Name:         "crlf",
		Category:     "Exploitation",
		Subcategory:  "CRLF Injection",
		Requirements: []string{"git", "python3"},
		InstallCmd: `
		git clone https://github.com/MichaelStott/CRLF-Injection-Scanner.git crlf  &> /dev/null
		cd crlf
		python3 setup.py install &> /dev/null
		cd ..`,
	},
	{
		Name:         "xsrfprobe",
		Category:     "Exploitation",
		Subcategory:  "CSRF Injection",
		Requirements: []string{"pipx"},
		InstallCmd:   `pipx install xsrfprobe &> /dev/null`,
	},
	{
		Name:         "graphqlmap",
		Category:     "Exploitation",
		Subcategory:  "GraphQL Injection",
		Requirements: []string{"git", "python3"},
		InstallCmd: `
		git clone https://github.com/swisskyrepo/GraphQLmap &> /dev/null
		cd GraphQLmap
		python3 setup.py install &> /dev/null
		cd ..`,
	},
	{
		Name:         "clairvoyance",
		Category:     "Exploitation",
		Subcategory:  "GraphQL Injection",
		Requirements: []string{"pipx"},
		InstallCmd:   `pipx install clairvoyance &> /dev/null`,
	},
	{
		Name:         "shifter",
		Category:     "Exploitation",
		Subcategory:  "GraphQL Injection",
		Requirements: []string{"git", "python3"},
		InstallCmd: `
		git clone https://github.com/szski/shapeshifter.git &> /dev/null
		cd shapeshifter/shapeshifter
		python3 setup.py install &> /dev/null
		cd ../..
		`,
	},
	{
		Name:         "headi",
		Category:     "Exploitation",
		Subcategory:  "Header Injection",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/mlcsec/headi@latest &> /dev/null`,
	},
	{
		Name:         "openredirex",
		Category:     "Exploitation",
		Subcategory:  "Open Redirect",
		Requirements: []string{},
		InstallCmd: `
		git clone https://github.com/devanshbatham/openredirex &> /dev/null
		cd openredirex
		sudo chmod +x setup.sh
		./setup.sh &> /dev/null
		cd ..
`,
	},
	{
		Name:         "lorsrf",
		Category:     "Exploitation",
		Subcategory:  "SSRF",
		Requirements: []string{"cargo"},
		InstallCmd:   `cargo install --git https://github.com/knassar702/lorsrf &> /dev/null`,
	},
	{
		Name:         "whonow",
		Category:     "Exploitation",
		Subcategory:  "SSRF",
		Requirements: []string{"npm"},
		InstallCmd:   `npm install --cli -g whonow@latest &> /dev/null`,
	},
	{
		Name:         "sqlmap",
		Category:     "Exploitation",
		Subcategory:  "SQL Injection",
		Requirements: []string{"pipx"},
		InstallCmd:   `pipx install sqlmap &> /dev/null`,
	},
	{
		Name:         "nosqli",
		Category:     "Exploitation",
		Subcategory:  "SQL Injection",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/Charlie-belmer/nosqli@latest &> /dev/null`,
	},
	{
		Name:         "dalfox",
		Category:     "Exploitation",
		Subcategory:  "XSS",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/hahwul/dalfox/v2@latest &> /dev/null`,
	},
	{
		Name:         "XSpear",
		Category:     "Exploitation",
		Subcategory:  "XSS",
		Requirements: []string{"gem"},
		InstallCmd:   `gem install XSpear &> /dev/null`,
	},
	{
		Name:         "xsser",
		Category:     "Exploitation",
		Subcategory:  "XSS",
		Requirements: []string{"git", "python3"},
		InstallCmd: `
		git clone https://github.com/epsylon/xsser &> /dev/null
		cd xsser
		python3 setup.py install &> /dev/null
		cd ..`,
	},
	{
		Name:         "xxexploiter",
		Category:     "Exploitation",
		Subcategory:  "XXE Injection",
		Requirements: []string{"npm"},
		InstallCmd:   `npm install -g xxexploiter &> /dev/null`,
	},
	{
		Name:         "creds",
		Category:     "Misc",
		Subcategory:  "Bruteforce",
		Requirements: []string{},
		InstallCmd:   `pipx install defaultcreds-cheat-sheet &> /dev/null`,
	},
	{
		Name:         "brutex",
		Category:     "Misc",
		Subcategory:  "Bruteforce",
		Requirements: []string{"git"},
		InstallCmd: `
		git clone https://github.com/1N3/BruteX.git &> /dev/null
		cd BruteX
		chmod +x install.sh
		./install.sh &> /dev/null
		cd ..`,
	},
	{
		Name:         "hydra",
		Category:     "Misc",
		Subcategory:  "Bruteforce",
		Requirements: []string{"git", "make"},
		InstallCmd: `
		git clone https://github.com/vanhauser-thc/thc-hydra.git hydra &> /dev/null
		cd hydra
		./configure &> /dev/null
		make &> /dev/null
		make install &> /dev/null
		cd ..`,
	},
	{
		Name:         "git secrets",
		Category:     "Misc",
		Subcategory:  "Secrets",
		Requirements: []string{"git", "make"},
		InstallCmd: `
		git clone https://github.com/awslabs/git-secrets.git &> /dev/null
		cd git-secrets
		make install &> /dev/null
		cd ..`,
	},
	{
		Name:         "gitleaks",
		Category:     "Misc",
		Subcategory:  "Secrets",
		Requirements: []string{"git", "make"},
		InstallCmd: `
		git clone https://github.com/gitleaks/gitleaks.git &> /dev/null
		cd gitleaks
		make build &> /dev/null
		go install . &> /dev/null
		cd ..`,
	},
	{
		Name:         "trufflehog",
		Category:     "Misc",
		Subcategory:  "Secrets",
		Requirements: []string{"curl", "sh"},
		InstallCmd:   `curl -sSfL https://raw.githubusercontent.com/trufflesecurity/trufflehog/main/scripts/install.sh | sh -s -- -b /usr/local/bin &> /dev/null`,
	},
	{
		Name:         "commit-stream",
		Category:     "Misc",
		Subcategory:  "Secrets",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/x1sec/commit-stream@latest &> /dev/null`,
	},
	{
		Name:         "shhgit",
		Category:     "Misc",
		Subcategory:  "Secrets",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/eth0izzle/shhgit@latest &> /dev/null`,
	},
	{
		Name:         "detect-secrets",
		Category:     "Misc",
		Subcategory:  "Secrets",
		Requirements: []string{"pipx"},
		InstallCmd:   `pipx install detect-secrets &> /dev/null`,
	},
	{
		Name:         "whispers",
		Category:     "Misc",
		Subcategory:  "Secrets",
		Requirements: []string{"pipx"},
		InstallCmd:   `pipx install whispers &> /dev/null`,
	},
	{
		Name:         "yar",
		Category:     "Misc",
		Subcategory:  "Secrets",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/nielsing/yar@latest &> /dev/null`,
	},
	{
		Name:         "go-earlybird",
		Category:     "Misc",
		Subcategory:  "Secrets",
		Requirements: []string{"git"},
		InstallCmd: `
		git clone https://github.com/americanexpress/earlybird.git &> /dev/null
		cd earlybird
		./build.sh &> /dev/null
		./install.sh &> /dev/null
		cd ..`,
	},
	{
		Name:         "ankamali_hog",
		Category:     "Misc",
		Subcategory:  "Secrets",
		Requirements: []string{"cargo"},
		InstallCmd: `
			git clone https://github.com/newrelic/rusty-hog.git &> /dev/null
			cd rusty-hog
			cargo build --release &> /dev/null
			mv target/release/*_hog /usr/local/bin
			mv target/release/*_hog_lambda /usr/local/bin
			cd ..`,
	},
	{
		Name:         "noseyparker",
		Category:     "Misc",
		Subcategory:  "Secrets",
		Requirements: []string{"cargo", "cmake", "git"},
		InstallCmd: `
			git clone https://github.com/praetorian-inc/noseyparker.git &> /dev/null
			cd noseyparker
			cargo build --locked --profile release
			mv target/release/noseyparker-cli /usr/local/bin/noseyparker
			cd ..`,
	},
	{
		Name:         "gitjacker",
		Category:     "Misc",
		Subcategory:  "Git",
		Requirements: []string{"git", "bash"},
		InstallCmd:   `curl -s "https://raw.githubusercontent.com/liamg/gitjacker/master/scripts/install.sh" | bash &> /dev/null`,
	},
	{
		Name:         "git-dumper",
		Category:     "Misc",
		Subcategory:  "Git",
		Requirements: []string{"pipx"},
		InstallCmd:   `pipx install git-dumper &> /dev/null`,
	},
	{
		Name:         "GitHunter",
		Category:     "Misc",
		Subcategory:  "Git",
		Requirements: []string{"git", "go"},
		InstallCmd: `
		git clone https://github.com/digininja/GitHunter &> /dev/null
		cd GitHunter
		go build &> /dev/null
		mv GitHunter /usr/local/bin
		cd ..`,
	},
	{
		Name:         "gato",
		Category:     "Misc",
		Subcategory:  "Git",
		Requirements: []string{"git", "pipx"},
		InstallCmd: `
		git clone https://github.com/praetorian-inc/gato &> /dev/null
		cd gato
		pipx install . &> /dev/null
		cd ..`,
	},
	{
		Name:         "s3scanner",
		Category:     "Misc",
		Subcategory:  "Buckets",
		Requirements: []string{"go"},
		InstallCmd:   `go install -v github.com/sa7mon/s3scanner@latest &> /dev/null`,
	},
	{
		Name:         "festin",
		Category:     "Misc",
		Subcategory:  "Buckets",
		Requirements: []string{"pipx"},
		InstallCmd:   `pipx install festin &> /dev/null`,
	},
	{
		Name:         "s3reverse",
		Category:     "Misc",
		Subcategory:  "Buckets",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/hahwul/s3reverse@latest &> /dev/null`,
	},
	{
		Name:         "dirlstr",
		Category:     "Misc",
		Subcategory:  "Buckets",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/cybercdh/dirlstr@latest &> /dev/null`,
	},
	{
		Name:         "kicks3.py",
		Category:     "Misc",
		Subcategory:  "Buckets",
		Requirements: []string{"pipx"},
		InstallCmd:   `pipx install kick-s3 &> /dev/null`,
	},
	{
		Name:         "2tearsinabucket",
		Category:     "Misc",
		Subcategory:  "Buckets",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/Revenant40/2tearsinabucket@latest &> /dev/null`,
	},
	{
		Name:         "s3tk",
		Category:     "Misc",
		Subcategory:  "Buckets",
		Requirements: []string{"pipx"},
		InstallCmd:   `pipx install s3tk &> /dev/null`,
	},
	{
		Name:         "cloudbrute",
		Category:     "Misc",
		Subcategory:  "Buckets",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/0xsha/cloudbrute@latest &> /dev/null`,
	},
	{
		Name:         "wpscan",
		Category:     "Misc",
		Subcategory:  "CMS",
		Requirements: []string{"gem"},
		InstallCmd:   `gem install wpscan &> /dev/null`,
	},
	{
		Name:         "wprecon",
		Category:     "Misc",
		Subcategory:  "CMS",
		Requirements: []string{"go"},
		InstallCmd: `
		git clone github.com/blackcrw/wpreconx.git &> /dev/null
		cd wpreconx/cli
		go build &> /dev/null
		mv cli /usr/local/bin/wprecon
		cd ../..
		`,
	},
	{
		Name:         "jwtear",
		Category:     "Misc",
		Subcategory:  "JWT",
		Requirements: []string{"gem"},
		InstallCmd:   `gem install jwtear &> /dev/null`,
	},
	{
		Name:         "jwt-hack",
		Category:     "Misc",
		Subcategory:  "JWT",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/hahwul/jwt-hack@latest &> /dev/null`,
	},
	{
		Name:         "jwt-cracker",
		Category:     "Misc",
		Subcategory:  "JWT",
		Requirements: []string{"npm"},
		InstallCmd:   `npm install --global jwt-cracker &> /dev/null`,
	},
	{
		Name:         "csprecon",
		Category:     "Misc",
		Subcategory:  "Uncategorized",
		Requirements: []string{"go"},
		InstallCmd:   `go install github.com/edoardottt/csprecon/cmd/csprecon@latest &> /dev/null`,
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
