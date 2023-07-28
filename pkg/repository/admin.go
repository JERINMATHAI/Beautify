package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"beautify/pkg/domain"
	"beautify/pkg/repository/interfaces"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"

	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB: db}
}

// SaveAdmin - Inserts a new admin record into the database.

func (a *adminDatabase) SaveAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error) {
	query := `INSERT INTO admins(user_name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	createdAt := time.Now()
	updatedAt := time.Now()
	if a.DB.Exec(query, admin.UserName, admin.Email, admin.Password, createdAt, updatedAt).Error != nil {
		return admin, errors.New("Failed to save admin")
	}
	return admin, nil
}

// Get admin - Retrieves an admin record from the database based on the provided email or username.

func (a *adminDatabase) GetAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error) {
	query := `SELECT * FROM admins WHERE email =? OR user_name =?`
	if a.DB.Raw(query, admin.Email, admin.UserName).Scan(&admin).Error != nil {
		return admin, errors.New("Failed to find admin")
	}
	return admin, nil
}

func (a *adminDatabase) GetAllUser(ctx context.Context, page request.ReqPagination) (users []response.UserRespStrcut, err error) {
	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	query := `SELECT * FROM users ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	err = a.DB.Raw(query, limit, offset).Scan(&users).Error

	return users, err
}

func (a *adminDatabase) BlockUser(ctx context.Context, userID uint) error {
	// Check user if valid or not
	var user domain.Users
	query := `SELECT * FROM users WHERE id=?`
	a.DB.Raw(query, userID).Scan(&user)
	if user.Email == "" {
		// check user email with user ID
		return errors.New("Invalid user id user doesn't exist")
	}

	query = `UPDATE users SET block_status = $1 WHERE id = $2`
	if a.DB.Exec(query, !user.BlockStatus, userID).Error != nil {
		return fmt.Errorf("Failed to update user block_status to %v", !user.BlockStatus)
	}
	return nil
}
