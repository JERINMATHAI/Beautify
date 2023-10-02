// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interfaces/userInterface.go

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

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// Addaddress mocks base method.
func (m *MockUserService) Addaddress(ctx context.Context, address request.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addaddress", ctx, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// Addaddress indicates an expected call of Addaddress.
func (mr *MockUserServiceMockRecorder) Addaddress(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addaddress", reflect.TypeOf((*MockUserService)(nil).Addaddress), ctx, address)
}

// DeleteAddress mocks base method.
func (m *MockUserService) DeleteAddress(ctx context.Context, userID, addressID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", ctx, userID, addressID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockUserServiceMockRecorder) DeleteAddress(ctx, userID, addressID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockUserService)(nil).DeleteAddress), ctx, userID, addressID)
}

// GetAllAddress mocks base method.
func (m *MockUserService) GetAllAddress(ctx context.Context, userId uint) ([]response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAddress", ctx, userId)
	ret0, _ := ret[0].([]response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAddress indicates an expected call of GetAllAddress.
func (mr *MockUserServiceMockRecorder) GetAllAddress(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAddress", reflect.TypeOf((*MockUserService)(nil).GetAllAddress), ctx, userId)
}

// GetCartItemsbyCartId mocks base method.
func (m *MockUserService) GetCartItemsbyCartId(ctx context.Context, page request.ReqPagination, userID uint) ([]response.CartItemResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartItemsbyCartId", ctx, page, userID)
	ret0, _ := ret[0].([]response.CartItemResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartItemsbyCartId indicates an expected call of GetCartItemsbyCartId.
func (mr *MockUserServiceMockRecorder) GetCartItemsbyCartId(ctx, page, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartItemsbyCartId", reflect.TypeOf((*MockUserService)(nil).GetCartItemsbyCartId), ctx, page, userID)
}

// GetWalletHistory mocks base method.
func (m *MockUserService) GetWalletHistory(ctx context.Context, userId uint) ([]domain.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletHistory", ctx, userId)
	ret0, _ := ret[0].([]domain.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletHistory indicates an expected call of GetWalletHistory.
func (mr *MockUserServiceMockRecorder) GetWalletHistory(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletHistory", reflect.TypeOf((*MockUserService)(nil).GetWalletHistory), ctx, userId)
}

// GetWishListItemsbyCartId mocks base method.
func (m *MockUserService) GetWishListItemsbyCartId(ctx context.Context, page request.ReqPagination, userID uint) ([]response.WishItemResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWishListItemsbyCartId", ctx, page, userID)
	ret0, _ := ret[0].([]response.WishItemResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWishListItemsbyCartId indicates an expected call of GetWishListItemsbyCartId.
func (mr *MockUserServiceMockRecorder) GetWishListItemsbyCartId(ctx, page, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWishListItemsbyCartId", reflect.TypeOf((*MockUserService)(nil).GetWishListItemsbyCartId), ctx, page, userID)
}

// Login mocks base method.
func (m *MockUserService) Login(ctx context.Context, user domain.Users) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, user)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceMockRecorder) Login(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserService)(nil).Login), ctx, user)
}

// OTPLogin mocks base method.
func (m *MockUserService) OTPLogin(ctx context.Context, user domain.Users) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OTPLogin", ctx, user)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OTPLogin indicates an expected call of OTPLogin.
func (mr *MockUserServiceMockRecorder) OTPLogin(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OTPLogin", reflect.TypeOf((*MockUserService)(nil).OTPLogin), ctx, user)
}

// Profile mocks base method.
func (m *MockUserService) Profile(ctx context.Context, userId uint) (response.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Profile", ctx, userId)
	ret0, _ := ret[0].(response.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Profile indicates an expected call of Profile.
func (mr *MockUserServiceMockRecorder) Profile(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Profile", reflect.TypeOf((*MockUserService)(nil).Profile), ctx, userId)
}

// RemoveCartItem mocks base method.
func (m *MockUserService) RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItemReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCartItem", ctx, DelCartItem)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCartItem indicates an expected call of RemoveCartItem.
func (mr *MockUserServiceMockRecorder) RemoveCartItem(ctx, DelCartItem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCartItem", reflect.TypeOf((*MockUserService)(nil).RemoveCartItem), ctx, DelCartItem)
}

// RemoveWishListItem mocks base method.
func (m *MockUserService) RemoveWishListItem(ctx context.Context, DelWishListItem request.DeleteWishItemReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveWishListItem", ctx, DelWishListItem)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveWishListItem indicates an expected call of RemoveWishListItem.
func (mr *MockUserServiceMockRecorder) RemoveWishListItem(ctx, DelWishListItem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveWishListItem", reflect.TypeOf((*MockUserService)(nil).RemoveWishListItem), ctx, DelWishListItem)
}

// SaveCartItem mocks base method.
func (m *MockUserService) SaveCartItem(ctx context.Context, addToCart request.AddToCartReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCartItem", ctx, addToCart)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCartItem indicates an expected call of SaveCartItem.
func (mr *MockUserServiceMockRecorder) SaveCartItem(ctx, addToCart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCartItem", reflect.TypeOf((*MockUserService)(nil).SaveCartItem), ctx, addToCart)
}

// SaveWishListItem mocks base method.
func (m *MockUserService) SaveWishListItem(ctx context.Context, addToWishList request.AddToWishReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveWishListItem", ctx, addToWishList)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveWishListItem indicates an expected call of SaveWishListItem.
func (mr *MockUserServiceMockRecorder) SaveWishListItem(ctx, addToWishList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveWishListItem", reflect.TypeOf((*MockUserService)(nil).SaveWishListItem), ctx, addToWishList)
}

// SignUp mocks base method.
func (m *MockUserService) SignUp(ctx context.Context, user domain.Users) (response.UserSignUp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", ctx, user)
	ret0, _ := ret[0].(response.UserSignUp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUserServiceMockRecorder) SignUp(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUserService)(nil).SignUp), ctx, user)
}

// UpdateAddress mocks base method.
func (m *MockUserService) UpdateAddress(ctx context.Context, address request.AddressPatchReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", ctx, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockUserServiceMockRecorder) UpdateAddress(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockUserService)(nil).UpdateAddress), ctx, address)
}

// UpdateCart mocks base method.
func (m *MockUserService) UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", ctx, cartUpadates)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCart indicates an expected call of UpdateCart.
func (mr *MockUserServiceMockRecorder) UpdateCart(ctx, cartUpadates interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockUserService)(nil).UpdateCart), ctx, cartUpadates)
}
