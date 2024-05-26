package util_test

import (
	"testing"

	"dealls-technical-test-dating-service/pkg/util"
)

func TestGetFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed: Target is a directory",
			args: args{
				filename: "tests",
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: args{
				filename: ".env",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := util.GetFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFile() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
		})
	}
}
