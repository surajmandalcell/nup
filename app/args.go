package app

import (
	"os"

	"nup/statics"
	"nup/types"
)

func ParseArgs(args []string) types.Args {
	parsedArgs := types.Args{
		Latency: false,
		Status:  false,
		Verbose: false,
		Domains: []string{"https://www.google.com"},
	}

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-s":
			parsedArgs.Status = true
		case "-t":
			parsedArgs.Latency = true
		case "-v":
			parsedArgs.Verbose = true
		case "-h", "--help":
			statics.HelpMsg()
			os.Exit(0)
		}
	}

	return parsedArgs
}
