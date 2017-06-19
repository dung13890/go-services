package repo

import (
	"time"

	"github.com/dung13890/go-services/users/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepo struct {
	C *mgo.Collection
}

func (u *UserRepo) GetById(id string) (models.User, error) {
	user := models.User{}
	err := u.C.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)
	return user, err
}

func (u *UserRepo) GetAll() []models.User {
	users := []models.User{}
	iter := u.C.Find(nil).Iter()
	user := models.User{}
	for iter.Next(&user) {
		users = append(users, user)
	}
	return users
}

func (u *UserRepo) Create(user *models.User) error {
	user.Id = bson.NewObjectId()
	user.Date = time.Now()
	err := u.C.Insert(&user)
	return err
}

func (u *UserRepo) Delete(id string) error {
	err := u.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
