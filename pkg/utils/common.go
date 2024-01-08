package utils

var StartInstallTpl = `
#!/usr/bin/env bash

function banner {
  echo "                   __          __  __"
  echo "   ________  _____/ /_  __  __/ /_/ /__  _____"
  echo "  / ___/ _ \/ ___/ __ \/ / / / __/ / _ \/ ___/"
  echo " (__  )  __/ /__/ /_/ / /_/ / /_/ /  __/ /"    
  echo "/____/\___/\___/_.___/\__,_/\__/_/\___/_/"
  echo ""
}

function bold {
  echo -e "\033[1m$1\033[0m"
}

function info {
  echo -e "\e[33m[INFO] $1\e[0m"
}

function success {
  echo -e "\e[32m[SUCC] $1\e[0m"
}

function error {
  echo -e "\e[31m[ERRO] $1\e[0m"
}

`

var StartCheckRequiremenstFunc = `function check_requirements {
	bold "Checking requirements"	
`

var EndCheckRequirementsFunc = `
  for requirement in ${REQUIREMENTS[@]}; do
		info "Checking if requirement '$requirement' is installed"
    if ! command -v $requirement &> /dev/null; then
      error "Requirement '$requirement' is not installed, impossible to continue installation"
      exit 1
    else
      success "Requirement '$requirement' is correctly installed"
    fi
  done
	echo ""
}`
