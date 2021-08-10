package etcaid

type logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}
