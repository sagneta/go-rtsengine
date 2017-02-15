package rtsengine

// HealthAndAttack maintains the health statistics of a unit.
type HealthAndAttack struct {

	// Init at Life == Hitpoints. Life reduces as hits are absorbed
	Life int

	// Default hit points for a particular unit. The more hit points
	// the more abuse a unit can maintain.
	HitPoints int

	// The number of points a single clash deducts from some other
	// units health
	AttackPoints int

	// The number of acres away a unit can attack.
	// 1 means only adjacent acres. 2 means anything two acres away
	// etc
	AttackRange int
}

// IPlayer
func (health *HealthAndAttack) life() int {
	return health.Life
}

// IsHealthy returns TRUE if this unit has >=40% of its Hitpoints
func (health *HealthAndAttack) IsHealthy() bool {
	return health.Life > int(float32(health.HitPoints)*0.40)
}

// IsPerfectlyHealthy returns TRUE if this unit has no damage
// at all.
func (health *HealthAndAttack) IsPerfectlyHealthy() bool {
	return health.Life >= health.HitPoints
}

// IsDead returns TRUE if this unit has expired.
func (health *HealthAndAttack) IsDead() bool {
	return health.Life <= 0
}

// ReduceHealth reduces the health of this unit by
// the AttackPoints held within otherUnit. The
// argument otherUnit presumably is the Health of some
// other attacking unit.
func (health *HealthAndAttack) ReduceHealth(otherUnit *HealthAndAttack) {
	health.Life -= otherUnit.AttackPoints
}

// IncreaseHealth will increase health using the AttackPoints
// Presumably this would be used by a medic or repair unit.
// Maxes out at HitPoints obviously.
func (health *HealthAndAttack) IncreaseHealth(otherUnit *HealthAndAttack) {
	health.Life += otherUnit.AttackPoints
	if health.Life > health.HitPoints {
		health.Life = health.AttackPoints
	}
}
