package gofn

type Enumerated[Item any] struct {
	Index int
	Item  Item
}

func EnumerateArray[Item any](items []Item) []Enumerated[Item] {
	var enumeratedItems []Enumerated[Item]
	for i, item := range items {
		enumeratedItems = append(enumeratedItems, Enumerate(i, item))
	}
	return enumeratedItems
}

func Enumerate[Item any](i int, item Item) Enumerated[Item] {
	return Enumerated[Item]{
		Index: i,
		Item:  item,
	}
}
