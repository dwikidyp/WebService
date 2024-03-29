package monitoring

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	NPM      int             `bson:"npm,omitempty" json:"npm,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
}

type OrangTua struct {
	Nama_OT    string      `bson:"nama_ot,omitempty" json:"nama_ot,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Anak 	Mahasiswa 	`bson:"anak,omitempty" json:"anak,omitempty"`
}

type Tema struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Tema    string            `bson:"nama_tema,omitempty" json:"nama_tema,omitempty"`
	Tanggal     string            `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
	Hari	string  	`bson:"hari,omitempty" json:"hari,omitempty"`
}

