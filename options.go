package xdb

const _defaultSingularTable = true

type options struct {
	// SingularTable use singular table by default
	singularTable bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func SingularTable(singularTable bool) Option {
	return optionFunc(func(o *options) {
		o.singularTable = singularTable
	})
}
