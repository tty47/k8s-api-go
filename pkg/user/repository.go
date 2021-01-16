package user

import "context"

// Interface repository for user struct
type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	GetUser(ctx context.Context, id uint) (User, error)
	Create(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uint) error
}
