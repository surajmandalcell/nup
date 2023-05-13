package types

type Config struct {
	Domains      []string
	FlagLatency  bool
	FlagStatus   bool
	FlagVerbose  bool
	TimeoutSecs  uint64
	IntervalSecs int8
}
