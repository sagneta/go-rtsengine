package rtsengine

/*
 Implements the  unit

*/

// Castle is an IUnit that maintains a military fortification
type Castle struct {
	Poolable
}

func (unit *Castle) name() string {
	return "Castle"
}
