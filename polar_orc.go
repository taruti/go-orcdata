package orcdata

func AllowanceToSpeed(allowance float64) float64 {
	return 3600 / allowance
}

// OrcAllowance numbers are in allowances = seconds per nautical mile.
// For knots use AllowanceToSpeed.
type OrcAllowance struct {
	WindSpeeds []int32
	WingAngles []int32
	// R* are SOG
	R52  [9]float64
	R60  [9]float64
	R75  [9]float64
	R90  [9]float64
	R110 [9]float64
	R120 [9]float64
	R135 [9]float64
	R150 [9]float64
	// DW* are VMG
	DW165     [9]float64
	DW180     [9]float64
	BeatAngle [9]float64
	GybeAngle [9]float64
	// Beat is VMG allowance at BeatAngle at selected wind
	Beat [9]float64
	// Run is VMG allowance at RunAngle at selected wind
	Run [9]float64
	// Scoring
	// Scroring DW180, DW165, DW150, WL, CR, OC
}
