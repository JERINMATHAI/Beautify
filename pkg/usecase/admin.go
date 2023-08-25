package usecase

import (
	"beautify/pkg/domain"
	"beautify/pkg/repository/interfaces"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"context"
	"errors"

	service "beautify/pkg/usecase/interfaces"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

// type adminService struct {
// 	adminRepository interfaces.AdminRepository
// }

// func NewAdminService(repo interfaces.AdminRepository) service.AdminService {
// 	return &adminService{adminRepository: repo}
// }

type adminService struct {
	adminRepository   interfaces.AdminRepository
	PaymentRepository interfaces.PaymentRepository
}

func NewAdminService(repo interfaces.AdminRepository, PaymentRepo interfaces.PaymentRepository) service.AdminService {
	return &adminService{adminRepository: repo,
		PaymentRepository: PaymentRepo}
}

// Signup
func (a *adminService) Signup(c context.Context, admin domain.Admin) (domain.Admin, error) {
	if dbAdmin, err := a.adminRepository.GetAdmin(c, admin); err != nil {
		return admin, err
	} else if dbAdmin.ID != 0 {
		return admin, errors.New("User already exists")
	}

	// Hash password
	PassHash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return admin, errors.New("Hashing failed")
	}

	admin.Password = string(PassHash)
	return a.adminRepository.SaveAdmin(c, admin)

}

//Login

func (a *adminService) Login(c context.Context, admin domain.Admin) (domain.Admin, error) {
	// Check admin exist in db
	dbAdmin, err := a.adminRepository.GetAdmin(c, admin)
	if err != nil {
		return admin, err
	}
	// compare password with hash password
	if bcrypt.CompareHashAndPassword([]byte(dbAdmin.Password), []byte(admin.Password)) != nil {
		return admin, errors.New("Wrong password")
	}
	return dbAdmin, nil

}

// List all users in admin side
func (a *adminService) GetAllUser(c context.Context, page request.ReqPagination) (users []response.UserRespStrcut, err error) {
	users, err = a.adminRepository.GetAllUser(c, page)

	if err != nil {
		return nil, err
	}

	// if no error then copy users details to an array response struct
	var response []response.UserRespStrcut
	copier.Copy(&response, &users)

	return response, nil
}

// to block or unblock a user
func (a *adminService) BlockUser(c context.Context, userID uint) error {

	return a.adminRepository.BlockUser(c, userID)
}

func (o *adminService) ApproveReturnOrder(c context.Context, data request.ApproveReturnRequest) error {
	// get payment data
	// ID 2 is for status "Paid"
	payment, err := o.PaymentRepository.GetPaymentDataByOrderId(c, data.OrderID)

	if err != nil {
		return err
	}

	data.OrderTotal = payment.OrderTotal
	err = o.adminRepository.ApproveReturnOrder(c, data)
	if err != nil {
		return err
	}
	return nil
}
