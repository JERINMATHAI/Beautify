package interfaces

import (
	"beautify/pkg/domain"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"context"
	"mime/multipart"
)

type ProductService interface {
	AddProduct(ctx context.Context, product domain.Product) error
	AddCategory(ctx context.Context, Category request.Category) error
	GetAllBrands(ctx context.Context) (brand []response.Brand, err error)
	GetProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error)
	UpdateProduct(ctx context.Context, product domain.Product) error
	DeleteProduct(ctx context.Context, productID uint) (domain.Product, error)
	AddProductItem(ctx context.Context, productItem request.ProductItemReq) error
	GetProductItem(ctx context.Context, productId uint) (ProductItems []response.ProductItemResp, count int, err error)
	AddImage(c context.Context, pid int, files []*multipart.FileHeader) ([]domain.ProductImage, error)
	//SKUhelper(ctx context.Context, productId uint) (interface{}, error)
}
