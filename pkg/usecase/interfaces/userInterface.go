package interfaces

import (
	"beautify/pkg/domain"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"context"
)

type UserService interface {
	SignUp(ctx context.Context, user domain.Users) (usersignup response.UserSignUp, err error)
	Login(ctx context.Context, user domain.Users) (domain.Users, error)
	OTPLogin(ctx context.Context, user domain.Users) (domain.Users, error)
	Addaddress(ctx context.Context, address request.Address) error
	UpdateAddress(ctx context.Context, address request.AddressPatchReq) error
	DeleteAddress(ctx context.Context, userID, addressID uint) error
	GetAllAddress(ctx context.Context, userId uint) (address []response.Address, err error)
	SaveCartItem(ctx context.Context, addToCart request.AddToCartReq) error
	GetCartItemsbyCartId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error)
	UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error
	RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItemReq) error
	SaveWishListItem(ctx context.Context, addToWishList request.AddToWishReq) error
	GetWishListItemsbyCartId(ctx context.Context, page request.ReqPagination, userID uint) (WishItems []response.WishItemResp, err error)
	RemoveWishListItem(ctx context.Context, DelWishListItem request.DeleteWishItemReq) error

	Profile(ctx context.Context, userId uint) (profile response.Profile, err error)
	GetWalletHistory(ctx context.Context, userId uint) (wallet []domain.Wallet, err error)
}
