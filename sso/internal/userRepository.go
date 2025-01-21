package internal

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(ctx context.Context, email Email, password string) (*User, error)
	GetByEmailAndPassword(ctx context.Context, email, password string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type UserRepositoryImpl struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, email Email, password string) (*User, error) {

	existingUser, err := r.GetByEmailAndPassword(ctx, string(email), password)
	if existingUser != nil {
		return nil, fmt.Errorf("duplicate_email")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("password_error")
	}

	user := &User{
		ID:              uuid.New(),
		Email:           email,
		PasswordBCrypto: string(hashedPassword),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if _, err := r.db.ModelContext(ctx, user).Insert(); err != nil {
		return nil, fmt.Errorf("database_error")
	}

	role := &UserRole{
		ID:            uuid.New(),
		UserID:        user.ID,
		Role:          "PT_READ",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		ApplicationID: "backend",
	}
	if _, err := r.db.ModelContext(ctx, role).Insert(); err != nil {
		return nil, fmt.Errorf("role_error")
	}

	role = &UserRole{
		ID:            uuid.New(),
		UserID:        user.ID,
		Role:          "PT_SHARE",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		ApplicationID: "proxy",
	}
	if _, err := r.db.ModelContext(ctx, role).Insert(); err != nil {
		return nil, fmt.Errorf("role_error")
	}

	role = &UserRole{
		ID:            uuid.New(),
		UserID:        user.ID,
		Role:          "PT_SHARE",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		ApplicationID: "backend",
	}
	if _, err := r.db.ModelContext(ctx, role).Insert(); err != nil {
		return nil, fmt.Errorf("role_error")
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetByEmailAndPassword(ctx context.Context, email, password string) (*User, error) {
	user := new(User)
	err := r.db.ModelContext(ctx, user).
		Relation("Roles").
		Where("email = ?", email).
		Select()
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordBCrypto), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetByEmail(ctx context.Context, email string) (*User, error) {
	user := new(User)
	err := r.db.ModelContext(ctx, user).
		Relation("Roles").
		Where("email = ?", email).
		Select()
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return user, nil
}

func InitializeDatabase(db *pg.DB) {
	models := []interface{}{
		(*User)(nil),
		(*UserRole)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			log.Fatalf("Failed to create table for %T: %v", model, err)
		}
	}
	log.Println("Database schema initialized successfully.")
}
