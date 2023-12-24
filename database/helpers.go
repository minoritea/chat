package database

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"time"

	"github.com/oklog/ulid/v2"
)

func IsRecordNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func NewID() string {
	return ulid.Make().String()
}

type Time time.Time

func (t *Time) Scan(src any) error {
	switch v := src.(type) {
	case time.Time:
		*t = Time(v)
		return nil
	case *time.Time:
		*t = Time(*v)
		return nil
	case string:
		tt, err := time.Parse(time.RFC3339Nano, v)
		if err != nil {
			return err
		}
		*t = Time(tt)
		return nil
	case *string:
		tt, err := time.Parse(time.RFC3339Nano, *v)
		if err != nil {
			return err
		}
		*t = Time(tt)
		return nil
	case nil:
		return nil
	default:
		return errors.New("invalid type")
	}
}
func (t Time) Value() (driver.Value, error) { return time.Time(t), nil }
func (t Time) Time() time.Time              { return time.Time(t) }
