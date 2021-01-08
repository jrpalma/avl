package avl

type node struct {
	key    Key
	data   interface{}
	height int16
	left   *node
	right  *node
}
