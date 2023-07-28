package interfaces

import (
	"beautify/pkg/domain"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"context"
)

type AdminRepository interface {
	GetAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error)
	SaveAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error)
	GetAllUser(ctx context.Context, page request.ReqPagination) (users []response.UserRespStrcut, err error)
	BlockUser(ctx context.Context, userID uint) error
}
