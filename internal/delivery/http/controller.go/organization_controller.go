package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lutfiandri/golang-clean-architecture/internal/helper"
	"github.com/lutfiandri/golang-clean-architecture/internal/model"
	"github.com/lutfiandri/golang-clean-architecture/internal/usecase"
)

type OrganizationController interface {
	Create(c *fiber.Ctx) error
	GetMany(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type organizationController struct {
	app                 *fiber.App
	validate            *validator.Validate
	organizationUseCase usecase.OrganizationUseCase
}

func NewOrganizationController(app *fiber.App, validate *validator.Validate, organizationUseCase usecase.OrganizationUseCase) OrganizationController {
	return &organizationController{
		app:                 app,
		validate:            validate,
		organizationUseCase: organizationUseCase,
	}
}

func (controller *organizationController) Create(c *fiber.Ctx) error {
	var request model.CreateOrganizationRequest
	parseOption := helper.ParseOptions{ParseBody: true}
	if err := helper.ParseAndValidateRequest[model.CreateOrganizationRequest](c, controller.validate, &request, parseOption); err != nil {
		return err
	}

	response, err := controller.organizationUseCase.Create(&request)
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func (controller *organizationController) GetMany(c *fiber.Ctx) error {
	var request model.GetManyOrganizationRequest
	parseOption := helper.ParseOptions{ParseQuery: true}
	if err := helper.ParseAndValidateRequest[model.GetManyOrganizationRequest](c, controller.validate, &request, parseOption); err != nil {
		return err
	}

	response, err := controller.organizationUseCase.GetMany(&request)
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func (controller *organizationController) Get(c *fiber.Ctx) error {
	var request model.GetOrganizationRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[model.GetOrganizationRequest](c, controller.validate, &request, parseOption); err != nil {
		return err
	}

	response, err := controller.organizationUseCase.Get(&request)
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func (controller *organizationController) Update(c *fiber.Ctx) error {
	var request model.UpdateOrganizationRequest
	parseOption := helper.ParseOptions{ParseBody: true, ParseParams: true}
	if err := helper.ParseAndValidateRequest[model.UpdateOrganizationRequest](c, controller.validate, &request, parseOption); err != nil {
		return err
	}

	response, err := controller.organizationUseCase.Update(&request)
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func (controller *organizationController) Delete(c *fiber.Ctx) error {
	var request model.DeleteOrganizationRequest
	parseOption := helper.ParseOptions{ParseParams: true}
	if err := helper.ParseAndValidateRequest[model.DeleteOrganizationRequest](c, controller.validate, &request, parseOption); err != nil {
		return err
	}

	response, err := controller.organizationUseCase.Delete(&request)
	if err != nil {
		return err
	}

	return c.JSON(response)
}
