package main

import (
	"nup/app"
	"nup/types"

	"os"
)

func main() {
	args := app.ParseArgs(os.Args[1:])

	app.Ping(types.Config{
		Domains:      args.Domains,
		IntervalSecs: 1,
		TimeoutSecs:  5,
		FlagStatus:   args.Status,
		FlagLatency:  args.Latency,
		FlagVerbose:  args.Verbose,
	})
}
