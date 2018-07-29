package yamaha_rt_simulator

type DebugLogger interface {
	Debug(args ...interface{})
}

type InfoLogger interface {
	Info(args ...interface{})
}

type ErrorLogger interface {
	Error(args ...interface{})
}
