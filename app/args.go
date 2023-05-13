package app

import (
	"os"

	"nup/static"
	"nup/types"
)

func ParseArgs(args []string) types.Args {
	parsedArgs := types.Args{
		Latency: false,
		Status:  false,
		Verbose: false,
		LogAll:  false,
		Domains: []string{"https://www.google.com"},
	}

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-s":
			parsedArgs.Status = true // default is false, prints status code, should be used with -v
		case "-t":
			parsedArgs.Latency = true // default is false, prints latency, should be used with -v
		case "-v":
			parsedArgs.Verbose = true // default is false prints realtime output
		case "-a", "--all":
			parsedArgs.LogAll = true // deafult is false, only log failed requests
		case "-h", "--help":
			static.HelpMsg()
			os.Exit(0)
		}
	}

	return parsedArgs
}
