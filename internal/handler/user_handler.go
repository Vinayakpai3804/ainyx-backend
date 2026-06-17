package handler

import (
	"ainyx-backend/internal/models"
	"ainyx-backend/internal/service"
	"ainyx-backend/internal/logger"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserHandler struct {
	service  service.UserService
	validate *validator.Validate
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := h.validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	user, err := h.service.CreateUser(c.Context(), req)
	if err != nil {
		logger.Log.Error("Failed to create user", zap.Error(err))
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}
	logger.Log.Info("User created", zap.Int32("id", user.ID))
	return c.Status(201).JSON(user)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	user, err := h.service.GetUserByID(c.Context(), int32(id))
	if err != nil {
		logger.Log.Error("User not found", zap.Int("id", id))
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.Status(200).JSON(user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := h.validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	user, err := h.service.UpdateUser(c.Context(), int32(id), req)
	if err != nil {
		logger.Log.Error("Failed to update user", zap.Int("id", id))
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}
	logger.Log.Info("User updated", zap.Int32("id", user.ID))
	return c.Status(200).JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := h.service.DeleteUser(c.Context(), int32(id)); err != nil {
		logger.Log.Error("Failed to delete user", zap.Int("id", id))
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}
	logger.Log.Info("User deleted", zap.Int("id", id))
	return c.SendStatus(204)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	page := int32(c.QueryInt("page", 1))
	limit := int32(c.QueryInt("limit", 10))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	users, err := h.service.ListUsers(c.Context(), page, limit)
	if err != nil {
		logger.Log.Error("Failed to list users", zap.Error(err))
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}
	return c.Status(200).JSON(users)
}