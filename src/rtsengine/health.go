package rtsengine

// Health maintains the health statistics of a unit.
type Health struct {

	// Init at life == hitPoints. Life reduces as hits are absorbed
	Life float32

	// Default hit points for a particular unit. The more hit points
	// the more abuse a unit can maintain.
	HitPoints float32

	// The number of points a single clash deducts from some other
	// units health
	AttackPoints float32
}

// IsHealthy returns TRUE if this unit has >=40% of its Hitpoints
func (health *Health) IsHealthy() bool {
	return health.Life > (health.HitPoints * 0.40)
}

// IsPerfectlyHealthy returns TRUE if this unit has no damage
// at all.
func (health *Health) IsPerfectlyHealthy() bool {
	return health.Life >= health.HitPoints
}

// IsDead returns TRUE if this unit has expired.
func (health *Health) IsDead() bool {
	return health.Life <= 0.0
}

// ReduceHealth reduces the health of this unit by
// the AttackPoints held within otherUnit. The
// argument otherUnit presumably is the Health of some
// other attacking unit.
func (health *Health) ReduceHealth(otherUnit *Health) {
	health.Life -= otherUnit.AttackPoints
}

// IncreaseHealth will increase health using the AttackPoints
// Presumably this would be used by a medic or repair unit.
// Maxes out at HitPoints obviously.
func (health *Health) IncreaseHealth(otherUnit *Health) {
	health.Life += otherUnit.AttackPoints
	if health.Life > health.HitPoints {
		health.Life = health.AttackPoints
	}
}
