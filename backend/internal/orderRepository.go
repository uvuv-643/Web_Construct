package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type OrderRepository interface {
	GetAll(ctx context.Context, user string) ([]Order, error)
	GetOne(ctx context.Context, uuid uuid.UUID, user string) (*Order, error)
	Create(ctx context.Context, user string, request string) (*Order, error)
	Modify(ctx context.Context, uuid uuid.UUID, content string) error
	ModifyByUser(ctx context.Context, uuid uuid.UUID, content string) error
	Delete(ctx context.Context, order *Order) error
}

type OrderRepositoryImpl struct {
	db *pg.DB
}

func (r OrderRepositoryImpl) GetAll(ctx context.Context, user string) ([]Order, error) {
	var orders []Order
	fmt.Println(user)
	err := r.db.ModelContext(ctx, &orders).
		Where("user_jwt = ?", user).
		Select()
	if err != nil {
		return nil, fmt.Errorf("failed to get all orders for user %s: %w", user, err)
	}

	return orders, nil
}

func (r OrderRepositoryImpl) GetOne(ctx context.Context, uuid uuid.UUID, user string) (*Order, error) {
	order := new(Order)
	err := r.db.ModelContext(ctx, order).
		Where("id = ?", uuid).
		Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, fmt.Errorf("order not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get order: %w", err)
	}

	// Check ownership
	if order.UserJwt != user {
		return nil, fmt.Errorf("user not authorized to access this order: %s", user)
	}

	return order, nil
}

func NewOrderRepository(db *pg.DB) OrderRepository {
	return OrderRepositoryImpl{db: db}
}

func (r OrderRepositoryImpl) Create(ctx context.Context, user string, request string) (*Order, error) {

	order := &Order{
		ID:        uuid.New(),
		Request:   request,
		UserJwt:   user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := r.db.ModelContext(ctx, order).Insert(); err != nil {
		return nil, fmt.Errorf("database_error")
	}

	return order, nil
}

func (r OrderRepositoryImpl) Modify(ctx context.Context, uuid uuid.UUID, content string) error {
	_, err := r.db.Model(&Order{}).
		Set("response = ?", content).
		Where("id = ?", uuid).
		Update()
	if err != nil {
		return err
	}
	fmt.Println(uuid, content)
	return nil
}

func (r OrderRepositoryImpl) ModifyByUser(ctx context.Context, uuid uuid.UUID, content string) error {
	_, err := r.db.Model(&Order{}).
		Set("modified_response = ?", content).
		Where("id = ?", uuid).
		Update()
	if err != nil {
		return err
	}
	fmt.Println(uuid, content)
	return nil
}

func (r OrderRepositoryImpl) Delete(ctx context.Context, order *Order) error {
	_, err := r.db.ModelContext(ctx, order).Delete()
	if err != nil {
		return err
	}
	return nil
}

func InitializeDatabase(db *pg.DB) {
	models := []interface{}{
		(*Order)(nil),
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
