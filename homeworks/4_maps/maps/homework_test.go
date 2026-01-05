package main

import (
	"cmp"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type OrderedMap[T cmp.Ordered] struct {
	count int
	root  *Node[T]
}

type Node[T cmp.Ordered] struct {
	key   T
	value any
	left  *Node[T]
	right *Node[T]
}

func NewOrderedMap[T cmp.Ordered]() OrderedMap[T] {
	return OrderedMap[T]{}
}

func (m *OrderedMap[T]) Insert(key T, value any) {
	m.root = m.insert(key, value, m.root)
}

func (m *OrderedMap[T]) insert(key T, value any, node *Node[T]) *Node[T] {
	if node == nil {
		m.count++
		return &Node[T]{
			key:   key,
			value: value,
		}
	}
	if key < node.key {
		node.left = m.insert(key, value, node.left)
	}
	if key > node.key {
		node.right = m.insert(key, value, node.right)
	}
	node.value = value
	return node
}

func (m *OrderedMap[T]) Erase(key T) {
	if m.Size() == 0 {
		return
	}

	m.erase(key, m.root)
}

func (m *OrderedMap[T]) erase(key T, node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = m.erase(key, node.left)
		return node
	}

	if key > node.key {
		node.right = m.erase(key, node.right)
		return node
	}

	m.count--

	if node.left == nil {
		return node.right
	}

	if node.right == nil {
		return node.left
	}

	// target - min(left) leaf from right subtree
	target := m.min(node.right)

	node.right = m.erase(target.key, node.right)

	node.key = target.key
	node.value = target.value

	return node
}

func (m *OrderedMap[T]) min(node *Node[T]) *Node[T] {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (m *OrderedMap[T]) Contains(key T) bool {
	for node := m.root; node != nil; {
		if key == node.key {
			return true
		}
		if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}

	return false
}

func (m *OrderedMap[T]) Size() int {
	return m.count
}

func (m *OrderedMap[T]) ForEach(action func(key T, value any)) {
	if m.Size() == 0 {
		return
	}

	stack := make([]*Node[T], 0)
	current := m.root

	for current != nil || len(stack) != 0 {
		if current != nil {
			stack = append(stack, current)
			current = current.left
		} else {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			action(current.key, current.value)
			current = current.right
		}
	}
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap[int]()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key int, _ any) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key int, _ any) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
