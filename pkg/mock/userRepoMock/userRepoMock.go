// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interfaces/userInterface.go

// Package mock is a generated GoMock package.
package mock

import (
	domain "beautify/pkg/domain"
	request "beautify/pkg/utils/request"
	response "beautify/pkg/utils/response"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreditUserWallet mocks base method.
func (m *MockUserRepository) CreditUserWallet(ctx context.Context, data domain.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreditUserWallet", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreditUserWallet indicates an expected call of CreditUserWallet.
func (mr *MockUserRepositoryMockRecorder) CreditUserWallet(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreditUserWallet", reflect.TypeOf((*MockUserRepository)(nil).CreditUserWallet), ctx, data)
}

// DeleteAddress mocks base method.
func (m *MockUserRepository) DeleteAddress(ctx context.Context, userID, addressID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", ctx, userID, addressID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockUserRepositoryMockRecorder) DeleteAddress(ctx, userID, addressID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockUserRepository)(nil).DeleteAddress), ctx, userID, addressID)
}

// FindUser mocks base method.
func (m *MockUserRepository) FindUser(c context.Context, user domain.Users) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUser", c, user)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUser indicates an expected call of FindUser.
func (mr *MockUserRepositoryMockRecorder) FindUser(c, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockUserRepository)(nil).FindUser), c, user)
}

// GetAllAddress mocks base method.
func (m *MockUserRepository) GetAllAddress(ctx context.Context, userId uint) ([]response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAddress", ctx, userId)
	ret0, _ := ret[0].([]response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAddress indicates an expected call of GetAllAddress.
func (mr *MockUserRepositoryMockRecorder) GetAllAddress(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAddress", reflect.TypeOf((*MockUserRepository)(nil).GetAllAddress), ctx, userId)
}

// GetCartIdByUserId mocks base method.
func (m *MockUserRepository) GetCartIdByUserId(ctx context.Context, userId uint) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartIdByUserId", ctx, userId)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartIdByUserId indicates an expected call of GetCartIdByUserId.
func (mr *MockUserRepositoryMockRecorder) GetCartIdByUserId(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartIdByUserId", reflect.TypeOf((*MockUserRepository)(nil).GetCartIdByUserId), ctx, userId)
}

// GetCartItemsbyUserId mocks base method.
func (m *MockUserRepository) GetCartItemsbyUserId(ctx context.Context, page request.ReqPagination, userID uint) ([]response.CartItemResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartItemsbyUserId", ctx, page, userID)
	ret0, _ := ret[0].([]response.CartItemResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartItemsbyUserId indicates an expected call of GetCartItemsbyUserId.
func (mr *MockUserRepositoryMockRecorder) GetCartItemsbyUserId(ctx, page, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartItemsbyUserId", reflect.TypeOf((*MockUserRepository)(nil).GetCartItemsbyUserId), ctx, page, userID)
}

// GetDefaultAddress mocks base method.
func (m *MockUserRepository) GetDefaultAddress(ctx context.Context, userId uint) (response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultAddress", ctx, userId)
	ret0, _ := ret[0].(response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultAddress indicates an expected call of GetDefaultAddress.
func (mr *MockUserRepositoryMockRecorder) GetDefaultAddress(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultAddress", reflect.TypeOf((*MockUserRepository)(nil).GetDefaultAddress), ctx, userId)
}

// GetUserbyID mocks base method.
func (m *MockUserRepository) GetUserbyID(ctx context.Context, userId uint) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserbyID", ctx, userId)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserbyID indicates an expected call of GetUserbyID.
func (mr *MockUserRepositoryMockRecorder) GetUserbyID(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserbyID", reflect.TypeOf((*MockUserRepository)(nil).GetUserbyID), ctx, userId)
}

// GetWalletHistory mocks base method.
func (m *MockUserRepository) GetWalletHistory(ctx context.Context, userId uint) ([]domain.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletHistory", ctx, userId)
	ret0, _ := ret[0].([]domain.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletHistory indicates an expected call of GetWalletHistory.
func (mr *MockUserRepositoryMockRecorder) GetWalletHistory(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletHistory", reflect.TypeOf((*MockUserRepository)(nil).GetWalletHistory), ctx, userId)
}

// GetWishIdByUserId mocks base method.
func (m *MockUserRepository) GetWishIdByUserId(ctx context.Context, userId uint) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWishIdByUserId", ctx, userId)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWishIdByUserId indicates an expected call of GetWishIdByUserId.
func (mr *MockUserRepositoryMockRecorder) GetWishIdByUserId(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWishIdByUserId", reflect.TypeOf((*MockUserRepository)(nil).GetWishIdByUserId), ctx, userId)
}

// GetWishListItemsbyUserId mocks base method.
func (m *MockUserRepository) GetWishListItemsbyUserId(ctx context.Context, page request.ReqPagination, userID uint) ([]response.WishItemResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWishListItemsbyUserId", ctx, page, userID)
	ret0, _ := ret[0].([]response.WishItemResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWishListItemsbyUserId indicates an expected call of GetWishListItemsbyUserId.
func (mr *MockUserRepositoryMockRecorder) GetWishListItemsbyUserId(ctx, page, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWishListItemsbyUserId", reflect.TypeOf((*MockUserRepository)(nil).GetWishListItemsbyUserId), ctx, page, userID)
}

// RemoveCartItem mocks base method.
func (m *MockUserRepository) RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItemReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCartItem", ctx, DelCartItem)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCartItem indicates an expected call of RemoveCartItem.
func (mr *MockUserRepositoryMockRecorder) RemoveCartItem(ctx, DelCartItem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCartItem", reflect.TypeOf((*MockUserRepository)(nil).RemoveCartItem), ctx, DelCartItem)
}

// RemoveWishListItem mocks base method.
func (m *MockUserRepository) RemoveWishListItem(ctx context.Context, DeleteWishListItemReq request.DeleteWishItemReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveWishListItem", ctx, DeleteWishListItemReq)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveWishListItem indicates an expected call of RemoveWishListItem.
func (mr *MockUserRepositoryMockRecorder) RemoveWishListItem(ctx, DeleteWishListItemReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveWishListItem", reflect.TypeOf((*MockUserRepository)(nil).RemoveWishListItem), ctx, DeleteWishListItemReq)
}

// SaveAddress mocks base method.
func (m *MockUserRepository) SaveAddress(ctx context.Context, userAddress request.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAddress", ctx, userAddress)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAddress indicates an expected call of SaveAddress.
func (mr *MockUserRepositoryMockRecorder) SaveAddress(ctx, userAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAddress", reflect.TypeOf((*MockUserRepository)(nil).SaveAddress), ctx, userAddress)
}

// SaveUser mocks base method.
func (m *MockUserRepository) SaveUser(c context.Context, user domain.Users) (response.UserSignUp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUser", c, user)
	ret0, _ := ret[0].(response.UserSignUp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveUser indicates an expected call of SaveUser.
func (mr *MockUserRepositoryMockRecorder) SaveUser(c, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockUserRepository)(nil).SaveUser), c, user)
}

// SavetoCart mocks base method.
func (m *MockUserRepository) SavetoCart(ctx context.Context, addToCart request.AddToCartReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavetoCart", ctx, addToCart)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavetoCart indicates an expected call of SavetoCart.
func (mr *MockUserRepositoryMockRecorder) SavetoCart(ctx, addToCart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavetoCart", reflect.TypeOf((*MockUserRepository)(nil).SavetoCart), ctx, addToCart)
}

// SavetoWishList mocks base method.
func (m *MockUserRepository) SavetoWishList(ctx context.Context, addToWishList request.AddToWishReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavetoWishList", ctx, addToWishList)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavetoWishList indicates an expected call of SavetoWishList.
func (mr *MockUserRepositoryMockRecorder) SavetoWishList(ctx, addToWishList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavetoWishList", reflect.TypeOf((*MockUserRepository)(nil).SavetoWishList), ctx, addToWishList)
}

// UpdateAddress mocks base method.
func (m *MockUserRepository) UpdateAddress(ctx context.Context, userAddress request.AddressPatchReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", ctx, userAddress)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockUserRepositoryMockRecorder) UpdateAddress(ctx, userAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockUserRepository)(nil).UpdateAddress), ctx, userAddress)
}

// UpdateCart mocks base method.
func (m *MockUserRepository) UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", ctx, cartUpadates)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCart indicates an expected call of UpdateCart.
func (mr *MockUserRepositoryMockRecorder) UpdateCart(ctx, cartUpadates interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockUserRepository)(nil).UpdateCart), ctx, cartUpadates)
}
