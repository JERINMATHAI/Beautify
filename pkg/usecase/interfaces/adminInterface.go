package interfaces

import (
	"beautify/pkg/domain"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"context"
)

type AdminService interface {
	Signup(c context.Context, admin domain.Admin) (domain.Admin, error)
	Login(c context.Context, admin domain.Admin) (domain.Admin, error)
	GetAllUser(c context.Context, page request.ReqPagination) (users []response.UserRespStrcut, err error)
	BlockUser(c context.Context, userID uint) error
}
