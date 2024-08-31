package datetime

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const DateTimeFormat = "2006-01-02T15:04:05-0700"

type Time struct {
	time.Time
}

func (d *Time) UnmarshalJSON(b []byte) (err error) {

	s := strings.Trim(string(b), "\"") // remove quotes
	if s == JSONNull || s == "" {
		return nil
	}

	d.Time, err = time.Parse(DateTimeFormat, s)
	if err != nil {
		return err
	}

	return nil
}

func (d Time) MarshalJSON() ([]byte, error) {

	if d.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%v"`, d.Format(DateTimeFormat))), nil
}

func (d *Time) UnmarshalText(b []byte) (err error) {

	s := strings.Trim(string(b), "\"") // remove quotes
	if s == JSONNull || s == "" {
		return nil
	}

	d.Time, err = time.Parse(DateTimeFormat, s)
	if err != nil {
		return err
	}

	return nil
}

func (d *Time) Scan(src any) error {
	if t, ok := src.(time.Time); ok {
		d.Time = t
	}
	return nil
}

func (d *Time) Value() (driver.Value, error) {
	return d.Time, nil
}
