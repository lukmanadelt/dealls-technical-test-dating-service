package util_test

import (
	"testing"

	"dealls-technical-test-dating-service/pkg/util"
)

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Failed: Password too long",
			args: args{
				password: "CbcDRihwZ3EICkMw0FejKb9qyICcK6T0CbcDRihwZ3EICkMw0FejKb9qyICcK6T0123456789",
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: args{
				password: "password",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := util.HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
		})
	}
}

func TestIsValidPasswordHash(t *testing.T) {
	type args struct {
		password       string
		hashedPassword string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Failed: Invalid password hash",
			args: args{
				password:       "password",
				hashedPassword: "password",
			},
			want: false,
		},
		{
			name: "Success",
			args: args{
				password:       "password",
				hashedPassword: "$2a$10$JuRPQgR8fD07tm8z0frLoOKI0TyptJCjEG0R.xIAJmwwxjsBUzaqW",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.IsValidPasswordHash(tt.args.password, tt.args.hashedPassword); got != tt.want {
				t.Errorf("IsValidPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
