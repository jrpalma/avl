package avl

import (
	"fmt"
	"testing"
)

type Record struct {
	Key  int
	Data int
}

func (r *Record) Less(k Key) bool {
	return r.Key < k.(*Record).Key
}
func (r *Record) Equals(k Key) bool {
	return r.Key == k.(*Record).Key
}

type RecordTree struct {
	Tree
}

func (rt *RecordTree) Get(key Key) *Record {
	v := rt.Tree.Get(key)
	if v == nil {
		return nil
	}
	return v.(*Record)
}

func TestRecord(t *testing.T) {
	nt := &Tree{}
	key := 10
	nt.Put(&Record{Key: key, Data: 10})
	v := nt.Get(&Record{Key: key})
	if v == nil {
		t.Errorf("Failed to get %v", key)
	}
	r := v.(*Record)
	if r.Data != 10 {
		t.Errorf("Data should be 10, not %v", r.Data)
	}

}

func TestRecordTree(t *testing.T) {
	nt := &RecordTree{}
	key := 10
	nt.Put(&Record{Key: key, Data: 20})
	v := nt.Get(&Record{Key: key})
	if v == nil {
		t.Errorf("Failed to get %v", key)
	}
	if v.Data != 20 {
		t.Errorf("Data should be 10, not %v", v.Data)
	}
}

func TestTreeGet(t *testing.T) {
	nt := NewTree()
	nt.root = buildTree(7)
	n := nt.Get(testKey(10))
	if n != nil {
		t.Errorf("Key 10 should not exist: %v", n)
	}

	n = nt.Get(testKey(4))
	i, ok := n.(testKey)
	if !ok || i != 4 {
		t.Errorf("Get should return 4, not %v", i)
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
	nt.Put(testKey(4))
	if !nt.Has(testKey(4)) {
		t.Error("Key 4 should exist")
	}
}

func TestTreeDel(t *testing.T) {
	nt := NewTree()
	nt.Put(testKey(4))

	nt.Del(testKey(4))
	if nt.Has(testKey(4)) {
		t.Error("Key 4 should not exist")
	}
}
func TestTreeSize(t *testing.T) {
	nt := NewTree()

	if nt.Size() != 0 {
		t.Errorf("Empty tree should have 0 size, not %v", nt.Size())
	}

	nt.Put(testKey(4))
	if nt.Size() != 1 {
		t.Errorf("Expected tree size 1, not %v", nt.Size())
	}

	nt.Del(testKey(4))
	if nt.Size() != 0 {
		t.Errorf("Empty tree should have 0 size, not %v", nt.Size())
	}
}

func TestTreeVisitAscending(t *testing.T) {
	nt := NewTree()
	nt.root = buildTree(7)
	ints := make([]int, 0)

	nt.VisitAscending(func(d Key) {
		ints = append(ints, int(d.(testKey)))
	})

	expected := []int{1, 2, 3, 4, 5, 6, 7}

	if len(expected) != len(ints) {
		t.Errorf("Expected nodes %#v, not %#v", expected, ints)
		return
	}
	for i, v := range ints {
		if expected[i] != v {
			t.Errorf("Expected %v, not %v", expected[i], v)
		}
	}
}

func TestTreeVisitDescending(t *testing.T) {
	nt := NewTree()
	nt.root = buildTree(7)
	ints := make([]int, 0)

	nt.VisitDescending(func(d Key) {
		ints = append(ints, int(d.(testKey)))
	})

	expected := []int{7, 6, 5, 4, 3, 2, 1}

	if len(expected) != len(ints) {
		t.Errorf("Expected nodes %#v, not %#v", expected, ints)
		return
	}
	for i, v := range ints {
		if expected[i] != v {
			t.Errorf("Expected %v, not %v", expected[i], v)
		}
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
	nt := &Tree{}
	n := nt.insert(nil, testKey(10))
	if n.key.(testKey) != 10 {
		t.Errorf("Insert should have inserted 10")
	}
	n = nt.insert(n, testKey(15))
	n = nt.insert(n, testKey(20))
	if n.key.(testKey) != 15 {
		t.Errorf("Expected 5, not %v", n.key.(testKey))
	}
	if n.left.key.(testKey) != 10 {
		t.Errorf("Expected 10, not %v", n.key.(testKey))
	}
	if n.right.key.(testKey) != 20 {
		t.Errorf("Expected 20, not %v", n.key.(testKey))
	}
}

func TestInsertLL(t *testing.T) {
	nt := &Tree{}
	n := nt.insert(nil, testKey(20))
	if n.key.(testKey) != 20 {
		t.Errorf("Insert should have inserted 20")
	}
	n = nt.insert(n, testKey(15))
	n = nt.insert(n, testKey(10))
	if n.key.(testKey) != 15 {
		t.Errorf("Expected 15, not %v", n.key.(testKey))
	}
	if n.left.key.(testKey) != 10 {
		t.Errorf("Expected 10, not %v", n.key.(testKey))
	}
	if n.right.key.(testKey) != 20 {
		t.Errorf("Expected 20, not %v", n.key.(testKey))
	}
}

func TestInsertRL(t *testing.T) {
	nt := &Tree{}
	n := nt.insert(nil, testKey(10))
	if n.key.(testKey) != 10 {
		t.Errorf("Insert should have inserted 10")
	}
	n = nt.insert(n, testKey(20))
	n = nt.insert(n, testKey(15))
	if n.key.(testKey) != 15 {
		t.Errorf("Expected 15, not %v", n.key.(testKey))
	}
	if n.left.key.(testKey) != 10 {
		t.Errorf("Expected 10, not %v", n.key.(testKey))
	}
	if n.right.key.(testKey) != 20 {
		t.Errorf("Expected 20, not %v", n.key.(testKey))
	}
}

func TestInsertLR(t *testing.T) {
	nt := &Tree{}
	n := nt.insert(nil, testKey(20))
	if n.key.(testKey) != 20 {
		t.Errorf("Insert should have inserted 20")
	}
	n = nt.insert(n, testKey(10))
	n = nt.insert(n, testKey(15))
	if n.key.(testKey) != 15 {
		t.Errorf("Expected 15, not %v", n.key.(testKey))
	}
	if n.left.key.(testKey) != 10 {
		t.Errorf("Expected 10, not %v", n.key.(testKey))
	}
	if n.right.key.(testKey) != 20 {
		t.Errorf("Expected 20, not %v", n.key.(testKey))
	}
}

func TestLookup(t *testing.T) {
	nt := &Tree{}
	n := nt.insert(nil, testKey(20))
	n = nt.insert(n, testKey(30))
	n = nt.insert(n, testKey(40))
	n = nt.insert(n, testKey(50))
	n = nt.insert(n, testKey(60))

	v := lookup(n, testKey(70))
	if v != nil {
		t.Errorf("Lookup should fail")
	}

	v = lookup(n, testKey(50))
	if v == nil {
		t.Errorf("Lookup should be successful")
	}
	if v != nil && v.key.(testKey) != 50 {
		t.Errorf("Lookup should return 50")
	}

	v = lookup(n, testKey(20))
	if v == nil {
		t.Errorf("Lookup should be successful")
	}
	if v != nil && v.key.(testKey) != 20 {
		t.Errorf("Lookup should return 20")
	}

}
func TestRemoveNonExisting(t *testing.T) {
	nt := &Tree{}
	nt.root = buildTree(7)

	nt.root = nt.remove(nt.root, testKey(20))
	nodes := []int{1, 2, 3, 4, 5, 6, 7}
	for _, err := range verifyTree(nt.root, nodes) {
		t.Error(err.Error())
	}
}
func TestRemoveLeaf(t *testing.T) {
	nt := &Tree{}
	nt.root = buildTree(7)

	nt.root = nt.remove(nt.root, testKey(7))
	nodes := []int{1, 2, 3, 4, 5, 6}
	for _, err := range verifyTree(nt.root, nodes) {
		t.Error(err.Error())
	}
}
func TestRemoveLeftChild(t *testing.T) {
	nt := &Tree{}
	nt.root = buildTree(7)
	nt.root = nt.remove(nt.root, testKey(7))
	nt.root = nt.remove(nt.root, testKey(6))
	nodes := []int{1, 2, 3, 4, 5}
	for _, err := range verifyTree(nt.root, nodes) {
		t.Error(err.Error())
	}
}

func TestRemoveRightChild(t *testing.T) {
	nt := &Tree{}
	nt.root = buildTree(7)
	nt.root = nt.remove(nt.root, testKey(5))
	nt.root = nt.remove(nt.root, testKey(6))
	nodes := []int{1, 2, 3, 4, 7}
	for _, err := range verifyTree(nt.root, nodes) {
		t.Error(err.Error())
	}
}
func TestRemoveRoot(t *testing.T) {
	nt := &Tree{}
	nt.root = buildTree(7)
	nt.root = nt.remove(nt.root, testKey(4))
	nodes := []int{1, 2, 3, 5, 6, 7}
	for _, err := range verifyTree(nt.root, nodes) {
		t.Error(err.Error())
	}
}

func TestRebalance(t *testing.T) {
	n := rebalance(nil)
	if n != nil {
		t.Errorf("Rebalance should return nil")
	}
}

func TestVisitAscending(t *testing.T) {
	n := buildTree(7)
	ints := make([]int, 0)
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	visitAscending(n, func(d Key) {
		ints = append(ints, int(d.(testKey)))
	})

	if len(expected) != len(ints) {
		t.Errorf("Expected nodes %#v, not %#v", expected, ints)
		return
	}
	for i, v := range ints {
		if expected[i] != v {
			t.Errorf("Expected %v, not %v", expected[i], v)
		}
	}
}

func TestVisitDescending(t *testing.T) {
	n := buildTree(7)
	ints := make([]int, 0)
	expected := []int{7, 6, 5, 4, 3, 2, 1}
	visitDescending(n, func(d Key) {
		ints = append(ints, int(d.(testKey)))
	})

	if len(expected) != len(ints) {
		t.Errorf("Expected nodes %#v, not %#v", expected, ints)
		return
	}
	for i, v := range ints {
		if expected[i] != v {
			t.Errorf("Expected %v, not %v", expected[i], v)
		}
	}
}

func buildTree(max int) *node {
	nt := &Tree{}
	if max < 1 {
		return nil
	}
	for i := 1; i <= max; i++ {
		nt.root = nt.insert(nt.root, testKey(i))
	}
	return nt.root
}

func verifyTree(n *node, vals []int) []error {
	var errors []error
	ints := make([]int, 0)

	visitAscending(n, func(d Key) {
		ints = append(ints, int(d.(testKey)))
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
