package avl

type Visit func(interface{})

type Key interface {
	Less(Key) bool
	Equals(Key) bool
}

func NewTree() *Tree {
	return &Tree{}
}

type Tree struct {
	root *node
}

func (t *Tree) Clear() {
	t.root = nil
}
func (t *Tree) Has(k Key) bool {
	return lookup(t.root, k) != nil
}
func (t *Tree) Put(k Key, v interface{}) {
	t.root = insert(t.root, k, v)
}
func (t *Tree) Traverse(v Visit) {
	traverse(t.root, v)
}
func (t *Tree) Get(k Key) interface{} {
	n := lookup(t.root, k)
	if n == nil {
		return nil
	}
	return n.data
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

func insert(n *node, k Key, d interface{}) *node {
	if n == nil {
		n = &node{key: k, data: d, height: 0}
		return n
	}
	if k.Less(n.key) {
		n.left = insert(n.left, k, d)
	} else {
		n.right = insert(n.right, k, d)
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
func traverse(n *node, visit Visit) {
	if n == nil {
		return
	}
	traverse(n.left, visit)
	visit(n.data)
	traverse(n.right, visit)
}

func remove(n *node, k Key) *node {
	if n == nil {
		return n
	} else if k.Less(n.key) {
		n.left = remove(n.left, k)
	} else if !k.Equals(n.key) {
		n.right = remove(n.right, k)
	} else {
		if n.left == nil && n.right == nil {
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
			n.data = smallestRight.data
			n.right = remove(n.right, n.key)
		}
	}

	return rebalance(n)
}

func smallest(n *node) *node {
	small := n
	for small.left != nil {
		small = small.left
	}
	return small
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
