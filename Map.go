package gofn

import "fmt"

func Map[A any, B any](items []A, transform func(item A) B) []B {
	var b []B
	for _, item := range items {
		b = append(b, transform(item))
	}
	return b
}

func FlatMap[A any, B any](items []A, transform func(item A) *B) []B {
	var bs []B
	for _, item := range items {
		b := transform(item)
		if b != nil {
			bs = append(bs, *b)
		}
	}
	return bs
}

func ToString[Item any](item Item) string {
    return fmt.Sprintf("%s", item)
}
