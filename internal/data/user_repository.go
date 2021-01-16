package data

import (
	"context"

	"github.com/jrmanes/k8s-api-go/pkg/user"
)

// UserRepository will be a bridge to data which will give us access to DBs
type UserRepository struct {
	Data *Data
}

// GetAll implement a user repository against infrastructure
func (ur *UserRepository) GetAll(ctx context.Context) ([]user.User, error) {
	q := `
	SELECT id, user_name, role
	FROM users;
	`
	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		rows.Scan(&u.ID, &u.UserName, &u.Role)
		users = append(users, u)
	}

	return users, nil
}

// GetUser implement a user repository method against infrastructure
func (ur *UserRepository) GetUser(ctx context.Context, id uint) (user.User, error) {
	q := `
	SELECT id, user_name, role
	FROM users 
	WHERE id = $1;
	`

	row := ur.Data.DB.QueryRowContext(ctx, q, id)

	var u user.User
	err := row.Scan(&u.ID, &u.UserName, &u.Role)
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}

// Create implement a user repository method against infrastructure
func (ur *UserRepository) Create(ctx context.Context, u *user.User) error {
	q := `
	INSERT INTO users (user_name, role)
	VALUES ($1, $2)
	RETURNING id;
	`
	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, u.UserName, u.Role)

	err = row.Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete implement a user repository method against infrastructure
func (ur *UserRepository) Delete(ctx context.Context, id uint) error {
	q := `
	DELETE FROM users WHERE id=$1;
	`

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
