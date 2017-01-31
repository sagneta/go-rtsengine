package rtsengine

/*
 Implements the  unit

*/

// Tower is an IUnit that maintains a military watch tower with some defensive capability
type Tower struct {
	Poolable
}

func (unit *Tower) name() string {
	return "Tower"
}
