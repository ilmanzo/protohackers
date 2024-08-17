package problem06

import (
	"math"
	"time"
)

type (
	// [Road]tickets
	ticketQueue chan *Ticket
)

func checkViolation(o observation, past []*observation, limit float64) *Ticket {
	for _, prev := range past {
		// Check prev timestamp is within a day (86.4k secs)
		if o.timestamp.Sub(prev.timestamp).Abs() > time.Second*86400 {
			continue
		}

		// Calc speed
		miles := math.Abs(float64(prev.mile) - float64(o.mile))
		dur := prev.timestamp.Sub(o.timestamp).Abs()
		speed := miles / dur.Hours()

		if speed > limit+0.5 {
			first, second := orderObservations(o, *prev)
			return &Ticket{
				Plate:      o.plate,
				Speed:      uint16(speed * 100),
				Mile1:      first.mile,
				Timestamp1: UnixTime(first.timestamp.Unix()),
				Mile2:      second.mile,
				Timestamp2: UnixTime(second.timestamp.Unix()),
			}
		}
	}
	return nil
}

func orderObservations(obv1, obv2 observation) (earlier, later observation) {
	if obv1.timestamp.Before(obv2.timestamp) {
		return obv1, obv2
	}
	return obv2, obv1
}