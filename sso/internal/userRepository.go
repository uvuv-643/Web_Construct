package internal

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"strings"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(ctx context.Context, email Email, password, fullName string) error
	GetByEmailAndPassword(ctx context.Context, email, password string) (*User, error)
}

type UserRepositoryImpl struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, email Email, password, fullName string) error {
	if email == "" || password == "" || fullName == "" {
		return fmt.Errorf("all fields are required")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user := &User{
		ID:              uuid.New(),
		Email:           email,
		PasswordBCrypto: string(hashedPassword),
		FullName:        fullName,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Roles:           []*UserRole{{Role: "PT_READ", CreatedAt: time.Now(), UpdatedAt: time.Now()}},
	}

	_, err = r.db.ModelContext(ctx, user).Insert()
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *UserRepositoryImpl) GetByEmailAndPassword(ctx context.Context, email, password string) (*User, error) {
	if email == "" || password == "" {
		return nil, fmt.Errorf("email and password are required")
	}

	user := new(User)
	err := r.db.ModelContext(ctx, user).
		Where("email = ?", strings.ToLower(email)).
		Select()
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordBCrypto), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
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

func ConnectToDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "your_user",
		Password: "your_password",
		Database: "your_database",
	})
	InitializeDatabase(db)
	return db
}
