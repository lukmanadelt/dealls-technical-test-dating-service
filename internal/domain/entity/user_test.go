package entity_test

import (
	"testing"

	"dealls-technical-test-dating-service/internal/domain/entity"
)

func TestUserSignupRequest_Validate(t *testing.T) {
	type fields struct {
		email     string
		password  string
		name      string
		birthDate string
		gender    string
		location  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Failed: Empty email",
			fields: fields{
				email: "",
			},
			wantErr: true,
		},
		{
			name: "Failed: Invalid email format",
			fields: fields{
				email: "invalid@email",
			},
			wantErr: true,
		},
		{
			name: "Failed: Empty password",
			fields: fields{
				email:    "valid@email.com",
				password: "",
			},
			wantErr: true,
		},
		{
			name: "Failed: Invalid password length",
			fields: fields{
				email:    "valid@email.com",
				password: "pass",
			},
			wantErr: true,
		},
		{
			name: "Failed: Empty name",
			fields: fields{
				email:    "valid@email.com",
				password: "password",
				name:     "",
			},
			wantErr: true,
		},
		{
			name: "Failed: Empty birth date",
			fields: fields{
				email:     "valid@email.com",
				password:  "password",
				name:      "name",
				birthDate: "",
			},
			wantErr: true,
		},
		{
			name: "Failed: Invalid birth date format",
			fields: fields{
				email:     "valid@email.com",
				password:  "password",
				name:      "name",
				birthDate: "01-01-2024",
			},
			wantErr: true,
		},
		{
			name: "Failed: Empty gender",
			fields: fields{
				email:     "valid@email.com",
				password:  "password",
				name:      "name",
				birthDate: "2024-01-01",
				gender:    "",
			},
			wantErr: true,
		},
		{
			name: "Failed: Invalid gender value",
			fields: fields{
				email:     "valid@email.com",
				password:  "password",
				name:      "name",
				birthDate: "2024-01-01",
				gender:    "INVALID",
			},
			wantErr: true,
		},
		{
			name: "Failed: Empty location",
			fields: fields{
				email:     "valid@email.com",
				password:  "password",
				name:      "name",
				birthDate: "2024-01-01",
				gender:    "MALE",
			},
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				email:     "valid@email.com",
				password:  "password",
				name:      "name",
				birthDate: "2024-01-01",
				gender:    "MALE",
				location:  "Indonesia",
			},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := &entity.UserSignupRequest{
				Email:     test.fields.email,
				Password:  test.fields.password,
				Name:      test.fields.name,
				BirthDate: test.fields.birthDate,
				Gender:    test.fields.gender,
				Location:  test.fields.location,
			}
			if err := req.Validate(); (err != nil) != test.wantErr {
				t.Errorf("UserSignupRequest.Validate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestUserLoginRequest_Validate(t *testing.T) {
	type fields struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Failed: Empty email",
			fields: fields{
				email: "",
			},
			wantErr: true,
		},
		{
			name: "Failed: Invalid email format",
			fields: fields{
				email: "invalid@email",
			},
			wantErr: true,
		},
		{
			name: "Failed: Empty password",
			fields: fields{
				email:    "valid@email.com",
				password: "",
			},
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				email:    "valid@email.com",
				password: "password",
			},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := &entity.UserLoginRequest{
				Email:    test.fields.email,
				Password: test.fields.password,
			}
			if err := u.Validate(); (err != nil) != test.wantErr {
				t.Errorf("UserLoginRequest.Validate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
