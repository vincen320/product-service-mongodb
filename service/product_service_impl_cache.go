package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/vincen320/product-service-mongodb/cache"
	"github.com/vincen320/product-service-mongodb/model/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ps *ProductServiceImpl) FindByIdCache(ctx context.Context, idProduct primitive.ObjectID) (web.ProductResponse, error) {
	key := idProduct.String()
	jsonString, errCache := cache.FindKey(ctx, ps.RDB, key)

	if errCache != nil {
		//redis.Nil berarti key tidak ada == product belum di cache
		//Get manual, kemudian set
		if errCache == redis.Nil {
			//Simulation Slow Response
			time.Sleep(5 * time.Second)
			//find by id
			result, err := ps.FindById(ctx, idProduct)
			if err != nil {
				return result, err //404 not found
			}
			//set product cache
			cache.SetProductCache(ctx, ps.RDB, key, result)
			//selesai
			return result, nil
		} else {
			//error Cache yang lain
			return web.ProductResponse{}, errCache //500 internal server error
		}
	}
	//berarti product ada/sudah di cache
	var result web.ProductResponse
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return result, err //500 internal server error
	}
	result.NamaProduk = result.NamaProduk + "(cached)"
	return result, nil
}

func (ps *ProductServiceImpl) FindAllCache(ctx context.Context) (web.ProductResponses, error) {
	key := "allproducts"
	jsonString, errCache := cache.FindKey(ctx, ps.RDB, key)

	if errCache != nil {
		//redis.Nil berarti key tidak ada == product belum di cache
		//Get manual, kemudian set
		if errCache == redis.Nil {
			//Simulation Slow Response
			time.Sleep(5 * time.Second)
			//find all
			results, err := ps.FindAll(ctx)
			if err != nil {
				return results, err //404 not found
			}
			//set product cache
			cache.SetProductCache(ctx, ps.RDB, key, results)
			//selesai
			return results, nil
		} else {
			//error Cache yang lain
			return []web.ProductResponse{}, errCache //500 internal server error
		}
	}
	//berarti product ada/sudah di cache
	var results []web.ProductResponse
	err := json.Unmarshal([]byte(jsonString), &results)
	if err != nil {
		return results, err //500 internal server error
	}
	results[0].NamaProduk = results[0].NamaProduk + "(cached)"
	return results, nil
}
