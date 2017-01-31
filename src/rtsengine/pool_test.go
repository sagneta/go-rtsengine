package rtsengine

import "testing"

func TestFarm(t *testing.T) {
	pool := Pool{}
	pool.Generate(100)
	farms := pool.Farms(10)
	if len(farms) != 10 {
		t.Error("Length of returned array should have been 10 but got ", len(farms))
	}
}
