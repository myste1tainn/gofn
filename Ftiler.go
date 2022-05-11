package fn

func Filter[Item any](items []Item, predicate func(item Item) bool) *Item {
	for _, item := range items {
		if predicate(item) {
			return &item
		}
	}
	return nil
}
