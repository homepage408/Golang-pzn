package repository

import "unit-testing-golang/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
