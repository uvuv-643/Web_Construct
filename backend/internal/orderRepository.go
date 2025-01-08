package internal

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type OrderRepository interface {
	Create(ctx context.Context, user string, request string) (*Order, error)
	Modify(ctx context.Context, uuid uuid.UUID, content string) (*Order, error)
	Delete(ctx context.Context, order *Order) error
}

type OrderRepositoryImpl struct {
	db *pg.DB
}

func NewOrderRepository(db *pg.DB) OrderRepository {
	return OrderRepositoryImpl{db: db}
}

func (r OrderRepositoryImpl) Create(ctx context.Context, user string, request string) (*Order, error) {

	order := &Order{
		ID:        uuid.New(),
		Request:   request,
		User:      user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := r.db.ModelContext(ctx, order).Insert(); err != nil {
		return nil, fmt.Errorf("database_error")
	}

	return order, nil
}

func (r OrderRepositoryImpl) Modify(ctx context.Context, uuid uuid.UUID, content string) (*Order, error) {
	_, err := r.db.Model(&Order{}).Where("uuid = ?", uuid).Update("response", content)
	if err != nil {
		return nil, err
	}
	return nil, nil
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
