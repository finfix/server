package date

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const DateFormat = "2006-01-02"
const DateTimeFormat = "2006-01-02 15:04:05"

type Date struct {
	time.Time
}

func Unix(sec int64) Date {
	return Date{time.Unix(sec, 0)}
}

func Parse(s string) (Date, error) {
	date, err := time.Parse(DateFormat, s)
	if err != nil {
		return Date{}, err
	}
	return Date{date}, nil
}

func NewDate(year int, month time.Month, day int) Date {
	return Date{time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}

func Now() Date {
	return Date{time.Now()}
}

func (d Date) Format() string {
	return d.Time.Format(DateFormat)
}

func (d *Date) UnmarshalJSON(b []byte) (err error) {

	s := strings.Trim(string(b), "\"") // remove quotes
	if s == "null" || s == "" {
		return nil
	}

	d.Time, err = time.Parse(DateFormat, s)
	if err != nil {
		return err
	}

	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {

	if d.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%v"`, d.Format())), nil
}

func (d *Date) UnmarshalText(b []byte) (err error) {

	s := strings.Trim(string(b), "\"") // remove quotes
	if s == "null" || s == "" {
		return nil
	}

	d.Time, err = time.Parse(DateFormat, s)
	if err != nil {
		return err
	}

	return nil
}

func (d *Date) Scan(src any) error {
	if t, ok := src.(time.Time); ok {
		d.Time = t
	}
	return nil
}

func (d Date) Value() (driver.Value, error) {
	return d.Time, nil
}
