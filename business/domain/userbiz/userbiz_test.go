package userbiz_test

import (
	"testing"
	"time"

	"ousloob.io/booksales/business/domain/userbiz"
	"ousloob.io/booksales/business/sys/validate"
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

	for _, tc := range userTests {
		t.Run("test1", func(t *testing.T) {
			u := userbiz.User{
				Name: "Ouss",
			}
			if u.Name != tc.User.Name {
				t.Error("name is not the same")
			}
		})
	}

}
