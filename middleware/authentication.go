package middleware

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/tapeds/go-fiber-template/dto"
	"github.com/tapeds/go-fiber-template/service"
	"github.com/tapeds/go-fiber-template/utils"
)

func Authenticate(jwtService service.JWTService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_FOUND, nil)
			return ctx.Status(http.StatusUnauthorized).JSON(response)
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_VALID, nil)
			return ctx.Status(http.StatusUnauthorized).JSON(response)
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_VALID, nil)
			return ctx.Status(http.StatusUnauthorized).JSON(response)
		}
		if !token.Valid {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_DENIED_ACCESS, nil)
			return ctx.Status(http.StatusUnauthorized).JSON(response)
		}
		userId, err := jwtService.GetUserIDByToken(authHeader)
		if err != nil {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, err.Error(), nil)
			return ctx.Status(http.StatusUnauthorized).JSON(response)
		}
		ctx.Locals("token", authHeader)
		ctx.Locals("user_id", userId)
		return ctx.Next()
	}
}
