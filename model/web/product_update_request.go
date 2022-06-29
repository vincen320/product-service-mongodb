package web

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductUpdateRequest struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	IdUser     primitive.ObjectID `bson:"id_user,omitempty" json:"id_user,omitempty"`
	NamaProduk string             `validate:"omitempty,min=6,max=20" bson:"nama_produk,omitempty" json:"nama_produk,omitempty"`
	Harga      *int               `validate:"omitempty,numeric,gte=0" bson:"harga,omitempty" json:"harga,omitempty"`
	Kategori   string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
	Deskripsi  *string            `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Stok       *int               `validate:"omitempty,numeric,gte=0" bson:"stok,omitempty" json:"stok,omitempty"`
}

//validate:"omiempty"
//ketika ada tag validate maka akan bersifat wajib, maka itu harus ditaruh omiempty agar field tersebuh boleh kosong sehingga validasi tidak berjalan
//dan peletakkannya harus didepan sekali, jika seperti validate:"min=6,omitempty" > maka ini validasi akan tetap berjalan
/**
Allows conditional validation, for example if a field is not set with a value (Determined by the "required" validator)
then other validation such as min or max won't run, but if a value is set validation will run.
===
Mengizinkan validasi bersyarat, misalnya jika bidang tidak disetel dengan nilai (Ditentukan oleh validator "wajib")
maka validasi lain seperti min atau maks tidak akan berjalan, tetapi jika nilai ditetapkan, validasi akan berjalan.
**/
