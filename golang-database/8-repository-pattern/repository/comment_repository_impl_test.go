package repository_test

import (
	"context"
	"fmt"
	"golang-database/8-repository-pattern/entity"
	"golang-database/8-repository-pattern/repository"
	"golang-database/database"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	repository := repository.NewCommentRepository(database.GetConnection())

	context := context.Background()
	content := entity.Comment{
		Email:   "am@gmail.com",
		Comment: "Hallo Ilham",
	}

	_, err := repository.Insert(context, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert new Data comment")
}

func TestCommentFindById(t *testing.T) {

	// Assign your Repository
	repository := repository.NewCommentRepository(database.GetConnection())
	// Create Context
	ctx := context.Background()

	// Execute the repository
	result, err := repository.FindById(ctx, 8)

	if err != nil {
		panic(err)
	}

	fmt.Println("email : ", result.Email)
}

func TestCommentFindAll(t *testing.T) {
	// Assign your Repository
	repository := repository.NewCommentRepository(database.GetConnection())
	// Create Context
	ctx := context.Background()

	result, err := repository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for comment := range result {
		fmt.Println(comment)
	}

}
