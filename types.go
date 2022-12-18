package main

import "fmt"

type Card interface {
	fmt.Stringer

	Name() string
}
