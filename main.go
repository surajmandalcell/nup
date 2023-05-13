package main

import (
	"nup/app"
	"nup/services"
	"nup/types"

	"os"
	"sync"
)

func main() {
	args := app.ParseArgs(os.Args[1:]) // Initialize args
	wg := new(sync.WaitGroup)          // Initialize wait group
	db := services.DbInit()            // Initialize database
	db.SetConfig(args)

	pingInstance := app.Init(wg, db, types.Config{ // Initialize ping instance
		Domains:      args.Domains,
		IntervalSecs: 1,
		TimeoutSecs:  1,
		FlagStatus:   args.Status,
		FlagLatency:  args.Latency,
		FlagVerbose:  args.Verbose,
	})

	wg.Add(1)
	go pingInstance.Ping()

	if args.Verbose != true {
		app.MainPrompt(pingInstance)
	}
	wg.Wait()
}
