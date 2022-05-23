package gofn

func Reduce[Item any, B any](items []Item, seed B, reduction func(partialResult B, next Item, index int) B) B {
	result := seed
	for i, item := range items {
		result = reduction(result, item, i)
	}
	return result
}

func ReduceInto[Item any, B any](items []Item, seed B, reduction func(partialResult *B, next Item, index int)) B {
	result := &seed
	for i, item := range items {
		reduction(result, item, i)
	}
	return *result
}
