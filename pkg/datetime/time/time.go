package time

import (
	"time"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"pkg/proto/pbDatetime"
)

type Time struct {
	time.Time
}

func (d Time) ConvertToProto() *pbDatetime.Timestamp {

	if d.IsZero() {
		return &pbDatetime.Timestamp{
			Timestamp: &timestamppb.Timestamp{},
		}
	}

	pb := &pbDatetime.Timestamp{
		Timestamp: timestamppb.New(d.Time),
	}

	// Получаем смещение относительно UTC
	_, offset := d.Zone()
	pb.Zone = int32(offset)

	return pb
}

func (d *Time) ConvertToOptionalProto() *pbDatetime.Timestamp {
	if d == nil || d.IsZero() {
		return nil
	}
	return d.ConvertToProto()
}

type PbTime struct {
	*pbDatetime.Timestamp
}

func (d PbTime) ConvertToTime() time.Time {
	var myTime time.Time
	if d.Timestamp.Timestamp == nil {
		return myTime
	}
	myTime = time.Unix(d.Timestamp.Timestamp.Seconds, 0)
	return myTime.In(time.FixedZone("", int(d.Zone)))
}

func (d PbTime) ConvertToOptionalTime() *time.Time {
	if d.Timestamp == nil {
		return nil
	}
	myTime := d.ConvertToTime()
	return &myTime
}
