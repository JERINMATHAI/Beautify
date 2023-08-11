package repository

import (
	"beautify/pkg/domain"
	"beautify/pkg/repository/interfaces"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB: db}
}

//________________________________________CATEGORY__________________________________________

// Find brand
func (p *productDatabase) FindBrand(ctx context.Context, brand request.Category) (request.Category, error) {
	query := `SELECT * FROM categories WHERE id=? OR category_name=?`
	if p.DB.Raw(query, brand.ID, brand.CategoryName).Scan(&brand).Error != nil {
		return request.Category{}, errors.New("failed to get brand")
	}
	return brand, nil
}

// To add brand
func (p *productDatabase) AddCategory(ctx context.Context, brand request.Category) (err error) {

	// query := `INSERT INTO categories (category_name) VALUES ($1)`
	// err = p.DB.Exec(query).Error
	err = p.DB.Create(&brand).Error

	// query := `INSERT INTO categories (parent_id, category_name)VALUES($1,$2)`
	// err = p.DB.Exec(query, brand.ParentID, brand.CategoryName).Error

	if err != nil {
		return errors.New("Failed to save brand")
	}
	return nil
}

// Get all brands from database
func (p *productDatabase) GetAllBrand(ctx context.Context) (brand []response.Brand, err error) {

	query := `SELECT c.id, c.category_name FROM categories as c`
	if p.DB.Raw(query).Scan(&brand).Error != nil {
		return brand, fmt.Errorf("failed to get brands data from db")
	}
	fmt.Println(brand)

	return brand, nil
}

//________________________________________PRODUCTS________________________________________

// Add product
func (p *productDatabase) SaveProduct(ctx context.Context, product domain.Product) error {
	query := `INSERT INTO products (name, description, category_id, price, created_at) VALUES ($1, $2, $3, $4, $5)`

	createdAt := time.Now()
	if p.DB.Exec(query, product.Name, product.Description, product.CategoryID, product.Price, createdAt).Error != nil {
		return errors.New("failed to save product on database")
	}
	return nil
}

//Add Image
func (pd *productDatabase) AddImage(c context.Context, pid int, filename string) (domain.ProductImage, error) {

	// Store the image record in the database
	image := domain.ProductImage{ProductId: uint(pid), Image: filename}
	if err := pd.DB.Create(&image).Error; err != nil {

		return domain.ProductImage{}, errors.New("failed to store image record")
	}

	return image, nil
}

// Get product
func (p *productDatabase) GetProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	query := `SELECT * FROM products where id = ? product_name = ?`
	if p.DB.Raw(query, product.ID, product.Name).Scan(&product).Error != nil {
		return product, errors.New("failure to get product")
	}
	return product, nil
}

// Find product
func (p *productDatabase) FindProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	query := `SELECT * FROM products WHERE id = ? OR name=?`
	if p.DB.Raw(query, product.ID, product.Name).Scan(&product).Error != nil {
		return product, errors.New("failed to get product")
	}
	return product, nil
}

// Find product by id
func (p *productDatabase) FindProductByID(ctx context.Context, productID uint) (product domain.Product, err error) {
	query := `SELECT * FROM products WHERE id = $1`
	err = p.DB.Raw(query, productID).Scan(&product).Error
	if err != nil {
		return product, fmt.Errorf("failed find product with prduct_id %v", productID)
	}
	return product, nil
}

// get all products from database
func (p *productDatabase) GetAllProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error) {

	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	query := `SELECT p.id, p.name, p.description, c.category_name, p.price, p.discount_price,
	p.created_at, p.updated_at, pi.image
	FROM products p LEFT JOIN categories c ON p.category_id = c.id LEFT JOIN product_images pi
	ON pi.product_id=p.id ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	if p.DB.Raw(query, limit, offset).Scan(&products).Error != nil {
		return products, errors.New("failed to get products from database")
	}

	return products, nil
}

// update product
func (p *productDatabase) UpdateProduct(ctx context.Context, product domain.Product) error {
	existingProduct, err := p.FindProductByID(ctx, product.ID)
	if err != nil {
		return err
	}
	if product.Name == "" {
		product.Name = existingProduct.Name
	}
	if product.Description == "" {
		product.Description = existingProduct.Description
	}
	if product.Price == 0 {
		product.Price = existingProduct.Price
	}
	// if product.Image == "" {
	// 	product.Image = existingProduct.Image
	// }
	if product.CategoryID == 0 {
		product.CategoryID = existingProduct.CategoryID
	}
	query := `UPDATE products SET name = $1, description = $2, category_id = $3,
	price = $4, image = $5, updated_at = $6 WHERE id = $7`

	updatedAt := time.Now()

	if p.DB.Exec(query, product.Name, product.Description, product.CategoryID,
		product.Price, updatedAt, product.ID).Error != nil {
		return errors.New("failed to update product")
	}

	return nil
}

func (p *productDatabase) DeleteProduct(ctx context.Context, productID uint) (domain.Product, error) {
	// Check requested product is exist or not
	var existingProduct domain.Product
	existingProduct, err := p.FindProductByID(ctx, productID)
	if err != nil {
		return domain.Product{}, err
	} else if existingProduct.Name == "" {
		return domain.Product{}, errors.New("invalid product_id")
	}

	//delete query
	query := `DELETE FROM products WHERE id = $1`
	if err := p.DB.Exec(query, productID).Error; err != nil {
		return domain.Product{}, fmt.Errorf("failed to delete error : %v", err)
	}
	return existingProduct, nil
}

// ___________________________________________PRODUCT ITEM_____________________________________

func (p *productDatabase) AddProductItem(ctx context.Context, productItem request.ProductItemReq) error {
	tnx := p.DB.Begin()

	// Check if the product already exists
	existingProduct, err := p.FindProductByID(ctx, productItem.ProductID)
	if err != nil {
		return err
	}
	if existingProduct.ID != productItem.ProductID {
		tnx.Rollback()
		return errors.New("product does not exist for the requested product item")
	}

	// Save the product item to the database
	query := `INSERT INTO product_items (product_id, qty_in_stock, price, discount_price, sku, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`
	createdAt := time.Now()
	if err := tnx.Raw(query, productItem.ProductID, productItem.QtyInStock, productItem.Price,
		productItem.DiscountPrice, productItem.SKU, createdAt).Scan(&productItem.ProductItemId).Error; err != nil {
		tnx.Rollback()
		return fmt.Errorf("failed to add product item: %v", err)
	}

	// Save images for the product item
	query = `INSERT INTO product_images (product_item_id, image) 
	VALUES ($1, $2)`
	for _, image := range productItem.Images {
		if err := tnx.Exec(query, productItem.ProductItemId, image).Error; err != nil {
			tnx.Rollback()
			return fmt.Errorf("failed to add product images: %v", err)
		}
	}

	if err := tnx.Commit().Error; err != nil {
		tnx.Rollback()
		return fmt.Errorf("failed to commit the transaction: %v", err)
	}

	return nil
}

// to list product item
func (p *productDatabase) GetProductItems(ctx context.Context, productId uint) ([]response.ProductItemResp, error) {
	// Check if the product ID exists
	var productItems []response.ProductItemResp

	dbProduct, err := p.FindProductByID(ctx, productId)
	if err != nil {
		return productItems, err
	}
	if dbProduct.ID == 0 {
		return productItems, errors.New("invalid product ID")
	}

	// Get product items from the database
	query := `SELECT
		p.id AS product_id,
		pi.id AS product_item_id,
		pi.qty_in_stock AS stock_available,
		p.name AS product_name,
		c.category_name AS brand,
		p.description,
		p.price,
		pi.discount_price AS offer_price
	FROM
		products p
		JOIN categories c ON c.id = p.category_id
		JOIN product_items pi ON pi.product_id = p.id
	WHERE
		p.id = $1`
	if err := p.DB.Raw(query, productId).Scan(&productItems).Error; err != nil {
		return productItems, fmt.Errorf("failed to get product items: %v", err)
	}
	fmt.Println("Product Items: ", productItems)

	// Fetch product item images
	query = `SELECT
		pimg.image
	FROM
		product_images pimg
	WHERE product_item_id = $1`
	for i := range productItems {
		productItems[i].Images = []string{}
		p.DB.Raw(query, productItems[i].ProductItemID).Scan(&productItems[i].Images)
	}
	fmt.Println("product Id: ", productId)

	return productItems, nil
}
