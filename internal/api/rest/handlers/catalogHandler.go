package handlers

import (
	"strconv"

	"go-ecommerce-app2/internal/api/rest"
	"go-ecommerce-app2/internal/dto"
	"go-ecommerce-app2/internal/repository"
	"go-ecommerce-app2/internal/service"

	"github.com/gofiber/fiber/v2"
)

type CatalogHandler struct {
	// svc
	svc service.CatalogService
}

func SetupCatalogRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.CatalogService{
		Repo:   repository.NewCatalogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}
	// create handler having related routes in it
	handler := CatalogHandler{svc}

	// public endpoints
	// listing products and categories
	app.Get("/products", handler.GetProducts)
	app.Get("/products/:id", handler.GetProduct)
	app.Get("/categories", handler.GetCategories)
	app.Get("/categories/:id", handler.GetCategoryByID)

	selRoutes := app.Group("/seller", rh.Auth.AuthorizeSeller)
	// private endpoints
	// manage products and categories
	// categories

	selRoutes.Post("/categories", handler.CreateCategory)
	selRoutes.Patch("/categories/:id", handler.EditCategories)
	selRoutes.Delete("/categories/:id", handler.DeleteCategories)

	// products
	selRoutes.Post("/products", handler.CreateProduct)
	selRoutes.Get("/products", handler.GetProducts)
	selRoutes.Get("/products/:id", handler.GetProduct)
	selRoutes.Patch("/products/:id", handler.UpdateProduct)
	selRoutes.Put("/products/:id", handler.UpdateStock)
	selRoutes.Delete("/products/:id", handler.DeleteProduct)
}

func (ch CatalogHandler) GetCategories(ctx *fiber.Ctx) error {
	categories, err := ch.svc.GetCategories()
	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}
	return rest.ResponseMessage(ctx, "categories", categories)
}

func (ch CatalogHandler) GetCategoryByID(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	category, err := ch.svc.GetCategory(uint(id))
	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}
	return rest.ResponseMessage(ctx, "category", category)
}

func (ch CatalogHandler) CreateCategory(ctx *fiber.Ctx) error {
	req := dto.CreateCategoryRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "invalid request for create category")
	}
	err = ch.svc.CreateCategory(req)
	if err != nil {
		return rest.InternalError(ctx, err)
	}
	return rest.ResponseMessage(ctx, "category created", nil)
}

func (ch CatalogHandler) EditCategories(ctx *fiber.Ctx) error {
	req := dto.CreateCategoryRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "invalid request for create category")
	}
	updatedCategory, err := ch.svc.EditCategory(req)
	if err != nil {
		return rest.InternalError(ctx, err)
	}
	return rest.ResponseMessage(ctx, "category updated", updatedCategory)
}

func (ch CatalogHandler) DeleteCategories(ctx *fiber.Ctx) error {
	return rest.ResponseMessage(ctx, "category deleted", nil)
}

func (ch CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {
	return rest.ResponseMessage(ctx, "product created", nil)
}

func (ch CatalogHandler) GetProducts(ctx *fiber.Ctx) error {
	return rest.ResponseMessage(ctx, "products fetched", nil)
}

func (ch CatalogHandler) GetProduct(ctx *fiber.Ctx) error {
	return rest.ResponseMessage(ctx, "product fetched", nil)
}

func (ch CatalogHandler) UpdateProduct(ctx *fiber.Ctx) error {
	return rest.ResponseMessage(ctx, "product updated", nil)
}

func (ch CatalogHandler) UpdateStock(ctx *fiber.Ctx) error {
	return rest.ResponseMessage(ctx, "stock updated", nil)
}

func (ch CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {
	return rest.ResponseMessage(ctx, "product deleted", nil)
}
