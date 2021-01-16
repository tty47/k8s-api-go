package user

import "testing"

func Test_User(t *testing.T) {
	user := User{
		ID:       0,
		UserName: "jrmanes",
		Role:     "devops",
	}

	t.Run("check user id", func(t *testing.T) {
		got := user.ID
		var want int32 = 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("validate username", func(t *testing.T) {
		got := user.UserName
		want := "jrmanes"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("validate user role", func(t *testing.T) {
		got := user.Role
		want := "devops"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
