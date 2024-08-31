package testUtils

import (
	"testing"

	"server/pkg/errors"
)

func Test_checkError(t *testing.T) {

	firstType := errors.New("first")
	secondType := errors.New("second")

	type args struct {
		wantErr       error
		gotErr        error
		compareErrors bool
	}
	tests := []struct {
		name            string
		args            args
		wantErr         error
		wasPassedErrors bool
	}{
		{
			name: "1. Обе ошибки пустые",
			args: args{
				wantErr:       nil,
				gotErr:        nil,
				compareErrors: false,
			},
			wantErr:         nil,
			wasPassedErrors: false,
		},
		{
			name: "2. Одна из ошибок пустая, а другая нет (пришла ошибка, но не должна была)",
			args: args{
				wantErr:       nil,
				gotErr:        errors.BadRequest.New(""),
				compareErrors: false,
			},
			wantErr:         errGotErrorIsNotEmpty,
			wasPassedErrors: true,
		},
		{
			name: "3. Одна из ошибок пустая, а другая нет (не пришла ошибка, но должна была)",
			args: args{
				wantErr:       errors.BadRequest.New(""),
				gotErr:        nil,
				compareErrors: false,
			},
			wantErr:         errWantErrorIsEmpty,
			wasPassedErrors: true,
		},
		{
			name: "4. Пришедшая целевая ошибка не обернута",
			args: args{
				wantErr:       errors.New(""),
				gotErr:        errors.BadRequest.New(""),
				compareErrors: false,
			},
			wantErr:         errWantErrorIsNotWrapped,
			wasPassedErrors: true,
		},
		{
			name: "5. Пришедшая полученная ошибка не обернута",
			args: args{
				wantErr:       errors.BadRequest.New(""),
				gotErr:        errors.New(""),
				compareErrors: false,
			},
			wantErr:         errGotErrorIsNotWrapped,
			wasPassedErrors: true,
		},
		{
			name: "6. Соответствие кастомных типов ошибок",
			args: args{
				wantErr:       errors.BadRequest.New(""),
				gotErr:        errors.BadRequest.New(""),
				compareErrors: false,
			},
			wantErr:         nil,
			wasPassedErrors: true,
		},
		{
			name: "7. Несоответствие кастомных типов ошибок",
			args: args{
				wantErr:       errors.BadRequest.New(""),
				gotErr:        errors.InternalServer.New(""),
				compareErrors: false,
			},
			wantErr:         errCustomTypesMismatch,
			wasPassedErrors: true,
		},
		{
			name: "8. Соответствие изначальных типов ошибок",
			args: args{
				wantErr:       errors.BadRequest.Wrap(firstType),
				gotErr:        errors.BadRequest.Wrap(firstType),
				compareErrors: true,
			},
			wantErr:         nil,
			wasPassedErrors: true,
		},
		{
			name: "9. Несоответствие изначальных типов ошибок",
			args: args{
				wantErr:       errors.BadRequest.Wrap(firstType),
				gotErr:        errors.BadRequest.Wrap(secondType),
				compareErrors: true,
			},
			wantErr:         errWrappedTypesMismatch,
			wasPassedErrors: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr, gotWasPassedErrors := checkError(tt.args.wantErr, tt.args.gotErr, tt.args.compareErrors)
			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("checkError() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
			if gotWasPassedErrors != tt.wasPassedErrors {
				t.Errorf("checkError() gotWasPassedErrors = %v, sasPassedErrors %v", gotWasPassedErrors, tt.wasPassedErrors)
			}
		})
	}
}
