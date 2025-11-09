package userrepo

import (
	"context"

	usermodel "github.com/i-sub135/go-rest-blueprint/source/common/model/user_model"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) GetAll(ctx context.Context) (*[]usermodel.User, error) {
	var user []usermodel.User
	err := r.DB.WithContext(ctx).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetByID(ctx context.Context, id uint) (*usermodel.User, error) {
	var user usermodel.User
	err := r.DB.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*usermodel.User, error) {
	var user usermodel.User
	err := r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) Create(ctx context.Context, user *usermodel.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *UserRepo) Update(ctx context.Context, user *usermodel.User) error {
	return r.DB.WithContext(ctx).Save(user).Error
}

func (r *UserRepo) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&usermodel.User{}, id).Error
}
