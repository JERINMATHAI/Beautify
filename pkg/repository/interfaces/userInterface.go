package interfaces

import (
	"beautify/pkg/domain"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"context"
)

type UserRepository interface {
	SaveUser(c context.Context, user domain.Users) (response.UserSignUp, error)
	FindUser(c context.Context, user domain.Users) (domain.Users, error)
	GetUserbyID(ctx context.Context, userId uint) (domain.Users, error)

	//TODO
	SaveAddress(ctx context.Context, userAddress request.Address) error
	UpdateAddress(ctx context.Context, userAddress request.AddressPatchReq) error
	DeleteAddress(ctx context.Context, userID, addressID uint) error
	GetAllAddress(ctx context.Context, userId uint) (address []response.Address, err error)
	// GetEmailPhoneByUserId(ctx context.Context, userID uint) (contact response.UserContact, err error)
	GetDefaultAddress(ctx context.Context, userId uint) (address response.Address, err error)

	SavetoCart(ctx context.Context, addToCart request.AddToCartReq) error
	GetCartIdByUserId(ctx context.Context, userId uint) (cartId uint, err error)
	GetCartItemsbyUserId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error)
	UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error
	RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItemReq) error

	SavetoWishList(ctx context.Context, addToWishList request.AddToWishReq) error
	GetWishIdByUserId(ctx context.Context, userId uint) (wishId uint, err error)
	GetWishListItemsbyUserId(ctx context.Context, page request.ReqPagination, userID uint) (WishItems []response.WishItemResp, err error)
	RemoveWishListItem(ctx context.Context, DeleteWishListItemReq request.DeleteWishItemReq) error
}
