package redblacktree

import (
	"strconv"
	"strings"
	"testing"

	"github.com/mikekonan/gods-generic/utils"
	"github.com/stretchr/testify/assert"
)

// For example, [1 2,3] -> "123"
func intSliceToString(arr []int) string {
	arrText := make([]string, len(arr))
	for i, v := range arr {
		arrText[i] = strconv.Itoa(v)
	}
	return strings.Join(arrText, "")
}

func TestRedBlackTreePut(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") // overwrite

	assert.Equal(t, 7, tree.size)
	assert.Equal(t, "1234567", intSliceToString(tree.Keys()))
	assert.Equal(t, "abcdefg", strings.Join(tree.Values(), ""))

	tests1 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{8, "", false},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := tree.Get(test[0].(int))
		assert.Equal(t, test[1], actualValue)
		assert.Equal(t, test[2], actualFound)
	}
}

func TestRedBlackTreeRemove(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") // overwrite

	tree.Remove(5)
	tree.Remove(6)
	tree.Remove(7)
	tree.Remove(8)
	tree.Remove(5)

	assert.Equal(t, 4, tree.size)
	assert.Equal(t, "1234", intSliceToString(tree.Keys()))
	assert.Equal(t, "abcd", strings.Join(tree.Values(), ""))

	tests2 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "", false},
		{6, "", false},
		{7, "", false},
		{8, "", false},
	}

	for _, test := range tests2 {
		actualValue, actualFound := tree.Get(test[0].(int))
		assert.Equal(t, test[1], actualValue)
		assert.Equal(t, test[2], actualFound)
	}

	tree.Remove(1)
	tree.Remove(4)
	tree.Remove(2)
	tree.Remove(3)
	tree.Remove(2)
	tree.Remove(2)

	assert.Equal(t, 0, tree.size)
	assert.Equal(t, 0, len(tree.Keys()))
	assert.Equal(t, 0, len(tree.Values()))
}

func TestRedBlackTreeLeftAndRight(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])

	assert.Nil(t, tree.Left())
	assert.Nil(t, tree.Right())

	tree.Put(1, "a")
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x") // overwrite
	tree.Put(2, "b")

	assert.Equal(t, 1, tree.Left().Key)
	assert.Equal(t, "x", tree.Left().Value)
	assert.Equal(t, 7, tree.Right().Key)
	assert.Equal(t, "g", tree.Right().Value)
}

func TestRedBlackTreeCeilingAndFloor(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])

	node, found := tree.Floor(0)
	assert.Nil(t, node)
	assert.False(t, found)

	node, found = tree.Ceiling(0)
	assert.Nil(t, node)
	assert.False(t, found)

	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")

	node, found = tree.Floor(4)
	assert.Equal(t, 4, node.Key)
	assert.True(t, found)

	node, found = tree.Floor(0)
	assert.Nil(t, node)
	assert.False(t, found)

	node, found = tree.Ceiling(4)
	assert.Equal(t, 4, node.Key)
	assert.True(t, found)

	node, found = tree.Ceiling(8)
	assert.Nil(t, node)
	assert.False(t, found)
}

func TestRedBlackTreeIteratorNextOnEmpty(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])
	it := tree.Iterator()
	assert.False(t, it.Next())
}

func TestRedBlackTreeIteratorPrevOnEmpty(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])
	it := tree.Iterator()
	assert.False(t, it.Prev())
}

func TestRedBlackTreeIterator1Next(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") // overwrite
	// │   ┌── 7
	// └── 6
	//     │   ┌── 5
	//     └── 4
	//         │   ┌── 3
	//         └── 2
	//             └── 1
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		assert.Equal(t, count, it.Key())
	}
	assert.Equal(t, count, tree.Size())
}

func TestRedBlackTreeIterator1Prev(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") // overwrite
	// │   ┌── 7
	// └── 6
	//     │   ┌── 5
	//     └── 4
	//         │   ┌── 3
	//         └── 2
	//             └── 1
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		assert.Equal(t, countDown, it.Key())
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator2Next(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		assert.Equal(t, count, it.Key())
	}
	assert.Equal(t, tree.Size(), count)
}

func TestRedBlackTreeIterator2Prev(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		assert.Equal(t, countDown, it.Key())
		countDown--
	}
	assert.Equal(t, 0, countDown)
}

func TestRedBlackTreeIterator3Next(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])
	tree.Put(1, "a")
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		assert.Equal(t, count, it.Key())
	}
	assert.Equal(t, tree.Size(), count)
}

func TestRedBlackTreeIterator3Prev(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])
	tree.Put(1, "a")
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		assert.Equal(t, countDown, it.Key())
		countDown--
	}
	assert.Equal(t, 0, countDown)
}

func TestRedBlackTreeIterator4Next(t *testing.T) {
	tree := NewWithComparator[int, int](utils.NumbersComparator[int])
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	// │           ┌── 27
	// │       ┌── 25
	// │       │   └── 22
	// │   ┌── 17
	// │   │   └── 15
	// └── 13
	//     │   ┌── 11
	//     └── 8
	//         │   ┌── 6
	//         └── 1
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		assert.Equal(t, count, it.Value())
	}
	assert.Equal(t, tree.Size(), count)
}

func TestRedBlackTreeIterator4Prev(t *testing.T) {
	tree := NewWithComparator[int, int](utils.NumbersComparator[int])
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	// │           ┌── 27
	// │       ┌── 25
	// │       │   └── 22
	// │   ┌── 17
	// │   │   └── 15
	// └── 13
	//     │   ┌── 11
	//     └── 8
	//         │   ┌── 6
	//         └── 1
	it := tree.Iterator()
	count := tree.Size()
	for it.Next() {
	}
	for it.Prev() {
		assert.Equal(t, count, it.Value())
		count--
	}
	assert.Equal(t, 0, count)
}

func TestRedBlackTreeIteratorBegin(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()

	assert.Nil(t, it.node)

	it.Begin()

	assert.Nil(t, it.node)

	for it.Next() {
	}

	it.Begin()

	assert.Nil(t, it.node)

	it.Next()
	assert.Equal(t, 1, it.Key())
	assert.Equal(t, "a", it.Value())
}

func TestRedBlackTreeIteratorEnd(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])
	it := tree.Iterator()

	assert.Nil(t, it.node)

	it.End()
	assert.Nil(t, it.node)

	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it.End()
	assert.Nil(t, it.node)

	it.Prev()
	assert.Equal(t, 3, it.Key())
	assert.Equal(t, "c", it.Value())
}

func TestRedBlackTreeIteratorFirst(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	assert.True(t, it.First())
	assert.Equal(t, 1, it.Key())
	assert.Equal(t, "a", it.Value())
}

func TestRedBlackTreeIteratorLast(t *testing.T) {
	tree := NewWithComparator[int, string](utils.NumbersComparator[int])
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	assert.True(t, it.Last())
	assert.Equal(t, 3, it.Key())
	assert.Equal(t, "c", it.Value())
}

func TestRedBlackTreeSerialization(t *testing.T) {
	tree := NewWithComparator[string, string](utils.StringComparator)
	tree.Put("c", "3")
	tree.Put("b", "2")
	tree.Put("a", "1")

	assert.Equal(t, 3, tree.Size())
	assert.Equal(t, "abc", strings.Join(tree.Keys(), ""))
	assert.Equal(t, "123", strings.Join(tree.Values(), ""))
	// TODO: this library hasn't implemented ToJSON() yet
}

func benchmarkGet(b *testing.B, tree *Tree[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, tree *Tree[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Put(n, struct{}{})
		}
	}
}

func benchmarkRemove(b *testing.B, tree *Tree[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Remove(n)
		}
	}
}

func BenchmarkRedBlackTreeGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewWithComparator[int, struct{}](utils.NumbersComparator[int])
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreeGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := NewWithComparator[int, struct{}](utils.NumbersComparator[int])
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreeGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := NewWithComparator[int, struct{}](utils.NumbersComparator[int])
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreeGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := NewWithComparator[int, struct{}](utils.NumbersComparator[int])
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreePut100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewWithComparator[int, struct{}](utils.NumbersComparator[int])
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreePut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := NewWithComparator[int, struct{}](utils.NumbersComparator[int])
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreePut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := NewWithComparator[int, struct{}](utils.NumbersComparator[int])
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreePut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := NewWithComparator[int, struct{}](utils.NumbersComparator[int])
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreeRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewWithComparator[int, struct{}](utils.NumbersComparator[int])
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkRedBlackTreeRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := NewWithComparator[int, struct{}](utils.NumbersComparator[int])
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkRedBlackTreeRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := NewWithComparator[int, struct{}](utils.NumbersComparator[int])
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkRedBlackTreeRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := NewWithComparator[int, struct{}](utils.NumbersComparator[int])
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}
