package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restfull-api/helper"
	"golang-restfull-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	Sql := "insert into category(name) values(?)"
	resutl, err := tx.ExecContext(ctx, Sql, category.Name)
	helper.PanicIfError(err)

	id, err := resutl.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	Sql := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, Sql, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	Sql := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, Sql, category.Id)
	helper.PanicIfError(err)

}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	Sql := "select * from category where id = ?"
	rows, err := tx.QueryContext(ctx, Sql, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	var deletedAt *sql.NullString
	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt, &deletedAt)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	Sql := "select * from category"
	rows, err := tx.QueryContext(ctx, Sql)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	var deletedAt sql.NullString

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt, &deletedAt)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
