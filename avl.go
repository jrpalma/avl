package avl

//Key Interface used to keep a balanced AVL tree.
type Key interface {
	Less(Key) bool
	Equals(Key) bool
}

//Visit Function type used to visit all the nodes in a tree.
//Return true to continue visiting or fale to stop.
type Visit func(Key) bool

//NewTree Creates and returns a new AVL tree
func NewTree() *Tree {
	return &Tree{}
}

//Tree An AVL tree to insert, search, and delete in O(log n)
type Tree struct {
	root *node
	size int
}

//Clear Clears all the nodes in the AVL tree
func (t *Tree) Clear() {
	t.root = nil
}

//Size Returns the number of nodes in the AVL tree
func (t *Tree) Size() int {
	return t.size
}

//Has Returns true if the tree contains the Key k
func (t *Tree) Has(k Key) bool {
	return lookup(t.root, k) != nil
}

//Put Inserts a node with Key k into the tree. The node is not inserted
//if a node with the Key k alredy exist.
func (t *Tree) Put(k Key) {
	t.root = t.insert(t.root, k)
}

//Del Deletes the node with the Key k.
func (t *Tree) Del(k Key) {
	t.root = t.remove(t.root, k)
}

//Get Searches for node with the Key k and returns the key.
//It returns nil if the Key cannot be found in the tree.
func (t *Tree) Get(k Key) Key {
	n := lookup(t.root, k)
	if n == nil {
		return nil
	}
	return n.key
}

//VisitAscending Traverse the tree from the smallest to the largest Key by
//calling the Visit function v.
func (t *Tree) VisitAscending(v Visit) {
	visitAscending(t.root, v)
}

//VisitDescending Traverse the tree from the largest to the smalles Key by
//calling the Visit function v.
func (t *Tree) VisitDescending(v Visit) {
	visitDescending(t.root, v)
}

func leftHeavy(bf int16) bool {
	return bf > 0
}
func rightHeavy(bf int16) bool {
	return bf < 0
}

func max(a int16, b int16) int16 {
	if a < b {
		return b
	}
	return a
}

func height(n *node) int16 {
	if n == nil {
		return -1
	}
	return n.height
}

func updateHeight(n *node) {
	if n != nil {
		n.height = 1 + max(height(n.left), height(n.right))
	}
}

func balanceFactor(n *node) int16 {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

func smallest(n *node) *node {
	small := n
	for small.left != nil {
		small = small.left
	}
	return small
}

func (t *Tree) insert(n *node, k Key) *node {
	if n == nil {
		n = &node{key: k, height: 0}
		t.size++
		return n
	}
	if k.Less(n.key) {
		n.left = t.insert(n.left, k)
	} else {
		n.right = t.insert(n.right, k)
	}

	return rebalance(n)
}

func (t *Tree) remove(n *node, k Key) *node {
	if n == nil {
		return n
	} else if k.Less(n.key) {
		n.left = t.remove(n.left, k)
	} else if !k.Equals(n.key) {
		n.right = t.remove(n.right, k)
	} else {
		if n.left == nil && n.right == nil {
			t.size--
			return nil
		}
		if n.left == nil && n.right != nil {
			n = n.right
		}
		if n.left != nil && n.right == nil {
			n = n.left
		}
		if n.left != nil && n.right != nil {
			smallestRight := smallest(n.right)
			n.key = smallestRight.key
			n.right = t.remove(n.right, n.key)
			t.size--
		}
	}

	return rebalance(n)
}

func lookup(n *node, k Key) *node {
	if n == nil {
		return n
	}
	if k.Less(n.key) {
		return lookup(n.left, k)
	} else if k.Equals(n.key) {
		return n
	}
	return lookup(n.right, k)
}

func visitAscending(n *node, visit Visit) bool {
	if n == nil {
		return true
	}
	ok := visitAscending(n.left, visit)
	if !ok {
		return ok
	}
	ok = visit(n.key)
	if !ok {
		return ok
	}
	return visitAscending(n.right, visit)
}

func visitDescending(n *node, visit Visit) bool {
	if n == nil {
		return true
	}
	ok := visitDescending(n.right, visit)
	if !ok {
		return ok
	}
	ok = visit(n.key)
	if !ok {
		return ok
	}
	return visitDescending(n.left, visit)
}

func rebalance(n *node) *node {
	if n == nil {
		return n
	}

	updateHeight(n)

	nbf := balanceFactor(n)
	if nbf >= -1 && nbf <= 1 {
		return n
	}

	lbf := balanceFactor(n.left)
	rbf := balanceFactor(n.right)

	if rightHeavy(nbf) && rightHeavy(rbf) {
		n = rrRotation(n)
	} else if rightHeavy(nbf) && leftHeavy(rbf) {
		n = rlRotation(n)
	} else if leftHeavy(nbf) && leftHeavy(lbf) {
		n = llRotation(n)
	} else if leftHeavy(nbf) && rightHeavy(lbf) {
		n = lrRotation(n)
	}

	return n
}

func llRotation(n *node) *node {
	//fmt.Printf("LL Rotation\n")
	var x, y, z *node
	x = n
	y = x.left
	z = y.left
	x.left = y.right
	y.left = z
	y.right = x
	updateHeight(x)
	updateHeight(y)
	return y
}

func rrRotation(n *node) *node {
	//fmt.Printf("RR Rotation\n")
	var x, y, z *node
	x = n
	y = x.right
	z = y.right
	x.right = y.left
	y.left = x
	y.right = z
	updateHeight(x)
	updateHeight(y)
	return y
}

func rlRotation(n *node) *node {
	//fmt.Printf("RL Rotation\n")
	var x, y, z *node
	x = n
	y = x.right
	z = y.left
	z.left = x
	z.right = y
	x.right = nil
	y.left = nil
	updateHeight(y)
	updateHeight(x)
	updateHeight(z)
	return z
}

func lrRotation(n *node) *node {
	//fmt.Printf("LR Rotation\n")
	var x, y, z *node
	x = n
	y = x.left
	z = y.right
	z.left = y
	z.right = x
	x.right = nil
	y.right = nil
	updateHeight(y)
	updateHeight(x)
	updateHeight(z)
	return z
}
