package repository

import (
	"context"
	"errors"

	"github.com/vincen320/product-service-mongodb/exception"
	"github.com/vincen320/product-service-mongodb/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepositoryImpl struct {
}

func NewProductRespository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (pr *ProductRepositoryImpl) Create(ctx context.Context, db *mongo.Database, product domain.Product) (domain.Product, error) {
	result, err := db.Collection("product").InsertOne(ctx, product)
	if err != nil {
		return product, err //500 Internal Server Error
	}
	id := result.InsertedID
	productId, ok := id.(primitive.ObjectID)
	if !ok {
		return product, errors.New("unable to retrieve ID") // 500 Internal Server Error
	}
	product.Id = productId
	return product, nil
}

//https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/query-document/
func (pr *ProductRepositoryImpl) Update(ctx context.Context, db *mongo.Database, product domain.Product) (domain.Product, error) {
	filter := bson.M{
		"_id": product.Id,
	}

	// filter := bson.D{
	// 	{Key: "$and",
	// 		Value: bson.A{
	// 			bson.D{{Key: "_id", Value: product.Id}},
	// 			bson.D{{Key: "id_user", Value: product.IdUser}},
	// 		},
	// 	},
	// }

	result, err := db.Collection("product").UpdateOne(ctx, filter, bson.M{
		"$set": product,
	})
	if err != nil {
		return product, err //500 Internal Server Error
	}

	if result.ModifiedCount == 0 {
		return product, exception.NewNotFoundErr("no product was updated") //404 Not Found
	}

	return product, nil
}

func (pr *ProductRepositoryImpl) Delete(ctx context.Context, db *mongo.Database, idProduct primitive.ObjectID) error {
	filter := bson.M{
		"_id": idProduct,
	}

	// filter := bson.D{
	// 	{Key: "$and",
	// 		Value: bson.A{
	// 			bson.D{{Key: "_id", Value: idProduct}},
	// 			bson.D{{Key: "id_user", Value: idUser}},
	// 		},
	// 	},
	// }

	result, err := db.Collection("product").DeleteOne(ctx, filter)
	if err != nil {
		return err //500 Internal Server Error
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFoundErr("no product was deleted") // 404 Not Found
	}
	return nil
}

func (pr *ProductRepositoryImpl) FindById(ctx context.Context, db *mongo.Database, idProduct primitive.ObjectID) (domain.Product, error) {
	filter := bson.M{
		"_id": idProduct,
	}

	result := db.Collection("product").FindOne(ctx, filter)
	var product domain.Product

	err := result.Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return product, exception.NewNotFoundErr("product not found") //404 not found, return error trs smpe service ttp return error
		}
		return product, err // 500 Internal Server Error
	}

	return product, nil
}

func (pr *ProductRepositoryImpl) FindAll(ctx context.Context, db *mongo.Database) ([]domain.Product, error) {
	var products []domain.Product
	cursor, err := db.Collection("product").Find(ctx, bson.M{})
	if err != nil {
		return products, err // 500 Internal Server Error
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &products)
	if err != nil {
		return products, err // 500 Internal Server Error
	}
	return products, nil
}
