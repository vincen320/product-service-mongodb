package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	IdUser       primitive.ObjectID `bson:"id_user,omitempty" json:"id_user,omitempty"`
	NamaProduk   string             `bson:"nama_produk,omitempty" json:"nama_produk,omitempty"`
	Harga        *int               `bson:"harga,omitempty" json:"harga,omitempty"`
	Kategori     string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
	Deskripsi    *string            `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Stok         *int               `bson:"stok,omitempty" json:"stok,omitempty"`
	LastModified int64              `bson:"last_modified,omitempty" json:"last_modified,omitempty"`
}
