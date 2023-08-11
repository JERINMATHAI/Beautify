package repository

import (
	"beautify/pkg/domain"
	repository "beautify/pkg/repository/interfaces"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) repository.UserRepository {
	return &userDatabase{DB: DB}
}

// Find user is existing or not
func (u *userDatabase) FindUser(c context.Context, user domain.Users) (domain.Users, error) {
	query := `SELECT * FROM users where id=? OR user_name=? OR email=? OR phone=?`
	if err := u.DB.Raw(query, user.ID, user.UserName, user.Email, user.Phone).Scan(&user).Error; err != nil {
		return user, errors.New("Failed to find user")
	}
	return user, nil
}

// Save the user if the user is not existing
func (u *userDatabase) SaveUser(c context.Context, user domain.Users) (response.UserSignUp, error) {
	var usersignup response.UserSignUp
	query := `INSERT INTO users (user_name, first_name, last_name, age, email, phone, password,created_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	createdAt := time.Now()
	if u.DB.Exec(query, user.UserName, user.FirstName, user.LastName, user.Age, user.Email, user.Phone, user.Password, createdAt).Error != nil {
		return response.UserSignUp{}, errors.New("Failed to save user")
	}
	query2 := `SELECT id, user_name from users where first_name=?`
	if err := u.DB.Raw(query2, user.FirstName).Scan(&usersignup).Error; err != nil {
		return response.UserSignUp{}, errors.New("Failed to find user")
	}
	return usersignup, nil
}

//Save user adddress
func (i *userDatabase) SaveAddress(ctx context.Context, userAddress request.Address) error {
	var defaultAddressID uint
	userAddress.CreatedAt = time.Now()
	query := `INSERT INTO addresses (user_id ,house,address_line1,address_line2,city,state,pin_code,country,created_at, is_default) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	if err := i.DB.Raw(query, userAddress.UserID, userAddress.House, userAddress.AddressLine1,
		userAddress.AddressLine2, userAddress.City, userAddress.State, userAddress.PinCode, userAddress.Country, userAddress.CreatedAt, userAddress.IsDefault).Scan(&defaultAddressID).Error; err != nil {
		return err
	}
	// set as default if no existing default address
	query = `UPDATE addresses
	SET is_default = true
	WHERE user_id = $1
	AND is_default = false AND id = $2
	AND NOT EXISTS (
	  SELECT 1
	  FROM addresses
	  WHERE user_id = $1
	  AND is_default = true
	)`
	if err := i.DB.Exec(query, userAddress.UserID, defaultAddressID).Error; err != nil {
		return err
	}
	return nil
}

func (i *userDatabase) UpdateAddress(ctx context.Context, userAddress request.AddressPatchReq) error {
	tnx := i.DB.Begin()
	// Set all addresses of the user to false except for the new address if new address is default
	if userAddress.IsDefault {
		resetQuery := `UPDATE addresses
				SET is_default = false
				WHERE user_id = $1`
		if err := tnx.Exec(resetQuery, userAddress.UserID).Error; err != nil {
			tnx.Rollback()
			return err
		}
	}
	query := `UPDATE addresses
	SET
		house = COALESCE(NULLIF($1, ''), house),
		address_line1 = COALESCE(NULLIF($2, ''), address_line1),
		address_line2 = COALESCE(NULLIF($3, ''), address_line2),
		city = COALESCE(NULLIF($4, ''), city),
		state = COALESCE(NULLIF($5, ''), state),
		pin_code = COALESCE(NULLIF($6, ''), pin_code),
		country = COALESCE(NULLIF($7, ''), country),
		is_default = COALESCE($8, is_default),
		updated_at = $11
	WHERE
		user_id = $9
		AND id = $10`
	userAddress.UpdatedAt = time.Now()
	if err := tnx.Exec(query, userAddress.House, userAddress.AddressLine1, userAddress.AddressLine2,
		userAddress.City, userAddress.State, userAddress.PinCode,
		userAddress.Country, userAddress.IsDefault, userAddress.UserID, userAddress.ID, userAddress.UpdatedAt).Error; err != nil {
		tnx.Rollback()
		return err
	}
	err := tnx.Commit().Error
	if err != nil {
		return err
	}
	return nil
}

func (i *userDatabase) DeleteAddress(ctx context.Context, userID, addressID uint) error {
	query := `DELETE FROM addresses WHERE user_id = $1 AND id = $2`
	if err := i.DB.Exec(query, userID, addressID).Error; err != nil {
		return err
	}
	return nil
}
func (u *userDatabase) GetAllAddress(ctx context.Context, userId uint) (address []response.Address, err error) {
	query := `SELECT * FROM addresses WHERE user_id = ? ORDER BY is_default DESC, updated_at ASC`

	if err := u.DB.Raw(query, userId).Scan(&address).Error; err != nil {
		return address, err
	}
	return address, nil
}

func (i *userDatabase) GetDefaultAddress(ctx context.Context, userId uint) (address response.Address, err error) {
	//fmt.Println("user id", userId)
	query := `SELECT a.id, a.house, a.address_line1, a.address_line2, a.city, a.state, a.pin_code, a.country, a.is_default
FROM addresses as a
WHERE a.user_id = ? AND a.is_default = true`
	if err := i.DB.Raw(query, userId).Scan(&address).Error; err != nil {
		return address, err
	}
	//fmt.Println("Address", address)
	return address, nil
}

// user id
func (i *userDatabase) GetUserbyID(ctx context.Context, userId uint) (domain.Users, error) {
	var user domain.Users
	query := `SELECT * FROM users WHERE id = ?`
	if err := i.DB.Raw(query, userId).Scan(&user).Error; err != nil {
		return user, err
	}
	//fmt.Println(userId)
	return user, nil
}

func (i *userDatabase) SavetoCart(ctx context.Context, addToCart request.AddToCartReq) error {
	// Get product price from the product table
	query := `SELECT price FROM products WHERE id = $1`
	if err := i.DB.Raw(query, addToCart.ProductID).Scan(&addToCart.Price).Error; err != nil {
		return err
	}
	if addToCart.Price == 0 {
		return errors.New("invalid product ID")
	}

	// Get cart ID with user ID
	query = `SELECT id FROM carts WHERE user_id = $1`
	var cartID int
	if err := i.DB.Raw(query, addToCart.UserID).Scan(&cartID).Error; err != nil {
		return err
	}
	if cartID == 0 {
		// Create a cart for the user with the given userID if it doesn't exist
		query = `INSERT INTO carts (user_id) VALUES ($1) RETURNING id`
		if err := i.DB.Raw(query, addToCart.UserID).Scan(&cartID).Error; err != nil {
			return err
		}
	}

	// Check if the product already exists in the cart
	query = `SELECT id FROM cart_items WHERE product_id = $1 AND cart_id = $2`
	var cartItemID int
	if err := i.DB.Raw(query, addToCart.ProductID, cartID).Scan(&cartItemID).Error; err != nil {
		return err
	}
	if cartItemID != 0 {
		// Update the existing cart item with the new quantity and update the timestamp
		query = `UPDATE cart_items SET quantity = quantity + $1, updated_at = $2 WHERE id = $3`
		updatedAt := time.Now()
		if err := i.DB.Exec(query, addToCart.Quantity, updatedAt, cartItemID).Error; err != nil {
			return fmt.Errorf("failed to save cart item %v", addToCart.ProductID)
		}
	} else {
		// Insert the product into the cart_items table
		query = `INSERT INTO cart_items (cart_id, product_id, quantity, price, created_at)
		VALUES ($1, $2, $3, $4, $5)`
		createdAt := time.Now()
		if err := i.DB.Exec(query, cartID, addToCart.ProductID, addToCart.Quantity, addToCart.Price, createdAt).Error; err != nil {
			return fmt.Errorf("failed to save cart item %v", addToCart.ProductID)
		}
	}

	var cartItems []domain.CartItems
	if err := i.DB.Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		return err
	}

	// Calculate the new total based on the updated cart items
	var total float64
	for _, item := range cartItems {
		total += float64(item.Quantity) * item.Price
	}

	if err := i.DB.Exec("UPDATE carts SET total = $1 WHERE user_id = $2", total, addToCart.UserID).Error; err != nil {
		return err
	}

	return nil
}

func (i *userDatabase) GetCartIdByUserId(ctx context.Context, userId uint) (cartId uint, err error) {
	query := `SELECT id FROM carts WHERE user_id = $1`
	if err := i.DB.Raw(query, userId).Scan(&cartId).Error; err != nil {
		return cartId, err
	}
	return cartId, nil
}

func (i *userDatabase) GetCartItemsbyUserId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error) {

	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	// Get cart ID by user ID
	cartID, err := i.GetCartIdByUserId(ctx, userID)
	if err != nil {
		return CartItems, err
	}

	// Retrieve cart items with cart ID
	query := `
		SELECT ci.cart_id, p.name, p.price, ci.price AS discount_price,
			ci.quantity, ci.price * ci.quantity AS sub_total
		FROM cart_items ci
		JOIN products p ON ci.product_id = p.id
		WHERE cart_id = $1
		ORDER BY ci.created_at DESC
		LIMIT $2 OFFSET $3
	`
	if err := i.DB.Raw(query, cartID, limit, offset).Scan(&CartItems).Error; err != nil {
		return CartItems, err
	}

	return CartItems, nil
}

func (i *userDatabase) UpdateCart(ctx context.Context, cartUpdates request.UpdateCartReq) error {
	// Get cartID by user id
	cartID, err := i.GetCartIdByUserId(ctx, cartUpdates.UserID)
	if err != nil {
		return err
	}

	// Update cart
	query := `
		UPDATE cart_items 
		SET quantity = COALESCE($3, quantity),
		updated_at = $4
		WHERE cart_id = $1 AND product_id = $2
	`
	updatedAt := time.Now()

	if err := i.DB.Exec(query, cartID, cartUpdates.ProductID, cartUpdates.Quantity, updatedAt).Error; err != nil {
		return err
	}

	return nil
}
func (i *userDatabase) RemoveCartItem(ctx context.Context, DeleteCartItemReq request.DeleteCartItemReq) error {
	//get cart id by user id
	cartID, err := i.GetCartIdByUserId(ctx, DeleteCartItemReq.UserID)
	if err != nil {
		return err
	}
	//delete cart items
	query := `DELETE FROM cart_items WHERE cart_id =$1 AND product_id=$2`
	if err := i.DB.Exec(query, cartID, DeleteCartItemReq.ProductID).Error; err != nil {
		return err
	}
	return nil
}

//Wish List
func (i *userDatabase) SavetoWishList(ctx context.Context, addToWishList request.AddToWishReq) error {
	// Get product price from the product table
	query := `SELECT price FROM products WHERE id = $1`
	if err := i.DB.Raw(query, addToWishList.ProductID).Scan(&addToWishList.Price).Error; err != nil {
		return err
	}
	if addToWishList.Price == 0 {
		return errors.New("invalid product ID")
	}

	// Get Wish List ID with user ID
	query = `SELECT id FROM wish_lists WHERE user_id = $1`
	var wishID int
	if err := i.DB.Raw(query, addToWishList.UserID).Scan(&wishID).Error; err != nil {
		return err
	}
	if wishID == 0 {
		// Create a wish list for the user with the given userID if it doesn't exist
		query = `INSERT INTO wish_lists (user_id) VALUES ($1) RETURNING id`
		if err := i.DB.Raw(query, addToWishList.UserID).Scan(&wishID).Error; err != nil {
			return err
		}
	}

	// Check if the product already exists in the wish list
	query = `SELECT id FROM wish_list_items WHERE product_id = $1 AND wish_id = $2`
	var wishlistItemID int
	if err := i.DB.Raw(query, addToWishList.ProductID, wishID).Scan(&wishlistItemID).Error; err != nil {
		return err
	}
	if wishlistItemID != 0 {
		// Update the existing wish list item with  update the timestamp
		query = `UPDATE wish_list_items SET  updated_at = $1 WHERE id = $2`
		updatedAt := time.Now()
		if err := i.DB.Exec(query, updatedAt, wishlistItemID).Error; err != nil {
			return fmt.Errorf("failed to add wish list item %v", addToWishList.ProductID)
		}
	} else {
		// Insert the product into the wish_list_items table
		query = `INSERT INTO wish_list_items (wish_id, product_id, price, created_at)
		VALUES ($1, $2, $3, $4, $5)`
		createdAt := time.Now()
		if err := i.DB.Exec(query, wishID, addToWishList.ProductID, addToWishList.Price, createdAt).Error; err != nil {
			return fmt.Errorf("failed to add wish list item %v", addToWishList.ProductID)
		}
	}

	var wishlistItems []domain.WishListItems
	if err := i.DB.Where("wish_id = ?", wishID).Find(&wishlistItems).Error; err != nil {
		return err
	}

	// Calculate the new total based on the updated wish list items
	var total float64
	for _, item := range wishlistItems {
		total += float64(item.Quantity) * item.Price
	}

	if err := i.DB.Exec("UPDATE wish_lists SET total = $1 WHERE user_id = $2", total, addToWishList.UserID).Error; err != nil {
		return err
	}

	return nil
}

func (i *userDatabase) GetWishIdByUserId(ctx context.Context, userId uint) (wishId uint, err error) {
	query := `SELECT id FROM wish_lists WHERE user_id = $1`
	if err := i.DB.Raw(query, userId).Scan(&wishId).Error; err != nil {
		return wishId, err
	}
	return wishId, nil
}

func (i *userDatabase) GetWishListItemsbyUserId(ctx context.Context, page request.ReqPagination, userID uint) (WishItems []response.WishItemResp, err error) {

	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	// Get cart ID by user ID
	wishID, err := i.GetWishIdByUserId(ctx, userID)
	if err != nil {
		return WishItems, err
	}

	// Retrieve cart items with cart ID
	query := `
		SELECT wi.wish_id, p.name, p.price, wi.price AS discount_price
		FROM wish_list_items wi
		JOIN products p ON wi.product_id = p.id
		WHERE wish_id = $1
		ORDER BY wi.created_at DESC
		LIMIT $2 OFFSET $3
	`
	if err := i.DB.Raw(query, wishID, limit, offset).Scan(&WishItems).Error; err != nil {
		return WishItems, err
	}

	return WishItems, nil
}

func (i *userDatabase) RemoveWishListItem(ctx context.Context, DeleteWishListItemReq request.DeleteWishItemReq) error {
	//get wish list id by user id
	wishID, err := i.GetWishIdByUserId(ctx, DeleteWishListItemReq.UserID)
	if err != nil {
		return err
	}
	//delete wish list items
	query := `DELETE FROM wish_list_items WHERE wish_id =$1 AND product_id=$2`
	if err := i.DB.Exec(query, wishID, DeleteWishListItemReq.ProductID).Error; err != nil {
		return err
	}
	return nil
}
