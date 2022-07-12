package usecase

import (
	"context"
	"fmt"

	"github.com/ironsail/whydah-go-clean-template/internal/entity"
)

// UserUseCase -.
type UserUseCase struct {
	Store *entity.UserStore
}

// New -.
func NewUserUseCase(store *entity.UserStore) *UserUseCase {
	return &UserUseCase{store}
}

// List - get all users
func (uc *UserUseCase) List(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	uc.Store.Session().WithContext(ctx) // set context
	err := uc.Store.Find().All(&users)

	if err != nil {
		return nil, fmt.Errorf("Users - get list: %w", err)
	}

	return users, nil
}

// Create -.
func (uc *UserUseCase) Create(ctx context.Context, u entity.User) (entity.User, error) {

	uc.Store.Session().WithContext(ctx) // set context
	err := uc.Store.Session().Save(u.ToRecord())
	if err != nil {
		return entity.User{}, fmt.Errorf("Users - create: %w", err)
	}

	return u, nil
}
