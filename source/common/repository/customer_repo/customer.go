package customerrepo

import (
	"context"

	customermodel "github.com/i-sub135/go-rest-blueprint/source/common/model/customer_model"
	"gorm.io/gorm"
)

type CustomerRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{db: db}
}

func (cs *CustomerRepo) GetCustomerFirstName(ctx context.Context, name string) (*[]customermodel.Customer, error) {
	var customers []customermodel.Customer
	err := cs.db.WithContext(ctx).
		Where("first_name = ?", name).
		Find(&customers).Error

	if err != nil {
		return nil, err
	}
	return &customers, nil
}
