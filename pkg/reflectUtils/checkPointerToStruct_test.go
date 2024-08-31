package reflectUtils

import (
	"testing"

	"server/pkg/errors"
	"server/pkg/testUtils"
)

func TestCheckPointerToStruct(t *testing.T) {
	type args struct {
		dest any
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "1. Передача в функцию указателя на структуру",
			args: args{
				dest: &struct{}{},
			},
		},
		{
			name: "2. Передача в функцию указателя на строку",
			args: args{
				dest: new(string),
			},
			wantErr: errors.InternalServer.New(""),
		},
		{
			name: "3. Передача в функцию указателя на число",
			args: args{
				dest: new(int),
			},
			wantErr: errors.InternalServer.New(""),
		},
		{
			name: "4. Передача в функцию копии структуры",
			args: args{
				dest: struct{}{},
			},
			wantErr: errors.InternalServer.New(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := CheckPointerToStruct(tt.args.dest)
			testUtils.CheckError(t, tt.wantErr, gotErr, false)
		})
	}
}
