package meters_miles_conv

import "fmt"

type Meters float64
type Miles float64

const (
	MileInMeter Meters = 1609.344
)

func (m Meters) String() string {
	return fmt.Sprintf("%g meters", m)
}

func (m Miles) String() string {
	return fmt.Sprintf("%g miles", m)
}

func MetersToMiles(m Meters) Miles {
	return Miles(m / MileInMeter)
}

func MilesToMeters(m Miles) Meters {
	return Meters(m) * MileInMeter
}
