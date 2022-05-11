package fn

func Throw(errs ...error) {
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
}
