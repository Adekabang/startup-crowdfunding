package user

import "gorm.io/gorm"

//Repository interface
type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

//NewRepository for create user object with db connection
func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
