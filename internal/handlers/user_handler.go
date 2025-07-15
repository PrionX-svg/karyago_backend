package handlers

import (
	"github.com/gofiber/fiber/v2"
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"
	"strconv"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req request.UserEmployeeReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	creatorID := c.Locals("user_id").(uint)
	if err := h.userService.CreateUser(req, creatorID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userUUID := c.Params("uuid")

	var req request.UserEmployeeReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	modifierID := c.Locals("user_id").(uint)
	if err := h.userService.UpdateUser(userUUID, req, modifierID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

func (h *UserHandler) GetMe(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid user_id type",
		})
	}

	user, err := h.userService.GetMe(userIDUint)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch user data",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	companyUUID := c.Query("company_uuid")
	if companyUUID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "company_uuid is required",
		})
	}

	users, err := h.userService.GetAllUsers(companyUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve users",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": users,
	})
}

func (h *UserHandler) GetUserByUUID(c *fiber.Ctx) error {
	userUUID := c.Params("uuid")
	companyUUID := c.Query("company_uuid")

	user, err := h.userService.GetUserByUUID(userUUID, companyUUID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func (h *UserHandler) GetUsersDataTable(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")
	roleUUID := c.Query("role_uuid", "")
	branchUUID := c.Query("branch_uuid", "")
	isTerminated := c.Query("is_terminated", "") // "true", "false", atau "" (semua)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	data, total, filtered, err := h.userService.GetUsersWithEmployeeDataTable(
		page, limit, search, roleUUID, branchUUID, isTerminated,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch users",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":     data,
		"page":     page,
		"limit":    limit,
		"total":    total,
		"filtered": filtered,
	})
}

func (h *UserHandler) RehireUser(c *fiber.Ctx) error {
	userUUID := c.Params("uuid")
	companyUUID := c.Query("company_uuid")

	if companyUUID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "company_uuid is required",
		})
	}

	var req request.RehireEmployeeReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request payload",
			"error":   err.Error(),
		})
	}

	modifierID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	if err := h.userService.RehireEmployee(userUUID, companyUUID, req, modifierID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to rehire employee",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "employee rehired successfully",
	})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userUUID := c.Params("uuid")
	companyUUID := c.Query("company_uuid")
	terminationReason := c.Query("reason") // opsional, misalnya: "resigned", "terminated", etc.

	if userUUID == "" || companyUUID == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Missing user_uuid or company_uuid")
	}

	if err := h.userService.DeleteUser(userUUID, companyUUID, terminationReason); err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, nil, "User marked as terminated successfully")
}
