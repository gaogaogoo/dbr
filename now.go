package dbr

import (
	"database/sql/driver"
	"time"
)

// Now is a value that serializes to the current time in Local.
var Now = nowSentinel{}

const timeFormat = "2006-01-02 15:04:05.000000"

type nowSentinel struct{}

// Value implements a valuer for compatibility
func (n nowSentinel) Value() (driver.Value, error) {
	now := time.Now()
	if timeLocation != nil {
		now = now.In(timeLocation)
	}
	return now.Format(timeFormat), nil
}
