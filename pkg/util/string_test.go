package util_test

import (
	"testing"

	"dealls-technical-test-dating-service/pkg/util"
)

func TestIsValidEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Failed: Invalid email format",
			args: args{
				email: "invalid@email",
			},
			want: false,
		},
		{
			name: "Success",
			args: args{
				email: "user@email.com",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.IsValidEmail(tt.args.email); got != tt.want {
				t.Errorf("IsValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
