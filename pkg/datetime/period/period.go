package period

import (
	"time"

	"pkg/datetime/date"
	"pkg/errors"
	"pkg/proto/pbDatetime"
)

type Period string

// enums:"year,quarter,month,week,day"
const (
	Year    = Period("year")
	Quarter = Period("quarter")
	Month   = Period("month")
	Week    = Period("week")
	Day     = Period("day")
	All     = Period("all")
)

func (p *Period) ValidatePeriod() error {
	if p == nil {
		return nil
	}
	switch *p {
	case Year, Quarter, Month, Week, Day, All:
		return nil
	default:
		return errors.BadRequest.New("invalid period")
	}
}

func (p Period) ConvertToProto() pbDatetime.Period {
	switch p {
	case Year:
		return pbDatetime.Period_Year
	case Quarter:
		return pbDatetime.Period_Quarter
	case Month:
		return pbDatetime.Period_Month
	case Week:
		return pbDatetime.Period_Week
	case Day:
		return pbDatetime.Period_Day
	}
	return pbDatetime.Period_Month
}

func (p *Period) ConvertToOptionalProto() *pbDatetime.Period {
	if p == nil {
		return nil
	}
	period := p.ConvertToProto()
	return &period
}

type PbPeriod struct {
	*pbDatetime.Period
}

func (pb PbPeriod) ConvertToEnum() Period {
	if pb.Period == nil {
		return ""
	}
	switch *pb.Period {
	case pbDatetime.Period_Year:
		return Year
	case pbDatetime.Period_Quarter:
		return Quarter
	case pbDatetime.Period_Month:
		return Month
	case pbDatetime.Period_Week:
		return Week
	case pbDatetime.Period_Day:
		return Day
	}
	return ""
}

func (p PbPeriod) ConvertToOptionalEnum() *Period {
	if p.Period == nil {
		return nil
	}
	period := p.ConvertToEnum()
	return &period
}

// AtBeginOfPeriod возвращает начало периода
func GetTimeInterval(timeInPeriod time.Time, period Period) (dateFrom date.Date, dateTo date.Date) {
	switch period {
	case Day:
		dateFrom = date.NewDate(timeInPeriod.Year(), timeInPeriod.Month(), timeInPeriod.Day())
		dateTo = date.Date{dateFrom.AddDate(0, 0, 1)}
	case Month:
		dateFrom = date.NewDate(timeInPeriod.Year(), timeInPeriod.Month(), 1)
		dateTo = date.Date{dateFrom.AddDate(0, 1, 0)}
	case Week:
		secondsInWeek := 7 * 24 * 60 * 60
		// Получаем начало недели (понедельник)
		begin := (int(timeInPeriod.Unix())/secondsInWeek)*secondsInWeek - 3*24*60*60 // - 3 дня
		dateFrom = date.Unix(int64(begin))
		dateTo = date.Date{dateFrom.AddDate(0, 0, 7)}
	case Quarter:
		// Получаем номер квартала
		quarter := (timeInPeriod.Month()-1)/3 + 1
		// Получаем начало квартала
		dateFrom = date.NewDate(timeInPeriod.Year(), time.Month((quarter-1)*3+1), 1)
		dateTo = date.Date{dateFrom.AddDate(0, 3, 0)}
	case Year:
		dateFrom = date.NewDate(timeInPeriod.Year(), 1, 1)
		dateTo = date.Date{dateFrom.AddDate(1, 0, 0)}
	default:
		dateFrom, dateTo = date.NewDate(0, 0, 0), date.NewDate(0, 0, 0)
	}
	return dateFrom, dateTo
}
