package gofn

func ForEach[Item any](items []Item, run func(item Item, index int)) {
	for index, item := range items {
		run(item, index)
	}
}

func ForEachOfMap[Key comparable, Value any](items map[Key]Value, run func(item Value, key Key)) {
	for key, item := range items {
		run(item, key)
	}
}
