package models

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id    bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string        `json:"name,omitempty"`
	Email string        `json:"email,omitempty"`
	Date  time.Time     `json:"date, omitempty"`
}

type UserRepo struct {
	C *mgo.Collection
}

func (u *UserRepo) GetAll() []User {
	users := []User{}
	iter := u.C.Find(nil).Iter()
	user := User{}
	for iter.Next(&user) {
		users = append(users, user)
	}
	return users
}

func (u *UserRepo) Create(user *User) error {
	user.Id = bson.NewObjectId()
	user.Date = time.Now()
	err := u.C.Insert(&user)
	return err
}

func (u *UserRepo) Delete(id string) error {
	err := u.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
