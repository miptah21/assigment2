package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"errors"

	"go.etcd.io/bbolt"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userRepository struct {
	filebasedDb *filebased.Data
}

func NewUserRepo(filebasedDb *filebased.Data) *userRepository {
	return &userRepository{filebasedDb}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	user, err := r.filebasedDb.GetUserByEmail(email)
	if err != nil {
		return model.User{}, err
	}

	var User model.User{}
	if user.Email == User.Email {
		Users = append(User, user)
		return user, nil
	}

	return user, nil
}



func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	createdUser, err := r.filebasedDb.CreateUser(user)

	if err != nil {
		return model.User{}, err
	}

	return createdUser, nil
}

func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	var userTaskCategories []model.UserTaskCategory
	err := r.filebasedDb.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("usertaskcategories"))
		if b == nil {
			return errors.New("usertaskcategories bucket not found")
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var utc model.UserTaskCategory
			if err := json.Unmarshal(v, &utc); err != nil {
				return err
			}
			userTaskCategories = append(userTaskCategories, utc)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return userTaskCategories, nil
}
