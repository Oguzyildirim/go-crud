package internal_test

import (
	"errors"
	"testing"

	"github.com/Oguzyildirim/go-crud/internal"
)

func TestUser_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   internal.User
		withErr bool
	}{
		{
			"OK",
			internal.User{
				Name:     "name",
				Lastname: "lastname",
				Username: "username",
				Country:  "UK",
			},
			false,
		},
		{
			"ERR: Name",
			internal.User{
				Lastname: "lastname",
				Username: "username",
				Country:  "UK",
			},
			true,
		},
		{
			"ERR: Lastname",
			internal.User{
				Name:     "name",
				Username: "username",
				Country:  "UK",
			},
			true,
		},
		{
			"ERR: Username",
			internal.User{
				Name:     "name",
				Lastname: "lastname",
				Country:  "UK",
			},
			true,
		},
		{
			"ERR: Country",
			internal.User{
				Name:     "name",
				Lastname: "lastname",
				Username: "username",
			},
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actualErr := tt.input.Validate()
			if (actualErr != nil) != tt.withErr {
				t.Fatalf("expected error %t, got %s", tt.withErr, actualErr)
			}

			var ierr *internal.Error
			if tt.withErr && !errors.As(actualErr, &ierr) {
				t.Fatalf("expected %T error, got %T", ierr, actualErr)
			}
		})
	}
}
