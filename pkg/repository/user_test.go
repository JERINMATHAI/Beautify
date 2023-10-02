package repository

import (
	"beautify/pkg/domain"
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gopkg.in/go-playground/assert.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestFindUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("SQL mock failed %v ", err)
	}
	defer db.Close()
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to create psql instance ")
	}
	err = gormDb.Statement.Error
	if err != nil {
		t.Fatalf("Failed to ping the database: %v", err)
	}
	user := domain.Users{
		ID:        2,
		UserName:  "JerinMathai",
		FirstName: "Jerin",
		LastName:  "Mathai",
		Age:       25,
		Email:     "jerin@gmail.com",
	}
	expectedError := errors.New("Failed to find user")
	authDb := &userDatabase{DB: gormDb}
	mock.ExpectQuery(
		`SELECT * FROM users where id=? OR user_name=? OR email=? OR phone=?`).
		WithArgs(user.ID, user.Email, user.Phone, user.UserName).
		WillReturnRows(sqlmock.NewRows([]string{"ID", "UserName", "FirstName", "LastName", "Age", "Email"}))
	result, resultErr := authDb.FindUser(context.Background(), user)
	assert.Equal(t, expectedError, resultErr)
	fmt.Println(result)

}
