package userbiz_test

import (
	"context"
	"testing"
	"time"

	"github.com/oussamm/booksales/business/domain/userbiz"
	"github.com/oussamm/booksales/business/sys/validate"
)

func TestCreate(t *testing.T) {
	now := time.Now()
	userTests := []struct {
		NewUser userbiz.NewUser
		User    userbiz.User
	}{
		{
			NewUser: userbiz.NewUser{
				Name:            "Ouss",
				Email:           "ouss@gmail.com",
				Roles:           []string{"admin"},
				Password:        "test",
				PasswordConfirm: "test",
			},
			User: userbiz.User{
				Name:         "Ouss",
				Email:        "ouss@gmail.com",
				ID:           validate.GenerateID(),
				Roles:        []string{"admin"},
				PasswordHash: nil,
				DateCreated:  now,
				DateUpdated:  now,
			},
		},
	}

	ctx := context.Background()
	b := userbiz.Bus{}

	for _, tc := range userTests {
		t.Run("test1", func(t *testing.T) {
			u, _ := b.Create(ctx, tc.NewUser, now)
			if u.Name != tc.User.Name {
				t.Error("name is not the same")
			}
		})
	}

}
