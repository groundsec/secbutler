package utils

import "fmt"

func Banner(version string) {
	banner := `
   ________  _____/ /_  __  __/ /_/ /__  _____
  / ___/ _ \/ ___/ __ \/ / / / __/ / _ \/ ___/
 (__  )  __/ /__/ /_/ / /_/ / /_/ /  __/ /    
/____/\___/\___/_.___/\__,_/\__/_/\___/_/																		 

v%s - https://github.com/groundsec/secbutler

`
	fmt.Printf(banner, version)
}
