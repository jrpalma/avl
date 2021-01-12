package avl

type node struct {
	key    Key
	height int16
	left   *node
	right  *node
	parent *node
}
