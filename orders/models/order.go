package models

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Order struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	User     string        `json:"user,omitempty"`
	Products []string      `json:"products,omitempty"`
	Date     time.Time     `json:"date, omitempty"`
}

type OrderRepo struct {
	C *mgo.Collection
}

func (u *OrderRepo) GetUser(id string) (Order, error) {
	order := Order{}
	err := u.C.Find(bson.M{"user": id}).One(&order)
	return order, err
}

func (u *OrderRepo) Create(order *Order) error {
	order.Id = bson.NewObjectId()
	order.Date = time.Now()
	err := u.C.Insert(&order)
	return err
}

func (u *OrderRepo) Delete(id string) error {
	err := u.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
