package rtsengine

import "testing"

func TestBasicUnitSyntax(t *testing.T) {
	a := Acre{}
	a.unit = &Farm{}

	if a.unit != a.unit {
		t.Error("Unit pointers are not equivalent!", a.unit)
	}

	var unitTest IUnit
	unitTest = &Farm{}
	a.unit = &WoodPile{}

	if a.unit == unitTest {
		t.Error("Unit pointers should not be equivalent ")
	}

	p := Pool{}
	p.Free(a.unit)
	farms := p.Farms(10)
	a.unit = farms[0]

}

func TestFarm(t *testing.T) {
	pool := Pool{}
	pool.Generate(5)
	farms := pool.Farms(10) // Force heap allocation for half of the units
	if len(farms) != 10 {
		t.Error("Length of returned array should have been 10 but got ", len(farms))
	}

	for i := range farms {
		if !farms[i].IsAllocated() {
			t.Error("Farm not allocated ", farms[i])
		}
	}

	for _, farm := range farms {
		pool.Free(farm)
	}

	for i := range farms {
		if farms[i].IsAllocated() {
			t.Error("Farm is still allocated ", farms[i])
		}
	}

}
