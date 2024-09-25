package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tapeds/go-fiber-template/dto"
	"github.com/tapeds/go-fiber-template/service"
	"github.com/tapeds/go-fiber-template/utils"
)

type (
	UserController interface {
		Register(ctx *fiber.Ctx) error
		Login(ctx *fiber.Ctx) error
		Me(ctx *fiber.Ctx) error
		GetAllUser(ctx *fiber.Ctx) error
		SendVerificationEmail(ctx *fiber.Ctx) error
		VerifyEmail(ctx *fiber.Ctx) error
		Update(ctx *fiber.Ctx) error
		Delete(ctx *fiber.Ctx) error
	}

	userController struct {
		userService service.UserService
	}
)

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

func (c *userController) Register(ctx *fiber.Ctx) error {
	var user dto.UserCreateRequest

	if err := ctx.BodyParser(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.userService.RegisterUser(ctx.Context(), user)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *userController) GetAllUser(ctx *fiber.Ctx) error {
	var req dto.PaginationRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.userService.GetAllUserWithPagination(ctx.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_USER,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	return ctx.Status(http.StatusOK).JSON(resp)
}

func (c *userController) Me(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(string)

	result, err := c.userService.GetUserById(ctx.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *userController) Login(ctx *fiber.Ctx) error {
	var req dto.UserLoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := c.userService.Verify(ctx.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *userController) SendVerificationEmail(ctx *fiber.Ctx) error {
	var req dto.SendVerificationEmailRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	err := c.userService.SendVerificationEmail(ctx.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SEND_VERIFICATION_EMAIL_SUCCESS, nil)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *userController) VerifyEmail(ctx *fiber.Ctx) error {
	var req dto.VerifyEmailRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.userService.VerifyEmail(ctx.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_VERIFY_EMAIL, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_VERIFY_EMAIL, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *userController) Update(ctx *fiber.Ctx) error {
	var req dto.UserUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	userId := ctx.Locals("user_id").(string)
	result, err := c.userService.UpdateUser(ctx.Context(), req, userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *userController) Delete(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(string)

	if err := c.userService.DeleteUser(ctx.Context(), userId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_USER, nil)
	return ctx.Status(http.StatusOK).JSON(res)
}
