package web

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductCreateRequest struct {
	IdUser     primitive.ObjectID `validate:"required" bson:"id_user,omitempty" json:"id_user,omitempty"`
	NamaProduk string             `validate:"required,min=6,max=20" bson:"nama_produk,omitempty" json:"nama_produk,omitempty"`
	Harga      *int               `validate:"required,numeric,gte=0" bson:"harga,omitempty" json:"harga,omitempty"`
	Kategori   string             `validate:"required" bson:"kategori,omitempty" json:"kategori,omitempty"`
	Deskripsi  *string            `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Stok       *int               `validate:"required,numeric,gte=0" bson:"stok,omitempty" json:"stok,omitempty"`
}
