package auth_test

import (
	"testing"
	"time"

	"dealls-technical-test-dating-service/internal/infrastructure/auth"
)

func TestJWTClaims_GenerateToken(t *testing.T) {
	type args struct {
		email string
		key   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				email: "user@email.com",
				key:   "key",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := auth.NewJWTClaims(24 * time.Hour)
			_, err := j.GenerateToken(tt.args.email, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWTClaims.GenerateToken() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
		})
	}
}
