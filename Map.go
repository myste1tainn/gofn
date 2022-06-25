package gofn

import "fmt"

func Map[Item any, Result any](items []Item, transform func(item Item, index int) Result) []Result {
	var results []Result
	for index, item := range items {
		results = append(results, transform(item, index))
	}
	return results
}

func FlatMap[Item any, Result any](items []Item, transform func(item Item, index int) *Result) []Result {
	var results []Result
	for index, item := range items {
		result := transform(item, index)
		if result != nil {
			results = append(results, *result)
		}
	}
	return results
}

func MapOfMap[Key comparable, Item any, Result any](items map[Key]Item, transform func(item Item, key Key) Result) map[Key]Result {
	var result map[Key]Result
	for key, item := range items {
		result[key] = transform(item, key)

	}
	return result
}

func FlatMapMapOfMap[Key comparable, Item any, Result any](items map[Key]Item, transform func(item Item, key Key) *Result) map[Key]Result {
	var result map[Key]Result
	for key, item := range items {
		res := transform(item, key)
		result[key] = *res

	}
	return result
}

func ToString[Item any](item Item) string {
	return fmt.Sprintf("%v", item)
}
