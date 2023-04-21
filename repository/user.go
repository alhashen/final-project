package repository

import (
	"finalproj/model"
)

type UserRepo interface {
	RegisterUser(model.User) (model.User, error)
	LogInUser(map[string]any) error
}

func (r Repo) RegisterUser(user model.User) (res model.User, err error) {

	err = r.gorm.Debug().Create(&user).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) LogInUser(user map[string]any) error {

	userm := model.User{}

	qry := r.gorm.Where("email = ?", user["user_email"]).Take(&userm)
	err := qry.Error

	if err != nil {
		return err
	}

	user["user_id"] = userm.ID
	user["user_password"] = userm.Password

	return nil
}
