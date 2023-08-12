package handler

import (
	"beautify/pkg/domain"
	service "beautify/pkg/usecase/interfaces"
	"beautify/pkg/utils"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type ProductHandler struct {
	ProductService service.ProductService
}

func NewProductHandler(prodUseCase service.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: prodUseCase,
	}
}

// AddCategory godoc
// @Summary Add a new product category or brand
// @Description Add a new product category or brand to the database
// @Tags Product
// @Param body body request.Category true "Category or Brand details"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{} "Missing or invalid entry or Failed to add brand"
// @Router /admin/brands [post]
func (p *ProductHandler) AddCategory(c *gin.Context) {
	var ProductBrand request.Category
	if err := c.ShouldBindJSON(&ProductBrand); err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, "Missing or invalid entry", err.Error(), ProductBrand)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := p.ProductService.AddCategory(c, ProductBrand)
	if err != nil {
		response := response.ErrorResponse(400, "Failed to add brand", err.Error(), ProductBrand)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly added a new brand / category in database", ProductBrand)
	c.JSON(200, response)
}

// GetAllBrands godoc
// @Summary Get a list of all product brands
// @Description Get a list of all product brands from the database
// @Tags Product
// @Success 200 {object} response.Response{} "Successfuly listed all brands"
// @Failure 500 {object} response.Response{} "Failed to get brands"
// @Router /admin/brands [get]
func (p *ProductHandler) GetAllBrands(c *gin.Context) {
	allBrands, err := p.ProductService.GetAllBrands(c)
	if err != nil {
		response := response.ErrorResponse(500, "Failed to get brands", err.Error(), allBrands)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	fmt.Println(allBrands)
	response := response.SuccessResponse(200, "Successfuly listed all brands", allBrands)
	c.JSON(200, response)
}

// AddProduct godoc
// @Summary Add a new product
// @Description Add a new product to the database
// @Tags Product
// @Param body body request.ProductReq true "Product details"
// @Success 200 {object} response.Response{} "Product added successful"
// @Failure 400 {object} response.Response{} "Missing or invalid entry or Failed to add product"
// @Router /admin/products/add [post]
func (p *ProductHandler) AddProduct(c *gin.Context) {
	var body request.ProductReq
	if err := c.ShouldBindJSON(&body); err != nil {
		responce := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, responce)
		return
	}
	var product domain.Product
	copier.Copy(&product, body)
	if err := p.ProductService.AddProduct(c, product); err != nil {
		response := response.ErrorResponse(400, "failed to add product", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(http.StatusOK, "Product added successful", body)
	c.JSON(http.StatusOK, response)
}

// AddImage godoc
// @Summary Add images to a product
// @Description Add images to a product in the database
// @Tags Product
// @Param product_id formData int true "Product ID"
// @Param image formData file true "Image file"
// @Success 200 {object} response.Response{} "Successfully added images"
// @Failure 400 {object} response.Response{} "Error while fetching product_id or Error while fetching image file or No image files found or Can't be add images"
// @Router /admin/products/addimage [post]
func (p *ProductHandler) AddImage(c *gin.Context) {
	pid, err := strconv.Atoi(c.PostForm("product_id"))
	if err != nil {
		response := response.ErrorResponse(400, "Error while fetching product_id", err.Error(), pid)
		c.JSON(400, response)
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		response := response.ErrorResponse(400, "Error while fetching image file", err.Error(), form)
		c.JSON(400, response)
		return
	}
	files := form.File["image"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image files found"})
		return
	}
	Images, err := p.ProductService.AddImage(c, pid, files)
	if err != nil {
		response := response.ErrorResponse(400, "Can't be add images", err.Error(), Images)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully added images", Images)
	c.JSON(200, response)
}

// ListProducts godoc
// @Summary List products
// @Description List products from the database
// @Tags Product
// @Param count query uint true "Count of items per page"
// @Param page_number query uint true "Page Number"
// @Success 200 {object} response.Response{} "Product listed successfuly"
// @Failure 400 {object} response.Response{} "invalid inputs"
// @Failure 500 {object} response.Response{} "failed to get all products"
// @Router /admin/products/list [get]
func (p *ProductHandler) ListProducts(c *gin.Context) {
	count, err1 := utils.StringToUint(c.Query("count"))
	if err1 != nil {
		response := response.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pageNumber, err2 := utils.StringToUint(c.Query("page_number"))
	if err2 != nil {
		response := response.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pagination := request.ReqPagination{
		PageNumber: pageNumber,
		Count:      count,
	}
	products, err := p.ProductService.GetProducts(c, pagination)
	if err != nil {
		response := response.ErrorResponse(500, "failed to get all products", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if products == nil {
		response := response.SuccessResponse(200, "Oops ! no products to show", nil)
		c.JSON(http.StatusOK, response)
		return
	}
	respones := response.SuccessResponse(200, "Product listed successfuly", products)
	c.JSON(http.StatusOK, respones)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product in the database
// @Tags Product
// @Param body body request.UpdateProductReq true "Updated product details"
// @Success 200 {object} response.Response{} "Product updated successful"
// @Failure 400 {object} response.Response{} "Missing or invalid input or failed to update product"
// @Router /admin/products/update [put]
func (p *ProductHandler) UpdateProduct(c *gin.Context) {
	var body request.UpdateProductReq
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var product domain.Product
	copier.Copy(&product, &body)
	err := p.ProductService.UpdateProduct(c, product)
	if err != nil {
		response := response.ErrorResponse(400, "failed to update product", err.Error(), body)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Product updated successful", body)
	c.JSON(200, response)
	c.Abort()
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product from the database
// @Tags Product
// @Param body body request.DeleteProductReq true "Product ID to be deleted"
// @Success 200 {object} response.Response{} "Successfully deleted product"
// @Failure 400 {object} response.Response{} "Missing or invalid input"
// @Failure 500 {object} response.Response{} "Failed to delete product"
// @Router /admin/products/delete [delete]
func (p *ProductHandler) DeleteProduct(c *gin.Context) {
	var body request.DeleteProductReq
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	productID := body.ID
	deletedProduct, err := p.ProductService.DeleteProduct(c, productID)
	if err != nil {
		response := response.ErrorResponse(500, "Failed to delete product", err.Error(), body)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := response.SuccessResponse(http.StatusOK, "Successfuly deleted product", deletedProduct)
	c.JSON(200, response)
}

func (p *ProductHandler) AddProductItem(c *gin.Context) {
	var body request.ProductItemReq

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := p.ProductService.AddProductItem(c, body)
	if err != nil {
		response := response.ErrorResponse(500, "failed to add product item", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Product item added successful", body)
	c.JSON(200, response)
	c.Abort()
}

// // Get all product item
func (p *ProductHandler) GetProductItem(c *gin.Context) {
	productID, err := utils.StringToUint(c.Param("product_id"))
	if err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, "invalid param input", err.Error(), productID)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	productItems, count, err := p.ProductService.GetProductItem(c, productID)
	// Create a map to combine the count and productItems
	fmt.Println("------------->", productItems)
	data := map[string]interface{}{
		"count":        count,
		"productItems": productItems,
	}
	if err != nil {
		response := response.ErrorResponse(400, "failed to get product item for given product id", err.Error(), nil)
		c.JSON(400, response)
		return
	}
	if count == 0 {
		response := response.ErrorResponse(http.StatusBadRequest, "No product items for this product id", "", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Fetching product item successful and listed below", data)
	c.JSON(200, response)

}
