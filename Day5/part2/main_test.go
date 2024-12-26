package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve5(t *testing.T) {
	t.Run("", func(t *testing.T) {
		data, err := loadPuzzel("test.input")
		assert.Nil(t, err)
		assert.NotEmpty(t, data)
		require.Nil(t, err)

		rules, updates, err := parseData(data)
		assert.Nil(t, err)
		assert.NotEmpty(t, rules)
		assert.NotEmpty(t, updates)
		require.Nil(t, err)
		p := p()
		assert.Equal(t, false, p.has(47))
		p.rulesMapping(rules)
		assert.Equal(t, true, p.has(47))
		assert.Equal(t, 3, getMiddle([]int{1, 2, 3, 4, 5}))
		assert.Equal(t, []int{1, 3, 4}, swap([]int{1, 4, 3}, 1))
		assert.Equal(t, 1, p.failsAt([]int{61, 13, 29}))
		assert.Equal(t, true, p.checkOrder(swap([]int{61, 13, 29}, 1)))
		assert.Equal(t, true, p.checkOrder(p.fix([]int{75, 97, 47, 61, 53})))
		assert.Equal(t, 123, solve(rules, updates))
	})
}
