package repository

import (
	"context"

	"github.com/vincen320/product-service-mongodb/model/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	Create(ctx context.Context, db *mongo.Database, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, db *mongo.Database, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, db *mongo.Database, idProduct primitive.ObjectID) error
	FindById(ctx context.Context, db *mongo.Database, idProduct primitive.ObjectID) (domain.Product, error)
	FindAll(ctx context.Context, db *mongo.Database) ([]domain.Product, error)
}
