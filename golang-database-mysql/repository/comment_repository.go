package repository

import (
	"context"
	"golang-database-mysql/entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, coment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
