package converter

import (
	"core/app/proto/pbCrontab"
	"jsonapi/app/internal/services/crontab/model"
)

type PbUpdateCurrenciesRes struct {
	*pbCrontab.UpdateCurrenciesRes
}

func (pb PbUpdateCurrenciesRes) ConvertToStruct() model.UpdateCurrenciesRes {
	var p model.UpdateCurrenciesRes
	p.Rates = pb.Rates
	return p
}
