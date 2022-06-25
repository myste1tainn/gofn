package gofn

func ForEach[A any](items []A, run func(item A)) {
	for _, item := range items {
		run(item)
	}
}
