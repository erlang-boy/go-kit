package binder

type BinderSpec struct {
	BindServerName string
	BindAddr       string
}

type Binder interface {
	Start() error
	Stop() error
	KeepAlive() error
	Status() (string, error)
}

func NewBinder(spec *BinderSpec) Binder {
	return binderFac(spec)
}
