package avl

import (
	"fmt"
	"testing"
)

func TestTreeGet(t *testing.T) {
	nt := NewTree()
	nt.root = buildTree(7)
	n := nt.Get(testKey(10))
	if n != nil {
		t.Errorf("Key 10 should not exist: %v", n)
	}

	n = nt.Get(testKey(4))
	i, ok := n.(int)
	if !ok || i != 4 {
		t.Errorf("Get should return 4, not %v", i)
	}
}
func TestTreeTraverse(t *testing.T) {
	nt := NewTree()
	nt.root = buildTree(7)
	ints := make([]int, 0)
	nodes := []int{1, 2, 3, 4, 5, 6, 7}

	nt.Traverse(func(d interface{}) {
		ints = append(ints, d.(int))
	})

	for _, err := range verifyTree(nt.root, nodes) {
		t.Error(err.Error())
	}
}

func TestTreeClear(t *testing.T) {
	nt := NewTree()
	nt.root = buildTree(7)
	nt.Clear()
	if nt.root != nil {
		t.Errorf("Root node should be null")
	}
}

func TestTreeHas(t *testing.T) {
	nt := NewTree()
	nt.root = buildTree(7)
	if nt.Has(testKey(10)) {
		t.Error("Key 10 should not exist")
	}

	if !nt.Has(testKey(4)) {
		t.Error("Key 4 should exist")
	}
}

func TestTreePut(t *testing.T) {
	nt := NewTree()
	nt.Put(testKey(4), 4)
	if !nt.Has(testKey(4)) {
		t.Error("Key 4 should exist")
	}
}

func TestMax(t *testing.T) {
	m := max(3, 2)
	if m != 3 {
		t.Errorf("Max should be 3, not %v", m)
	}
	m = max(2, 3)
	if m != 3 {
		t.Errorf("Max should be 2, not %v", m)
	}
	m = max(3, 3)
	if m != 3 {
		t.Errorf("Max should be 3, not %v", m)
	}
}

func TestHeight(t *testing.T) {
	h := height(nil)
	if h != -1 {
		t.Errorf("Height should be -1, not %v", h)
	}

	h = height(&node{height: 2})
	if h != 2 {
		t.Errorf("Height should be 2, not %v", h)
	}
}

func TestUpdateHeight(t *testing.T) {
	n := &node{height: 10}
	updateHeight(n)
	if n.height != 0 {
		t.Errorf("Height should be 0, not %v", n.height)
	}

	n.right = &node{height: 3}
	n.left = &node{height: 4}
	updateHeight(n)
	if n.height != 5 {
		t.Errorf("Height should be 5, not %v", n.height)
	}

}

func TestBalanceFactor(t *testing.T) {
	n := &node{}
	n.right = &node{height: 6}
	n.left = &node{height: 4}
	b := balanceFactor(n)
	if b != -2 {
		t.Errorf("Balance should be -2, not %v", b)
	}
	n.left.height = 8
	b = balanceFactor(n)
	if b != 2 {
		t.Errorf("Balance should be 2, not %v", b)
	}

}

func TestInsertRR(t *testing.T) {
	n := insert(nil, testKey(10), 10)
	if n.data.(int) != 10 {
		t.Errorf("Insert should have inserted 10")
	}
	n = insert(n, testKey(15), 15)
	n = insert(n, testKey(20), 20)
	if n.data.(int) != 15 {
		t.Errorf("Expected 5, not %v", n.data.(int))
	}
	if n.left.data.(int) != 10 {
		t.Errorf("Expected 10, not %v", n.data.(int))
	}
	if n.right.data.(int) != 20 {
		t.Errorf("Expected 20, not %v", n.data.(int))
	}
}

func TestInsertLL(t *testing.T) {
	n := insert(nil, testKey(20), 20)
	if n.data.(int) != 20 {
		t.Errorf("Insert should have inserted 20")
	}
	n = insert(n, testKey(15), 15)
	n = insert(n, testKey(10), 10)
	if n.data.(int) != 15 {
		t.Errorf("Expected 15, not %v", n.data.(int))
	}
	if n.left.data.(int) != 10 {
		t.Errorf("Expected 10, not %v", n.data.(int))
	}
	if n.right.data.(int) != 20 {
		t.Errorf("Expected 20, not %v", n.data.(int))
	}
}

func TestInsertRL(t *testing.T) {
	n := insert(nil, testKey(10), 10)
	if n.data.(int) != 10 {
		t.Errorf("Insert should have inserted 10")
	}
	n = insert(n, testKey(20), 20)
	n = insert(n, testKey(15), 15)
	if n.data.(int) != 15 {
		t.Errorf("Expected 15, not %v", n.data.(int))
	}
	if n.left.data.(int) != 10 {
		t.Errorf("Expected 10, not %v", n.data.(int))
	}
	if n.right.data.(int) != 20 {
		t.Errorf("Expected 20, not %v", n.data.(int))
	}
}

func TestInsertLR(t *testing.T) {
	n := insert(nil, testKey(20), 20)
	if n.data.(int) != 20 {
		t.Errorf("Insert should have inserted 20")
	}
	n = insert(n, testKey(10), 10)
	n = insert(n, testKey(15), 15)
	if n.data.(int) != 15 {
		t.Errorf("Expected 15, not %v", n.data.(int))
	}
	if n.left.data.(int) != 10 {
		t.Errorf("Expected 10, not %v", n.data.(int))
	}
	if n.right.data.(int) != 20 {
		t.Errorf("Expected 20, not %v", n.data.(int))
	}
}

func TestLookup(t *testing.T) {
	n := insert(nil, testKey(20), 20)
	n = insert(n, testKey(30), 30)
	n = insert(n, testKey(40), 40)
	n = insert(n, testKey(50), 50)
	n = insert(n, testKey(60), 60)

	v := lookup(n, testKey(70))
	if v != nil {
		t.Errorf("Lookup should fail")
	}

	v = lookup(n, testKey(50))
	if v == nil {
		t.Errorf("Lookup should be successful")
	}
	if v != nil && v.data.(int) != 50 {
		t.Errorf("Lookup should return 50")
	}

	v = lookup(n, testKey(20))
	if v == nil {
		t.Errorf("Lookup should be successful")
	}
	if v != nil && v.data.(int) != 20 {
		t.Errorf("Lookup should return 20")
	}

}
func TestRemoveNonExisting(t *testing.T) {
	n := buildTree(7)

	n = remove(n, testKey(20))
	nodes := []int{1, 2, 3, 4, 5, 6, 7}
	for _, err := range verifyTree(n, nodes) {
		t.Error(err.Error())
	}
}
func TestRemoveLeaf(t *testing.T) {
	n := buildTree(7)

	n = remove(n, testKey(7))
	nodes := []int{1, 2, 3, 4, 5, 6}
	for _, err := range verifyTree(n, nodes) {
		t.Error(err.Error())
	}
}
func TestRemoveLeftChild(t *testing.T) {
	n := buildTree(7)
	n = remove(n, testKey(7))
	n = remove(n, testKey(6))
	nodes := []int{1, 2, 3, 4, 5}
	for _, err := range verifyTree(n, nodes) {
		t.Error(err.Error())
	}
}

func TestRemoveRightChild(t *testing.T) {
	n := buildTree(7)
	n = remove(n, testKey(5))
	n = remove(n, testKey(6))
	nodes := []int{1, 2, 3, 4, 7}
	for _, err := range verifyTree(n, nodes) {
		t.Error(err.Error())
	}
}
func TestRemoveRoot(t *testing.T) {
	n := buildTree(7)
	n = remove(n, testKey(4))
	nodes := []int{1, 2, 3, 5, 6, 7}
	for _, err := range verifyTree(n, nodes) {
		t.Error(err.Error())
	}
}

func TestRebalance(t *testing.T) {
	n := rebalance(nil)
	if n != nil {
		t.Errorf("Rebalance should return nil")
	}
}

func TestTraverse(t *testing.T) {
	n := buildTree(7)
	nodes := []int{1, 2, 3, 4, 5, 6, 7}
	for _, err := range verifyTree(n, nodes) {
		t.Error(err.Error())
	}
}

func buildTree(max int) *node {
	var n *node
	if max < 1 {
		return nil
	}
	for i := 1; i <= max; i++ {
		n = insert(n, testKey(i), i)
	}
	return n
}

func verifyTree(n *node, vals []int) []error {
	var errors []error
	ints := make([]int, 0)

	traverse(n, func(d interface{}) {
		ints = append(ints, d.(int))
	})

	intsLen := len(ints)
	valsLen := len(vals)

	if valsLen != intsLen {
		msg := "Expected tree size %v, not %v"
		err := fmt.Errorf(msg, valsLen, intsLen)
		errors = append(errors, err)
		return errors
	}

	for i, v := range vals {
		if ints[i] != v {
			msg := "Incorrect traversal at: %v, %v"
			err := fmt.Errorf(msg, i, ints[i])
			errors = append(errors, err)
		}
	}

	return errors
}

type testKey int

func (tk testKey) Less(rhs Key) bool {
	return tk < rhs.(testKey)
}
func (tk testKey) Equals(rhs Key) bool {
	return tk == rhs.(testKey)
}
