package repository

import (
	"context"
	"fmt"
	golang_database_mysql "golang-database-mysql"
	"golang-database-mysql/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database_mysql.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "teguhRepo@test.com",
		Comment: "Test Repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database_mysql.GetConnection())

	ctx := context.Background()
	coment, err := commentRepository.FindById(ctx, 12)
	if err != nil {
		panic(err)
	}

	fmt.Println(coment)

}
