package event

import (
	"time"
)

// Dosage - Structure for dosage message
type Dosage struct {
	eventID string
	user    string
	carbs   int
	ts      time.Time
}
