package ucloud

type options struct {
	projectId string
}

type ClientOption interface {
	apply(*options)
}

type funcOption struct {
	f func(*options)
}

func (fdo *funcOption) apply(do *options) {
	fdo.f(do)
}

func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func WithProjectId(projectId string) ClientOption  {
	return newFuncOption(func(o *options) {
		o.projectId = projectId
	})
}