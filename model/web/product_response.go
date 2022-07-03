package web

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductResponse struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	IdUser       primitive.ObjectID `bson:"id_user,omitempty" json:"id_user,omitempty"`
	NamaProduk   string             `bson:"nama_produk,omitempty" json:"nama_produk,omitempty"`
	Harga        *int               `bson:"harga,omitempty" json:"harga,omitempty"`
	Kategori     string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
	Deskripsi    *string            `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Stok         *int               `bson:"stok,omitempty" json:"stok,omitempty"`
	LastModified int64              `bson:"last_modified,omitempty" json:"last_modified,omitempty"`
}

//impement encoding.BinaryMarshaler agar bisa di set di redis
func (pr ProductResponse) MarshalBinary() (data []byte, err error) {
	return json.Marshal(pr) // TODO: Implement
}

//===BARU DIBUAT SAAT ADA REDIS, BUAT ALIAS UNTUK ARRAY DARI PRODUCTRESPONSE AGAR BISA IMPLEMENT encoding.BinaryMarshaler
//(testing), digunakan di service yang findAll (cache dan tidak)
type ProductResponses []ProductResponse

//impement encoding.BinaryMarshaler agar bisa di set di redis
func (prs ProductResponses) MarshalBinary() (data []byte, err error) {
	return json.Marshal(prs) // TODO: Implement
}
