package handler

import (
	"go-tabungan-api/helper"
	"go-tabungan-api/tabungan"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type tabunganHandler struct {
	service tabungan.Service
}

func NewTabunganHandler(service tabungan.Service) *tabunganHandler {
	return &tabunganHandler{service}
}

func (h *tabunganHandler) CreateTabungan(c *fiber.Ctx) error {
	var input tabungan.CreateTabunganInput

	err := c.BodyParser(&input)
	if err != nil {
		errors := helper.FormatValidation(err)
		errorMessage := fiber.Map{"errors": errors}
		response := helper.APIResponse("Failed to Create Tabungan", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	isDataAvailable, err := h.service.IsTabunganAvailable(input)
	if err != nil {
		errorMessage := fiber.Map{"errors": "Server error"}
		response := helper.APIResponse("NIK or No HP checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	data := fiber.Map{
		"isAvailableData": isDataAvailable,
	}

	if isDataAvailable{
		response := helper.APIResponse("Data NIK or No HP is available", http.StatusOK, "success", data)
		return c.Status(http.StatusOK).JSON(response)
	}

	newTabungan, err := h.service.CreateTabungan(input)
	if err != nil {
		response := helper.APIResponse("Failed to Create Tabungan", http.StatusUnprocessableEntity, "error", nil)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	response := helper.APIResponse("Success create Tabungan", http.StatusOK, "success", tabungan.FormatTabungan(newTabungan))
	return c.Status(http.StatusOK).JSON(response)
}

func (h *tabunganHandler) TabungMutasi(c *fiber.Ctx) error {
	var input tabungan.CreateMutasiInput

	err := c.BodyParser(&input)
	if err != nil {
		errors := helper.FormatValidation(err)
		errorMessage := fiber.Map{"errors": errors}
		response := helper.APIResponse("Failed to Create Mutasi Tabungan", http.StatusUnprocessableEntity, "Input Data Required", errorMessage)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	input.Type = "C"

	NewMutasi, err := h.service.CreateMutasi(input)
	if err != nil {
		response := helper.APIResponse("Failed to Create Mutasi Tabungan", http.StatusUnprocessableEntity, "Nomor Rekening Not Available", nil)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	response := helper.APIResponse("Success create Mutasi Tabungan", http.StatusOK, "success", tabungan.FormatSaldo(NewMutasi))
	return c.Status(http.StatusOK).JSON(response)
}

func (h *tabunganHandler) TarikMutasi(c *fiber.Ctx) error {
	var input tabungan.CreateMutasiInput

	err := c.BodyParser(&input)
	if err != nil {
		errors := helper.FormatValidation(err)
		errorMessage := fiber.Map{"errors": errors}
		response := helper.APIResponse("Failed to Create Mutasi Penarikan", http.StatusUnprocessableEntity, "Input Data Required", errorMessage)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	input.Type = "D"

	NewMutasi, err := h.service.CreateMutasi(input)
	if err != nil {
		response := helper.APIResponse("Failed to Create Mutasi Penarikan", http.StatusUnprocessableEntity, "Nomor Rekening Not Available", err)
		return c.Status(http.StatusUnprocessableEntity).JSON(response)
	}

	response := helper.APIResponse("Success create campaign", http.StatusOK, "success", tabungan.FormatSaldo(NewMutasi))
	return c.Status(http.StatusOK).JSON(response)
}

func (h *tabunganHandler) GetSaldo(c *fiber.Ctx) error {
	NoRekening := c.Params("no_rekening")
	noRekening, err := strconv.Atoi(NoRekening)
	if err != nil {
		response := helper.APIResponse("Failed to get URI", http.StatusBadRequest, "Input Data Required", noRekening)
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	tabunganDetail, err := h.service.GetTabunganByNoRek(noRekening)
	if err != nil {
		response := helper.APIResponse("Failed to get detail Tabungan", http.StatusBadRequest, "error", err)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("Detail Saldo", http.StatusOK, "success", tabungan.SaldoTabunganDetail(tabunganDetail))
	return c.Status(http.StatusOK).JSON(response)
}

func (h *tabunganHandler) GetTabungan(c *fiber.Ctx) error {
	NoRekening := c.Params("no_rekening")
	noRekening, err := strconv.Atoi(NoRekening)
	if err != nil {
		response := helper.APIResponse("Failed to get URI", http.StatusBadRequest, "Input Data URI Required", noRekening)
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	tabunganDetail, err := h.service.GetTabunganByNoRek(noRekening)
	if err != nil {
		response := helper.APIResponse("Failed to get detail Tabungan", http.StatusBadRequest, "Nomor Rekening Not Available", err)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("Detail Tabungan", http.StatusOK, "success", tabungan.FormatTabunganDetail(tabunganDetail))
	return c.Status(http.StatusOK).JSON(response)
}