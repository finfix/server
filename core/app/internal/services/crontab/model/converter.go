package model

import "core/app/proto/pbCrontab"

func (p *UpdateCurrenciesRes) ConvertToProto() *pbCrontab.UpdateCurrenciesRes {
	if p == nil {
		return nil
	}
	pb := &pbCrontab.UpdateCurrenciesRes{}
	pb.Rates = p.Rates
	return pb
}
