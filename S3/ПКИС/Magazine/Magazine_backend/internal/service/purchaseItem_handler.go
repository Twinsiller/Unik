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

// @Summary		Get purchaseItems
// @Security		ApiKeyAuth
// @Description	Retrieve a list of purchaseItems for a specific account by account ID with pagination
// @Tags			purchaseItems
// @Accept			json
// @Produce		json
// @Param			page	query		int		false	"Page number (default: 1)"
// @Param			limit	query		int		false	"Number of purchaseItems per page (default: 5)"
// @Success		200		{array}		models.PurchaseItem
// @Failure		400		{object}	errorResponse
// @Failure		404		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/v1/purchaseItems [get]
func GetPurchaseItems(c *gin.Context) {
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

	var purchaseItems []models.PurchaseItem

	// Использование GORM для выборки с лимитом и смещением
	err := database.DbPostgres.Limit(limit).Offset(offset).Preload("Product").Preload("Purchase").Preload("OrderItems").Preload("PurchaseItems").Preload("Stocks").Preload("StockMovements").Find(&purchaseItems).Error
	if err != nil {
		utils.Logger.Error("Неудачный запрос|(purchaseItem_handler.go|GetPurchaseItems|):", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, purchaseItems)

	utils.Logger.Printf("GetPurchaseItems сработал: %v", purchaseItems)
}

// @Summary		Get purchaseItem by ID
// @Security		ApiKeyAuth
// @Description	Retrieve a specific purchaseItem by its ID
// @Tags			purchaseItems
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"PurchaseItem ID"
// @Success		200	{object}	models.PurchaseItem
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/purchaseItems/{id} [get]
func GetPurchaseItemById(c *gin.Context) {
	// Получаем параметр id из запроса
	id := c.Param("id")

	// Использование GORM для поиска профиля по ID
	var purchaseItem models.PurchaseItem
	err := database.DbPostgres.Preload("Product").Preload("Purchase").Preload("OrderItems").Preload("PurchaseItems").Preload("Stocks").Preload("StockMovements").First(&purchaseItem, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "purchaseItem not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(purchaseItem_handleurchaseItem|):", err)
		return
	}

	// Возвращаем профиль в ответе
	c.JSON(http.StatusOK, purchaseItem)
}

// @Summary		Create a new purchaseItem
// @Description	Creates a new purchaseItem by accepting purchaseItem details in the request body
// @Tags			sign
// @Accept			json
// @Produce		json
// @Param			purchaseItem	body		models.PurchaseItem	true	"PurchaseItem data"
// @Success		201	{object}	models.PurchaseItem
// @Failure		400	{object}	errorResponse
// @Failure		409	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/purchaseItems [post]
func CreatePurchaseItem(c *gin.Context) {
	req := models.CreatePurchaseItem{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(purchaseItem_handler.go|CreatePurchaseItem|)|:", err)
		return
	}

	purchaseItem := models.PurchaseItem{
		PurchaseID:    req.PurchaseID,
		ProductID:     req.ProductID,
		Quantity:      req.Quantity,
		PurchasePrice: req.PurchasePrice,
		ExpiryDate:    req.ExpiryDate,
		BatchNumber:   req.BatchNumber,
	}

	if err := database.DbPostgres.Create(&purchaseItem).Error; err != nil {
		utils.Logger.Error("Insert isn't done(purchaseItem_handler.go|CreatePurchaseItem|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	if err := database.DbPostgres.Preload("Product").Preload("Purchase").Preload("OrderItems").Preload("PurchaseItems").Preload("Stocks").Preload("StockMovements").Last(&purchaseItem).Error; err != nil {
		utils.Logger.Error("Сan't receive purchaseItem(purchaseItem_handler.go|CreatePurchaseItem|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		database.DbPostgres.Delete(&purchaseItem)
		return
	}
	c.JSON(http.StatusCreated, purchaseItem)
}

// @Summary		Update an existing purchaseItem
// @Security		ApiKeyAuth
// @Description	Update an existing purchaseItem's information by purchaseItem ID
// @Tags			purchaseItems
// @Accept			json
// @Produce		json
// @Param			id		path		int			true	"PurchaseItem ID"
// @Param			purchaseItem	body		models.PurchaseItem	true	"Updated purchaseItem data"
// @Success		202	{object}	models.PurchaseItem
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/purchaseItems/{id} [put]
func UpdatePurchaseItem(c *gin.Context) {
	id := c.Param("id")
	req := models.UpdatePurchaseItem{}

	// Парсим JSON из тела запроса
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(purchaseItem_handler.go|UpdatePurchaseItem|)|:", err)
		return
	}

	// Проверка, есть ли такой id
	var existing models.PurchaseItem
	if err := database.DbPostgres.Where("id = ?", id).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "purchaseItem not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		utils.Logger.Error("DB error when checking purchaseItem (purchaseItem_handler.go|UpdatePurchaseItem):", err)
		return
	}

	// Обновляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Model(&models.PurchaseItem{}).Where("id = ?", id).Updates(req).Error; err != nil {
		utils.Logger.Error("Update isn't done(purchaseItem_handler.go|UpdatePurchaseItem|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Использование GORM для поиска профиля по ID
	var purchaseItem models.PurchaseItem

	if err := database.DbPostgres.Preload("Product").Preload("Purchase").Preload("OrderItems").Preload("PurchaseItems").Preload("Stocks").Preload("StockMovements").First(&purchaseItem, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "purchaseItem not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(purchaseItem_handler.go|UpdatePurchaseItem|):", err)
		return
	}

	// Отправляем успешный ответ с обновленным профилем
	c.JSON(http.StatusAccepted, purchaseItem)
}

// @Summary		Delete a purchaseItem by ID
// @Security		ApiKeyAuth
// @Description	Delete a purchaseItem from the system by its ID
// @Tags			purchaseItem
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"PurchaseItem ID"
// @Success		202	{object}	string
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/purchaseItems/{id} [delete]
func DeletePurchaseItem(c *gin.Context) {
	id := c.Param("id")

	// Удаляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Delete(&models.PurchaseItem{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "purchaseItem wasn't deleted"})
		utils.Logger.Error("Delete isn't done(purchaseItem_handler.go|DeletePurchaseItem|):", err)
		return
	}

	// Отправляем успешный ответ о удалении
	c.JSON(http.StatusAccepted, gin.H{"message": "purchaseItem was deleted"})
}
