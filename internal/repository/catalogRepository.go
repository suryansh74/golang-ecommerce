package repository

import (
	"errors"

	"go-ecommerce-app2/internal/domain"

	"gorm.io/gorm"
)

type CatalogRespository interface {
	CreateCategory(e *domain.Category) error
	FindCategories() ([]*domain.Category, error)
	FindCategoryByID(id uint) (*domain.Category, error)
	EditCategory(id uint, e *domain.Category) (*domain.Category, error)
	DeleteCategory(id uint) error
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
	_, err := c.FindCategoryByID(e.ID)
	if err != nil {
		return nil, err
	}

	if err := c.db.Save(e).Error; err != nil {
		return nil, err
	}
	return e, nil
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

func NewCatalogRepository(db *gorm.DB) CatalogRespository {
	return catalogRespository{
		db: db,
	}
}
