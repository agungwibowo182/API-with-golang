package cart

import (
	"strconv"
	"tugas-bootcamp/helper"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service Service
}

func NewCartHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) CreateCart(c *fiber.Ctx) error {
	var cart Cart
	if err := c.BodyParser(&cart); err != nil {
		response := helper.APIResponse("Failed to parse request body", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	if err := h.service.CreateCart(&cart); err != nil {
		response := helper.APIResponse("Failed to create cart", fiber.StatusInternalServerError, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helper.APIResponse("Cart created successfully", fiber.StatusOK, "success", cart)
	return c.JSON(response)
}

func (h *Handler) GetAllCarts(c *fiber.Ctx) error {
	carts, err := h.service.GetAllCarts()
	if err != nil {
		response := helper.APIResponse("Failed to fetch carts", fiber.StatusInternalServerError, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helper.APIResponse("Carts fetched successfully", fiber.StatusOK, "success", carts)
	return c.JSON(response)
}

func (h *Handler) GetCartByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := helper.APIResponse("Invalid cart ID", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	cart, err := h.service.GetCartByID(id)
	if err != nil {
		response := helper.APIResponse("Failed to fetch cart", fiber.StatusInternalServerError, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helper.APIResponse("Cart fetched successfully", fiber.StatusOK, "success", cart)
	return c.JSON(response)
}

func (h *Handler) UpdateCart(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := helper.APIResponse("Invalid cart ID", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	var updatedCart Cart
	if err := c.BodyParser(&updatedCart); err != nil {
		response := helper.APIResponse("Failed to parse request body", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	updatedCart.ID = id
	if err := h.service.UpdateCart(&updatedCart); err != nil {
		response := helper.APIResponse("Failed to update cart", fiber.StatusInternalServerError, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helper.APIResponse("Cart updated successfully", fiber.StatusOK, "success", updatedCart)
	return c.JSON(response)
}

func (h *Handler) DeleteCart(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := helper.APIResponse("Invalid cart ID", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	if err := h.service.DeleteCart(id); err != nil {
		response := helper.APIResponse("Failed to delete cart", fiber.StatusInternalServerError, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helper.APIResponse("Cart deleted successfully", fiber.StatusOK, "success", nil)
	return c.JSON(response)
}
