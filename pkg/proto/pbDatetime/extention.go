package pbDatetime

import (
	"pkg/errors"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	if t == nil || t.Timestamp == nil || t.Timestamp.Seconds == 0 {
		return []byte("null"), nil
	}
	myTime := time.Unix(t.Timestamp.Seconds, 0).In(time.FixedZone("", int(t.Zone)))
	return []byte("\"" + myTime.Format(time.RFC3339) + "\""), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {

	s := strings.Trim(string(data), "\"") // remove quotes
	if s == "null" || s == "" {
		return nil
	}

	if len(s) == 10 {
		s += "T00:00:00Z"
	}

	myTime, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	_, offset := myTime.Zone()

	t.Timestamp = &timestamppb.Timestamp{
		Seconds: myTime.Unix(),
		Nanos:   0,
	}
	t.Zone = int32(offset)
	return nil
}
