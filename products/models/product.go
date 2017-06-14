package models

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Sku         string        `json:"sku,omitempty"`
	Price       int           `json:"price,omitempty"`
	Description string        `json:"description,omitempty"`
	Date        time.Time     `json:"date, omitempty"`
}

type ProductRepo struct {
	C *mgo.Collection
}

func (u *ProductRepo) GetAll() []Product {
	products := []Product{}
	iter := u.C.Find(nil).Iter()
	product := Product{}
	for iter.Next(&product) {
		products = append(products, product)
	}
	return products
}

func (u *ProductRepo) Create(product *Product) error {
	product.Id = bson.NewObjectId()
	product.Date = time.Now()
	err := u.C.Insert(&product)
	return err
}

func (u *ProductRepo) Delete(id string) error {
	err := u.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
