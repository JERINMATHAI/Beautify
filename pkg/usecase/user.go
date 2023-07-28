package usecase

import (
	"beautify/pkg/domain"
	"beautify/pkg/repository/interfaces"
	service "beautify/pkg/usecase/interfaces"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"beautify/pkg/verify"
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepository interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) service.UserService {
	return &UserUseCase{userRepository: repo}
}

func (u *UserUseCase) SignUp(ctx context.Context, user domain.Users) error {
	// Check if user already exist
	DBUser, err := u.userRepository.FindUser(ctx, user)
	if err != nil {
		return err
	}

	if DBUser.ID == 0 {
		// Hash user password
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			fmt.Println("Hashing failed")
			return err
		}
		user.Password = string(hashedPass)

		// Save user if not exist
		err = u.userRepository.SaveUser(ctx, user)
		if err != nil {
			return err
		}

	} else {
		return fmt.Errorf("%v user already exists", DBUser.UserName)
	}

	return nil
}

func (u *UserUseCase) Login(c context.Context, user domain.Users) (domain.Users, error) {
	// Find user in db
	DBUser, err := u.userRepository.FindUser(c, user)
	if err != nil {
		return user, err
	} else if DBUser.ID == 0 {
		return user, errors.New("user not exist")
	}
	//Check if the user blocked by admin
	if DBUser.BlockStatus {
		return user, errors.New("user blocked by admin")
	}

	if _, err := verify.TwilioSendOTP("+91" + DBUser.Phone); err != nil {
		return user, fmt.Errorf("failed to send otp %v",
			err)
	}
	// check password with hashed pass
	if bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(user.Password)) != nil {
		return user, errors.New("password incorrect")
	}

	return DBUser, nil

}

func (u *UserUseCase) OTPLogin(ctx context.Context, user domain.Users) (domain.Users, error) {
	// Find user in db
	DBUser, err := u.userRepository.FindUser(ctx, user)
	if err != nil {
		return user, err
	} else if DBUser.ID == 0 {
		return user, errors.New("user not exist")
	}
	return DBUser, nil
}

func (u *UserUseCase) Addaddress(ctx context.Context, address request.Address) error {
	if err := u.userRepository.SaveAddress(ctx, address); err != nil {
		return err
	}
	return nil
}
func (u *UserUseCase) UpdateAddress(ctx context.Context, address request.AddressPatchReq) error {
	if err := u.userRepository.UpdateAddress(ctx, address); err != nil {
		return err
	}
	return nil

}
func (u *UserUseCase) DeleteAddress(ctx context.Context, userID, addressID uint) error {
	if err := u.userRepository.DeleteAddress(ctx, userID, addressID); err != nil {
		return err
	}
	return nil
}
func (u *UserUseCase) GetAllAddress(ctx context.Context, userId uint) (address []response.Address, err error) {
	address, err = u.userRepository.GetAllAddress(ctx, userId)
	if err != nil {
		return address, err
	}
	return address, nil
}
func (u *UserUseCase) SaveCartItem(ctx context.Context, addToCart request.AddToCartReq) error {
	if err := u.userRepository.SavetoCart(ctx, addToCart); err != nil {
		return err
	}
	return nil
}
func (u *UserUseCase) GetCartItemsbyCartId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error) {
	cartItems, err := u.userRepository.GetCartItemsbyUserId(ctx, page, userID)
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}

func (u *UserUseCase) UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error {
	if err := u.userRepository.UpdateCart(ctx, cartUpadates); err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItemReq) error {
	if err := u.userRepository.RemoveCartItem(ctx, DelCartItem); err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) Profile(ctx context.Context, userId uint) (response.Profile, error) {
	user, err := u.userRepository.GetUserbyID(ctx, userId)
	if err != nil {
		return response.Profile{}, err
	}

	defaultAddress, err := u.userRepository.GetDefaultAddress(ctx, userId)
	if err != nil {
		return response.Profile{}, err
	}
	fmt.Println(defaultAddress)

	profile := response.Profile{
		ID:             user.ID,
		UserName:       user.UserName,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Age:            user.Age,
		Email:          user.Email,
		Phone:          user.Phone,
		DefaultAddress: defaultAddress,
	}

	return profile, nil
}

func (u *UserUseCase) SaveWishListItem(ctx context.Context, addToWishList request.AddToWishReq) error {
	if err := u.userRepository.SavetoWishList(ctx, addToWishList); err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) GetWishListItemsbyCartId(ctx context.Context, page request.ReqPagination, userID uint) (WishItems []response.WishItemResp, err error) {
	wishItems, err := u.userRepository.GetWishListItemsbyUserId(ctx, page, userID)
	if err != nil {
		return nil, err
	}
	return wishItems, nil
}

func (u *UserUseCase) RemoveWishListItem(ctx context.Context, DelWishListItem request.DeleteWishItemReq) error {
	if err := u.userRepository.RemoveWishListItem(ctx, DelWishListItem); err != nil {
		return err
	}
	return nil
}
