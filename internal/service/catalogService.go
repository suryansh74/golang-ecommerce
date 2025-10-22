package service

import (
	"errors"

	"go-ecommerce-app2/config"
	"go-ecommerce-app2/internal/domain"
	"go-ecommerce-app2/internal/dto"
	"go-ecommerce-app2/internal/helper"
	"go-ecommerce-app2/internal/repository"
)

// CatalogService
// @Description: Inside service inject interface(port) not repository(adapter) domain CatalogRepo
type CatalogService struct {
	Repo   repository.CatalogRespository
	Auth   helper.Auth
	Config config.AppConfig
}

func (c *CatalogService) CreateCategory(input dto.CreateCategoryRequest) error {
	err := c.Repo.CreateCategory(&domain.Category{
		Name:         input.Name,
		ImageURL:     input.ImageURL,
		DisplayOrder: input.DisplayOrder,
	})

	return err
}

func (c *CatalogService) EditCategory(id uint, input dto.CreateCategoryRequest) (*domain.Category, error) {
	updatedCategory, err := c.Repo.EditCategory(id, &domain.Category{
		Name:         input.Name,
		ImageURL:     input.ImageURL,
		DisplayOrder: input.DisplayOrder,
	})
	if err != nil {
		return nil, err
	}
	return updatedCategory, err
}

func (c *CatalogService) DeleteCategory(input any) error {
	return nil
}

func (c *CatalogService) GetCategory(id uint) (*domain.Category, error) {
	category, err := c.Repo.FindCategoryByID(id)
	if err != nil {
		return nil, errors.New("category doesnot exist")
	}
	return category, nil
}

func (c *CatalogService) GetCategories() ([]*domain.Category, error) {
	categories, err := c.Repo.FindCategories()
	if err != nil {
		return nil, errors.New("category is empty")
	}
	return categories, nil
}
