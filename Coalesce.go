package gofn

func Coalesce[Obj any](objs ...*Obj) *Obj {
	for _, obj := range objs {
		if obj != nil {
			return obj
		}
	}
	return nil
}

func Get0[R any](fn func() (R, error)) R {
	r, err := fn()
	if err != nil {
		panic(err)
	} else {
		return r
	}
}

func Get1[E1 any, R any](fn func(e1 E1) (R, error), e1 E1) R {
	r, err := fn(e1)
	if err != nil {
		panic(err)
	} else {
		return r
	}
}
