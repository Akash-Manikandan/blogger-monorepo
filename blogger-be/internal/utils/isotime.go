package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// ISOTime wraps time.Time for custom JSON serialization
type ISOTime time.Time

// MarshalJSON formats time as an ISO 8601 string
func (t ISOTime) MarshalJSON() ([]byte, error) {
	formatted := time.Time(t).Format(time.RFC3339)
	return []byte(`"` + formatted + `"`), nil
}

// UnmarshalJSON parses an ISO 8601 string into ISOTime
func (t *ISOTime) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	if err != nil {
		return err
	}
	*t = ISOTime(parsedTime)
	return nil
}

// Convert back to standard time.Time
func (t ISOTime) Time() time.Time {
	return time.Time(t)
}

// Implement `driver.Valuer` for storing in PostgreSQL
func (t ISOTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// Implement `sql.Scanner` for retrieving from PostgreSQL
func (t *ISOTime) Scan(value any) error {
	if value == nil {
		*t = ISOTime(time.Time{})
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*t = ISOTime(v)
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into ISOTime", value)
	}
}

func (t ISOTime) String() string {
	return time.Time(t).Format(time.RFC3339)
}
