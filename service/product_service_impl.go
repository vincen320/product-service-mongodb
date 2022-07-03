package service

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
	"github.com/vincen320/product-service-mongodb/exception"
	"github.com/vincen320/product-service-mongodb/model/domain"
	"github.com/vincen320/product-service-mongodb/model/web"
	"github.com/vincen320/product-service-mongodb/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductServiceImpl struct {
	Repository repository.ProductRepository
	Validator  *validator.Validate
	DB         *mongo.Database
	RDB        *redis.Client
}

func NewProductService(repository repository.ProductRepository, validator *validator.Validate, db *mongo.Database, rdb *redis.Client) ProductService {
	return &ProductServiceImpl{
		Repository: repository,
		Validator:  validator,
		DB:         db,
		RDB:        rdb,
	}
}

func (ps *ProductServiceImpl) Create(ctx context.Context, createRequest web.ProductCreateRequest) (web.ProductResponse, error) {
	var response web.ProductResponse
	var savedProduct domain.Product
	err := ps.Validator.Struct(createRequest)
	if err != nil {
		return response, err // 401 bad request
	}

	session, err := ps.DB.Client().StartSession()
	if err != nil {
		return response, err // 500 internal server error
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err = session.StartTransaction()
		if err != nil {
			return err // 500 Internal server error
		}

		timeNow := time.Now().UTC().UnixMilli()
		savedProduct, err = ps.Repository.Create(ctx, ps.DB, domain.Product{
			IdUser:       createRequest.IdUser,
			NamaProduk:   createRequest.NamaProduk,
			Harga:        createRequest.Harga,
			Kategori:     createRequest.Kategori,
			Deskripsi:    createRequest.Deskripsi,
			Stok:         createRequest.Stok,
			LastModified: timeNow,
		})

		if err != nil {
			return err //500 Internal Server Error
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err //500 Internal Server Error
		}
		return nil
	})

	if err != nil {
		errAbort := session.AbortTransaction(ctx)
		if errAbort != nil {
			return response, errAbort // Internal Server Error
		}
		return response, err // Err yang terdapat direturn di fungsi diatas (mongo.WithSession)
	}

	response.Id = savedProduct.Id
	response.IdUser = savedProduct.IdUser
	response.NamaProduk = savedProduct.NamaProduk
	response.Harga = savedProduct.Harga
	response.Kategori = savedProduct.Kategori
	response.Stok = savedProduct.Stok
	response.Deskripsi = savedProduct.Deskripsi
	response.LastModified = savedProduct.LastModified
	return response, nil
}

func (ps *ProductServiceImpl) Update(ctx context.Context, updateRequest web.ProductUpdateRequest) (web.ProductResponse, error) {
	var response web.ProductResponse
	var updatedProduct domain.Product
	err := ps.Validator.Struct(updateRequest)
	if err != nil {
		return response, err // 401 bad request
	}

	//Find Product is exist and Validate User
	product, err := ps.Repository.FindById(ctx, ps.DB, updateRequest.Id)
	if err != nil {
		return response, err //500 atau 404
	}

	if product.IdUser != updateRequest.IdUser {
		return response, exception.NewUnauthorizedErr("can't update product that not yours")
	}

	session, err := ps.DB.Client().StartSession()
	if err != nil {
		return response, err // 500 internal server error
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err = session.StartTransaction()
		if err != nil {
			return err // 500 Internal server error
		}

		timeNow := time.Now().UTC().UnixMilli()
		updatedProduct, err = ps.Repository.Update(ctx, ps.DB, domain.Product{
			Id:           updateRequest.Id,
			IdUser:       updateRequest.IdUser,
			NamaProduk:   updateRequest.NamaProduk,
			Harga:        updateRequest.Harga,
			Kategori:     updateRequest.Kategori,
			Deskripsi:    updateRequest.Deskripsi,
			Stok:         updateRequest.Stok,
			LastModified: timeNow,
		})

		if err != nil {
			return err ////500 atau 404
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err //500 Internal Server Error
		}
		return nil
	})

	if err != nil {
		errAbort := session.AbortTransaction(ctx)
		if errAbort != nil {
			return response, errAbort // Internal Server Error
		}
		return response, err // Err yang terdapat direturn di fungsi diatas (mongo.WithSession)
	}
	copier.CopyWithOption(&product, &updatedProduct, copier.Option{
		IgnoreEmpty: true,
	})

	response.Id = product.Id
	response.IdUser = product.IdUser
	response.NamaProduk = product.NamaProduk
	response.Harga = product.Harga
	response.Kategori = product.Kategori
	response.Stok = product.Stok
	response.Deskripsi = product.Deskripsi
	response.LastModified = product.LastModified
	return response, nil
}

func (ps *ProductServiceImpl) Delete(ctx context.Context, idProduct primitive.ObjectID, idUser primitive.ObjectID) error {
	//Find Product is exist and Validate User
	product, err := ps.Repository.FindById(ctx, ps.DB, idProduct)
	if err != nil {
		return err //500 atau 404
	}

	if product.IdUser != idUser {
		return exception.NewUnauthorizedErr("can't delete product that not yours")
	}

	session, err := ps.DB.Client().StartSession()
	if err != nil {
		return err // 500 internal server error
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err = session.StartTransaction()
		if err != nil {
			return err // 500 Internal server error
		}

		err = ps.Repository.Delete(ctx, ps.DB, idProduct)

		if err != nil {
			return err ////500 atau 404
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err //500 Internal Server Error
		}
		return nil
	})

	if err != nil {
		errAbort := session.AbortTransaction(ctx)
		if errAbort != nil {
			return errAbort // Internal Server Error
		}
		return err // Err yang terdapat direturn di fungsi diatas (mongo.WithSession)
	}

	return nil
}

func (ps *ProductServiceImpl) FindById(ctx context.Context, idProduct primitive.ObjectID) (web.ProductResponse, error) {
	var response web.ProductResponse
	var foundProduct domain.Product

	session, err := ps.DB.Client().StartSession()
	if err != nil {
		return response, err // 500 internal server error
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err = session.StartTransaction()
		if err != nil {
			return err // 500 Internal server error
		}

		foundProduct, err = ps.Repository.FindById(ctx, ps.DB, idProduct)

		if err != nil {
			return err ////500 atau 404
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err //500 Internal Server Error
		}
		return nil
	})

	if err != nil {
		errAbort := session.AbortTransaction(ctx)
		if errAbort != nil {
			return response, errAbort // Internal Server Error
		}
		return response, err // Err yang terdapat direturn di fungsi diatas (mongo.WithSession)
	}
	response.Id = foundProduct.Id
	response.IdUser = foundProduct.IdUser
	response.NamaProduk = foundProduct.NamaProduk
	response.Harga = foundProduct.Harga
	response.Kategori = foundProduct.Kategori
	response.Stok = foundProduct.Stok
	response.Deskripsi = foundProduct.Deskripsi
	response.LastModified = foundProduct.LastModified
	return response, nil
}

func (ps *ProductServiceImpl) FindAll(ctx context.Context) (web.ProductResponses, error) {
	var responses []web.ProductResponse
	var foundProducts []domain.Product

	session, err := ps.DB.Client().StartSession()
	if err != nil {
		return responses, err // 500 internal server error
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err = session.StartTransaction()
		if err != nil {
			return err // 500 Internal server error
		}

		foundProducts, err = ps.Repository.FindAll(ctx, ps.DB)

		if err != nil {
			return err ////500 atau 404
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err //500 Internal Server Error
		}
		return nil
	})

	if err != nil {
		errAbort := session.AbortTransaction(ctx)
		if errAbort != nil {
			return responses, errAbort // Internal Server Error
		}
		return responses, err // Err yang terdapat direturn di fungsi diatas (mongo.WithSession)
	}

	for _, foundProduct := range foundProducts {
		responses = append(responses, web.ProductResponse{
			Id:           foundProduct.Id,
			IdUser:       foundProduct.IdUser,
			NamaProduk:   foundProduct.NamaProduk,
			Harga:        foundProduct.Harga,
			Kategori:     foundProduct.Kategori,
			Stok:         foundProduct.Stok,
			Deskripsi:    foundProduct.Deskripsi,
			LastModified: foundProduct.LastModified,
		})
	}
	return responses, nil
}
