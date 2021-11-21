package algo_structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleArray(t *testing.T) {
	a := NewSingleArray()

	ast := assert.New(t)
	ast.Equal(0, a.Count())

	a.Add(10)
	ast.Equal(1, a.Count())
	ast.Equal(10, a.Get(0))

	a.Add(12)

	ast.Equal(2, a.Count())
	ast.Equal(12, a.Get(1))

	a.RemoveLast()

	ast.Equal(3, a.Count())
	ast.Panics(func() {
		a.Get(2)
	})
}