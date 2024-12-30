package internal

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordBCrypto), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.PasswordBCrypto = string(hashedPassword)

	_, err = r.db.ModelContext(ctx, user).Insert()
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *User) error {
	_, err := r.db.ModelContext(ctx, user).
		WherePK().
		Update()
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *UserRepositoryImpl) Get(ctx context.Context, opts *FindUserOptions) (*User, error) {
	user := new(User)
	query := r.db.ModelContext(ctx, user)

	if opts.Email != "" {
		query.Where("email = ?", opts.Email)
	}
	if opts.ID != uuid.Nil {
		query.Where("id = ?", opts.ID)
	}

	err := query.Select()
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func ConnectToDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "your_username",
		Password: "your_password",
		Database: "your_database",
	})
	return db
}
