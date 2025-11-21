package service

import (
	"Magazine_backend/internal/database"
	"Magazine_backend/internal/models"
	"Magazine_backend/pkg/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary		Get purchases
// @Security		ApiKeyAuth
// @Description	Retrieve a list of purchases for a specific account by account ID with pagination
// @Tags			purchases
// @Accept			json
// @Produce		json
// @Param			page	query		int		false	"Page number (default: 1)"
// @Param			limit	query		int		false	"Number of purchases per page (default: 5)"
// @Success		200		{array}		models.Purchase
// @Failure		400		{object}	errorResponse
// @Failure		404		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/v1/purchases [get]
func GetPurchases(c *gin.Context) {
	// Получение параметров из запроса
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))   // Номер страницы, по умолчанию 1
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5")) // Количество элементов на странице, по умолчанию 5

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}

	offset := (page - 1) * limit // Вычисление смещения

	var mans []models.Purchase

	// Использование GORM для выборки с лимитом и смещением
	err := database.DbPostgres.Limit(limit).Offset(offset).Preload("Supplier").Preload("Items").Find(&mans).Error
	if err != nil {
		utils.Logger.Error("Неудачный запрос|(purchases_handler.go|GetPurchases|):", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, mans)

	utils.Logger.Printf("GetPurchases сработал: %v", mans)
}

// @Summary		Get purchase by ID
// @Security		ApiKeyAuth
// @Description	Retrieve a specific purchase by its ID
// @Tags			purchases
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Account ID"
// @Success		200	{object}	models.Purchase
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/purchases/{id} [get]
func GetPurchaseById(c *gin.Context) {
	// Получаем параметр id из запроса
	id := c.Param("id")

	// Использование GORM для поиска профиля по ID
	var purchase models.Purchase
	err := database.DbPostgres.Preload("Supplier").Preload("Items").First(&purchase, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "purchase not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(purchases_handler.go|GetPurchaseById|):", err)
		return
	}

	// Возвращаем профиль в ответе
	c.JSON(http.StatusOK, purchase)
}

// @Summary		Create a new purchase
// @Description	Creates a new purchase by accepting purchase details in the request body
// @Tags			sign
// @Accept			json
// @Produce		json
// @Param			purchase	body		models.Purchase	true	"Purchase data"
// @Success		201	{object}	models.Purchase
// @Failure		400	{object}	errorResponse
// @Failure		409	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/purchases [post]
func CreatePurchase(c *gin.Context) {
	req := models.CreatePurchase{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(purchases_handler.go|CreatePurchase|)|:", err)
		return
	}

	purchase := models.Purchase{
		SupplierID:    req.SupplierID,
		TotalAmount:   req.TotalAmount,
		InvoiceNumber: req.InvoiceNumber,
	}

	if err := database.DbPostgres.Create(&purchase).Error; err != nil {
		utils.Logger.Error("Insert isn't done(purchases_handler.go|CreatePurchase|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	if err := database.DbPostgres.Preload("Supplier").Preload("Items").Last(&purchase).Error; err != nil {
		utils.Logger.Error("Сan't receive purchase(purchases_handler.go|CreatePurchase|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		database.DbPostgres.Delete(&purchase)
		return
	}
	c.JSON(http.StatusCreated, purchase)
}

// @Summary		Update an existing purchase
// @Security		ApiKeyAuth
// @Description	Update an existing purchase's information by purchase ID
// @Tags			purchases
// @Accept			json
// @Produce		json
// @Param			id		path		int			true	"Purchase ID"
// @Param			purchase	body		models.Purchase	true	"Updated purchase data"
// @Success		202	{object}	models.Purchase
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/purchases/{id} [put]
func UpdatePurchase(c *gin.Context) {
	id := c.Param("id")
	req := models.UpdatePurchase{}

	// Парсим JSON из тела запроса
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(purchases_handler.go|UpdatePurchase|)|:", err)
		return
	}

	// Проверка, есть ли такой id
	var existing models.Purchase
	if err := database.DbPostgres.Where("id = ?", id).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "purchase not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		utils.Logger.Error("DB error when checking purchase (purchases_handler.go|UpdatePurchase):", err)
		return
	}

	// Обновляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Model(&models.Purchase{}).Where("id = ?", id).Updates(req).Error; err != nil {
		utils.Logger.Error("Update isn't done(purchases_handler.go|UpdatePurchase|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Использование GORM для поиска профиля по ID
	var purchase models.Purchase

	if err := database.DbPostgres.Preload("Supplier").Preload("Items").First(&purchase, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "purchase not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(purchases_handler.go|UpdatePurchase|):", err)
		return
	}

	// Отправляем успешный ответ с обновленным профилем
	c.JSON(http.StatusAccepted, purchase)
}

// @Summary		Delete a purchase by ID
// @Security		ApiKeyAuth
// @Description	Delete a purchase from the system by its ID
// @Tags			purchase
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"Purchase ID"
// @Success		202	{object}	string
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/purchases/{id} [delete]
func DeletePurchase(c *gin.Context) {
	id := c.Param("id")

	// Удаляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Delete(&models.Purchase{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "purchase wasn't deleted"})
		utils.Logger.Error("Delete isn't done(purchases_handler.go|DeletePurchase|):", err)
		return
	}

	// Отправляем успешный ответ о удалении
	c.JSON(http.StatusAccepted, gin.H{"message": "purchase was deleted"})
}
