package data

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jrmanes/k8s-api-go/pkg/user"
	_ "github.com/lib/pq"
)



var u = user.User{
	ID:       1,
	UserName: "jrmanes",
	Role:     "devops",
}

func NewMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

// test get user by id -
// pending to fix mock due to: sql: unknown driver
/*
func TestUserRepository_GetUser(t *testing.T) {
	_, mock := NewMock(t)

	repo := &UserRepository{Data: New()}

	query := "SELECT id, user_name, role FROM users WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
		AddRow(u.ID, u.UserName, u.Role)

	mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows(rows)

	var ctx context.Context
	user, err := repo.GetUser(ctx ,uint(u.ID))
	assert.NotNil(t, user)
	assert.NoError(t, err)
}
*/