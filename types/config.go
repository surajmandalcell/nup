package types

type Config struct {
	Domains      []string
	IntervalSecs uint64
	TimeoutSecs  uint64
	FlagLatency  bool
	FlagStatus   bool
	FlagVerbose  bool
}
