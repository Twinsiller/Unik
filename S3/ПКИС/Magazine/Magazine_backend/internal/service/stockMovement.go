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

// @Summary		Get stockMovements
// @Security		ApiKeyAuth
// @Description	Retrieve a list of stockMovements for a specific account by account ID with pagination
// @Tags			stockMovements
// @Accept			json
// @Produce		json
// @Param			page	query		int		false	"Page number (default: 1)"
// @Param			limit	query		int		false	"Number of stockMovements per page (default: 5)"
// @Success		200		{array}		models.StockMovement
// @Failure		400		{object}	errorResponse
// @Failure		404		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/v1/stockMovements [get]
func GetStockMovements(c *gin.Context) {
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

	var stockMovements []models.StockMovement

	// Использование GORM для выборки с лимитом и смещением
	err := database.DbPostgres.Limit(limit).Offset(offset).Preload("Actor").Preload("Stock").Preload("Product").Preload("Order").Preload("Purchase").Find(&stockMovements).Error
	if err != nil {
		utils.Logger.Error("Неудачный запрос|(Movement.go|GetProducts|):", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, stockMovements)

	utils.Logger.Printf("GetProducts сработал: %v", stockMovements)
}

// @Summary		Get stockMovement by ID
// @Security		ApiKeyAuth
// @Description	Retrieve a specific stockMovement by its ID
// @Tags			stockMovements
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"StockMovement ID"
// @Success		200	{object}	models.StockMovement
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/stockMovements/{id} [get]
func GetStockMovementById(c *gin.Context) {
	// Получаем параметр id из запроса
	id := c.Param("id")

	// Использование GORM для поиска профиля по ID
	var stockMovement models.StockMovement
	err := database.DbPostgres.Preload("Actor").Preload("Stock").Preload("Product").Preload("Order").Preload("Purchase").First(&stockMovement, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "stockMovement not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(Movement.go|GetProductById|):", err)
		return
	}

	// Возвращаем профиль в ответе
	c.JSON(http.StatusOK, stockMovement)
}

// @Summary		Create a new stockMovement
// @Description	Creates a new stockMovement by accepting stockMovement details in the request body
// @Tags			sign
// @Accept			json
// @Produce		json
// @Param			stockMovement	body		models.StockMovement	true	"StockMovement data"
// @Success		201	{object}	models.StockMovement
// @Failure		400	{object}	errorResponse
// @Failure		409	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/stockMovements [post]
func CreateStockMovement(c *gin.Context) {
	req := models.CreateStockMovement{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(Movement.go|CreateStockMovement|)|:", err)
		return
	}

	stockMovement := models.StockMovement{
		ActorID:           req.ActorID,
		StockID:           req.StockID,
		ProductID:         req.ProductID,
		RelatedOrderID:    req.RelatedOrderID,
		RelatedPurchaseID: req.RelatedPurchaseID,
		Unit:              req.Unit,
		Change:            req.Change,
		Reason:            req.Reason,
	}

	temp := req.RelatedOrderID != nil || req.RelatedPurchaseID != nil
	if err := database.DbPostgres.Create(&stockMovement).Error; err != nil && temp {
		utils.Logger.Error("Insert isn't done(Movement.go|CreateStockMovement|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	if err := database.DbPostgres.Preload("Actor").Preload("Stock").Preload("Product").Preload("Order").Preload("Purchase").Last(&stockMovement).Error; err != nil {
		utils.Logger.Error("Сan't receive stockMovement(Movement.go|CreateStockMovement|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		database.DbPostgres.Delete(&stockMovement)
		return
	}
	c.JSON(http.StatusCreated, stockMovement)
}

// @Summary		Update an existing stockMovement
// @Security		ApiKeyAuth
// @Description	Update an existing stockMovement's information by stockMovement ID
// @Tags			stockMovements
// @Accept			json
// @Produce		json
// @Param			id		path		int			true	"StockMovement ID"
// @Param			stockMovement	body		models.StockMovement	true	"Updated stockMovement data"
// @Success		202	{object}	models.StockMovement
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/stockMovements/{id} [put]
func UpdateStockMovement(c *gin.Context) {
	id := c.Param("id")
	req := models.UpdateStockMovement{}

	// Парсим JSON из тела запроса
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(Movement.go|e|)|:", err)
		return
	}

	// Проверка, есть ли такой id
	var existing models.StockMovement
	if err := database.DbPostgres.Where("id = ?", id).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "stockMovement not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		utils.Logger.Error("DB error when checking stockMovement (Movement.go|e):", err)
		return
	}

	// Обновляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Model(&models.StockMovement{}).Where("id = ?", id).Updates(req).Error; err != nil {
		utils.Logger.Error("Update isn't done(Movement.go|e|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Использование GORM для поиска профиля по ID
	var stockMovement models.StockMovement

	if err := database.DbPostgres.Preload("Actor").Preload("Stock").Preload("Product").Preload("Order").Preload("Purchase").First(&stockMovement, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "stockMovement not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(Movement.go|e|):", err)
		return
	}

	// Отправляем успешный ответ с обновленным профилем
	c.JSON(http.StatusAccepted, stockMovement)
}

// @Summary		Delete a stockMovement by ID
// @Security		ApiKeyAuth
// @Description	Delete a stockMovement from the system by its ID
// @Tags			stockMovement
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"StockMovement ID"
// @Success		202	{object}	string
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/stockMovements/{id} [delete]
func DeleteStockMovement(c *gin.Context) {
	id := c.Param("id")

	// Удаляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Delete(&models.StockMovement{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "stockMovement wasn't deleted"})
		utils.Logger.Error("Delete isn't done(stockMovement_handler.go|DeleteStockMovement|):", err)
		return
	}

	// Отправляем успешный ответ о удалении
	c.JSON(http.StatusAccepted, gin.H{"message": "stockMovement was deleted"})
}
