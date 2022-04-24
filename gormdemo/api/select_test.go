package api

import (
	"fmt"
	"testing"
)

func TestFirst(t *testing.T) {
	first, _ := First()
	fmt.Printf("%+v", first)
}
func TestLast(t *testing.T) {
	last := Last()
	fmt.Printf("%+v", last)
}

func TestTake(t *testing.T) {
	take := Take()
	fmt.Printf("%+v", take)
}
