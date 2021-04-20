package avl

import (
	"fmt"
)

func NewUserStore() *UserStore {
	return &UserStore{tree: &Tree{}}
}

type UserVisit func(*User) bool

type User struct {
	ID  uint64
	Age uint8
}

func (u *User) Less(k Key) bool {
	return u.ID < k.(*User).ID
}
func (u *User) Equals(k Key) bool {
	return u.ID == k.(*User).ID
}

type UserStore struct {
	tree *Tree
}

func (us *UserStore) Get(id uint64) *User {
	v := us.tree.Get(&User{ID: id})
	if v == nil {
		return nil
	}
	return v.(*User)
}
func (us *UserStore) Size() int {
	return us.tree.Size()
}
func (us *UserStore) Clear() {
	us.tree.Clear()
}
func (us *UserStore) Has(id uint64) bool {
	return us.tree.Has(&User{ID: id})
}
func (us *UserStore) Del(id uint64) {
	us.tree.Del(&User{ID: id})
}
func (us *UserStore) Put(u *User) {
	us.tree.Put(u)
}
func (us *UserStore) VisitAscending(v UserVisit) {
	us.tree.VisitAscending(func(k Key) bool {
		v(k.(*User))
		return true
	})
}
func (us *UserStore) VisitDescending(v UserVisit) {
	us.tree.VisitDescending(func(k Key) bool {
		v(k.(*User))
		return true
	})
}

// ExampleUsage Shows how this package can be used
func Example_usage() {

	us := NewUserStore()
	for i := 1; i <= 256; i++ {
		us.Put(&User{ID: uint64(i)})
	}

	if us.Size() != 256 {
		panic("Size should be 256")
	}

	user := us.Get(128)
	if user == nil {
		panic("Should have user with ID 128")
	}

	fmt.Printf("%+v\n", user)
	// Output: &{ID:128 Age:0}

	if !us.Has(128) {
		panic("Should have user ID 128")
	}

	us.Del(128)

	if us.Has(128) {
		panic("Should not have user ID 128")
	}

	us.Clear()
}
