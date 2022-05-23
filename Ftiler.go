package gofn

func Filter[Item any](items []Item, predicate func(item Item) bool) []Item {
	var newItems []Item
	for _, item := range items {
		if predicate(item) {
			newItems = append(newItems, item)
		}
	}
	return newItems
}

func First[Item any](items []Item, predicate func(item Item) bool) *Item {
	for _, item := range items {
		if predicate(item) {
			return &item
		}
	}
	return nil
}

func FirstIndex[Item any](items []Item, predicate func(item Item) bool) *Item {
	for _, item := range items {
		if predicate(item) {
			return &item
		}
	}
	return nil
}
