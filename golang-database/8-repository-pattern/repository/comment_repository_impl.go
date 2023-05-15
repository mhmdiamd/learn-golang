package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang-database/8-repository-pattern/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT INTO comments VALUES(DEFAULT, $1, $2)"
	rows, err := repository.DB.QueryContext(ctx, query, comment.Email, comment.Comment)

	if err != nil {
		return comment, err
	}

	// Defer berguna jika smua kode di dalam scope dieksekusi maka dia akan di eksekusi
	defer rows.Close()

	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, commentId int32) (entity.Comment, error) {
	query := "SELECT FROM comments WHERE id = $1"
	rows, err := repository.DB.QueryContext(ctx, query, commentId)

	comment := entity.Comment{}

	if err != nil {
		return comment, err
	}

	defer rows.Close()

	fmt.Println(rows)

	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	}

	return comment, errors.New("comment not found")
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	query := "SELECT * FROM comments"
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	var comments []entity.Comment
	for rows.Next() {
		var comment entity.Comment
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil

}
