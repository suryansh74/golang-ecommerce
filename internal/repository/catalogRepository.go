package repository

import (
	"errors"

	"go-ecommerce-app2/internal/domain"

	"gorm.io/gorm"
)

type CatalogRespository interface {
	// categories
	CreateCategory(e *domain.Category) error
	FindCategories() ([]*domain.Category, error)
	FindCategoryByID(id uint) (*domain.Category, error)
	EditCategory(id uint, e *domain.Category) (*domain.Category, error)
	DeleteCategory(id uint) error

	// products
	CreateProduct(e *domain.Product) error
	FindProducts() ([]*domain.Product, error)
	FindProductByID(id uint) (*domain.Product, error)
	FindSellerProducts(id uint) ([]*domain.Product, error)
	EditProduct(id uint, e *domain.Product) (*domain.Product, error)
	DeleteProduct(id uint) error
}

type catalogRespository struct {
	db *gorm.DB
}

// CreateCategory implements CatalogRespository.
func (c catalogRespository) CreateCategory(e *domain.Category) error {
	if err := c.db.Create(e).Error; err != nil {
		return err
	}
	return nil
}

// DeleteCategory implements CatalogRespository.
func (c catalogRespository) DeleteCategory(id uint) error {
	// Check if category exists
	_, err := c.FindCategoryByID(id)
	if err != nil {
		return err
	}

	if err := c.db.Delete(&domain.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}

// EditCategory implements CatalogRespository.
func (c catalogRespository) EditCategory(id uint, e *domain.Category) (*domain.Category, error) {
	// Check if category exists
	_, err := c.FindCategoryByID(id)
	if err != nil {
		return nil, err
	}
	c.db.Model(domain.Category{}).Where("id = ?", id).Updates(e)
	category, _ := c.FindCategoryByID(id)

	return category, nil
}

// FindCategories implements CatalogRespository.
func (c catalogRespository) FindCategories() ([]*domain.Category, error) {
	var categories []*domain.Category
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// FindCategoryByID implements CatalogRespository.
func (c catalogRespository) FindCategoryByID(id uint) (*domain.Category, error) {
	var category domain.Category
	if err := c.db.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return &category, nil
}

// CreateProduct implements CatalogRespository.
func (c catalogRespository) CreateProduct(e *domain.Product) error {
	if err := c.db.Create(e).Error; err != nil {
		return err
	}
	return nil
}

// DeleteProduct implements CatalogRespository.
func (c catalogRespository) DeleteProduct(id uint) error {
	// Check if product exists
	_, err := c.FindProductByID(id)
	if err != nil {
		return err
	}

	if err := c.db.Delete(&domain.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

// EditProduct implements CatalogRespository.
func (c catalogRespository) EditProduct(id uint, e *domain.Product) (*domain.Product, error) {
	// Check if product exists
	_, err := c.FindProductByID(id)
	if err != nil {
		return nil, err
	}
	c.db.Model(domain.Product{}).Where("id = ?", id).Updates(e)
	product, _ := c.FindProductByID(id)

	return product, nil
}

// FindProducts implements CatalogRespository.
func (c catalogRespository) FindProducts() ([]*domain.Product, error) {
	var products []*domain.Product
	if err := c.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// FindProductByID implements CatalogRespository.
func (c catalogRespository) FindProductByID(id uint) (*domain.Product, error) {
	var product domain.Product
	if err := c.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

// FindSellerProducts implements CatalogRespository.
func (c catalogRespository) FindSellerProducts(id uint) ([]*domain.Product, error) {
	var products []*domain.Product
	if err := c.db.Where("seller_id = ?", id).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func NewCatalogRepository(db *gorm.DB) CatalogRespository {
	return catalogRespository{
		db: db,
	}
}
