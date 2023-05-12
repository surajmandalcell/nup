package main

import (
	"nup/app"
	"nup/services"
	"nup/types"

	"os"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	args := app.ParseArgs(os.Args[1:])
	db := services.DbService()
	db.SetConfig(args)

	app.GPingService = app.Init(wg, types.Config{
		Domains:      args.Domains,
		IntervalSecs: 1,
		TimeoutSecs:  5,
		FlagStatus:   args.Status,
		FlagLatency:  args.Latency,
		FlagVerbose:  args.Verbose,
	})

	wg.Add(1)

	go app.GPingService.Ping()

	app.MainPrompt()
	wg.Wait()
}
