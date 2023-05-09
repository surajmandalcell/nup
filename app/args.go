package app

import (
	f "fmt"
	"os"

	"nup/services"
	"nup/statics"
	"nup/types"
)

func ParseArgs(args []string) types.Args {
	// Experimental
	db := services.DbService()

	config, _ := db.Get("SELECT * FROM config")
	defer config.Close()

	for config.Next() {
		var (
			id         int
			key, value string
		)
		config.Scan(&id, &key, &value)
		f.Printf("Config: id=%d, key=%s, value=%s\n", id, key, value)
	}
	// Experimental End

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
