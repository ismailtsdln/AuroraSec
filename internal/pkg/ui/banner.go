package ui

import (
	"fmt"

	"github.com/fatih/color"
)

const Banner = `
   ____                              _____           
  / __ \                            / ____|          
 | |  | |_   _ _ __ ___  _ __ __ _ | (___   ___  ___ 
 | |  | | | | | '__/ _ \| '__/ _' | \___ \ / _ \/ __|
 | |__| | |_| | | | (_) | | | (_| | ____) |  __/ (__ 
  \____/ \__,_|_|  \___/|_|  \__,_||_____/ \___|\___|
                                                     
      Next-Gen AWS Auditing & Hardening Tool
`

func PrintBanner() {
	cyan := color.New(color.FgCyan, color.Bold)
	magenta := color.New(color.FgMagenta)

	cyan.Print(Banner)
	magenta.Println("  ------------------------------------------------")
	fmt.Println()
}
