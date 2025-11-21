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

// @Summary		Get stocks
// @Security		ApiKeyAuth
// @Description	Retrieve a list of stocks for a specific account by account ID with pagination
// @Tags			stocks
// @Accept			json
// @Produce		json
// @Param			page	query		int		false	"Page number (default: 1)"
// @Param			limit	query		int		false	"Number of stocks per page (default: 5)"
// @Success		200		{array}		models.Stock
// @Failure		400		{object}	errorResponse
// @Failure		404		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/v1/stocks [get]
func GetStocks(c *gin.Context) {
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

	var stocks []models.Stock

	// Использование GORM для выборки с лимитом и смещением
	err := database.DbPostgres.Limit(limit).Offset(offset).Preload("Product").Preload("StockMovements").Preload("OrderItems").Preload("PurchaseItems").Preload("Stocks").Preload("StockMovements").Find(&stocks).Error
	if err != nil {
		utils.Logger.Error("Неудачный запрос|(stock_handler.go|GetProducts|):", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, stocks)

	utils.Logger.Printf("GetProducts сработал: %v", stocks)
}

// @Summary		Get stock by ID
// @Security		ApiKeyAuth
// @Description	Retrieve a specific stock by its ID
// @Tags			stocks
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Stock ID"
// @Success		200	{object}	models.Stock
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/stocks/{id} [get]
func GetStockById(c *gin.Context) {
	// Получаем параметр id из запроса
	id := c.Param("id")

	// Использование GORM для поиска профиля по ID
	var stock models.Stock
	err := database.DbPostgres.Preload("Product").Preload("StockMovements").Preload("OrderItems").Preload("PurchaseItems").Preload("Stocks").Preload("StockMovements").First(&stock, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "stock not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(stock_handler.go|GetStockById|):", err)
		return
	}

	// Возвращаем профиль в ответе
	c.JSON(http.StatusOK, stock)
}

// @Summary		Create a new stock
// @Description	Creates a new stock by accepting stock details in the request body
// @Tags			sign
// @Accept			json
// @Produce		json
// @Param			stock	body		models.Stock	true	"Stock data"
// @Success		201	{object}	models.Stock
// @Failure		400	{object}	errorResponse
// @Failure		409	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/stocks [post]
func CreateStock(c *gin.Context) {
	req := models.CreateStock{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(stock_handler.go|CreateStock|)|:", err)
		return
	}

	stock := models.Stock{
		ProductID:     req.ProductID,
		Quantity:      req.Quantity,
		BatchNumber:   req.BatchNumber,
		PurchasePrice: req.PurchasePrice,
		ReceivedAt:    req.ReceivedAt,
		ExpiryDate:    req.ExpiryDate,
	}

	if err := database.DbPostgres.Create(&stock).Error; err != nil {
		utils.Logger.Error("Insert isn't done(stock_handler.go|CreateStock|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	if err := database.DbPostgres.Preload("Product").Preload("StockMovements").Preload("OrderItems").Preload("PurchaseItems").Preload("Stocks").Preload("StockMovements").Last(&stock).Error; err != nil {
		utils.Logger.Error("Сan't receive stock(stock_handler.go|CreateStock|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		database.DbPostgres.Delete(&stock)
		return
	}
	c.JSON(http.StatusCreated, stock)
}

// @Summary		Update an existing stock
// @Security		ApiKeyAuth
// @Description	Update an existing stock's information by stock ID
// @Tags			stocks
// @Accept			json
// @Produce		json
// @Param			id		path		int			true	"Stock ID"
// @Param			stock	body		models.Stock	true	"Updated stock data"
// @Success		202	{object}	models.Stock
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/stocks/{id} [put]
func UpdateStock(c *gin.Context) {
	id := c.Param("id")
	req := models.UpdateStock{}

	// Парсим JSON из тела запроса
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(stock_handler.go|UpdateStock|)|:", err)
		return
	}

	// Проверка, есть ли такой id
	var existing models.Stock
	if err := database.DbPostgres.Where("id = ?", id).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "stock not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		utils.Logger.Error("DB error when checking stock (stock_handler.go|UpdateStock):", err)
		return
	}

	// Обновляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Model(&models.Stock{}).Where("id = ?", id).Updates(req).Error; err != nil {
		utils.Logger.Error("Update isn't done(stock_handler.go|UpdateStock|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Использование GORM для поиска профиля по ID
	var stock models.Stock

	if err := database.DbPostgres.Preload("Product").Preload("StockMovements").Preload("OrderItems").Preload("PurchaseItems").Preload("Stocks").Preload("StockMovements").First(&stock, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "stock not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(stock_handler.go|UpdateStock|):", err)
		return
	}

	// Отправляем успешный ответ с обновленным профилем
	c.JSON(http.StatusAccepted, stock)
}

// @Summary		Delete a stock by ID
// @Security		ApiKeyAuth
// @Description	Delete a stock from the system by its ID
// @Tags			stock
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"Stock ID"
// @Success		202	{object}	string
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/stocks/{id} [delete]
func DeleteStock(c *gin.Context) {
	id := c.Param("id")

	// Удаляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Delete(&models.Stock{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "stock wasn't deleted"})
		utils.Logger.Error("Delete isn't done(stock_handler.go|DeleteProduct|):", err)
		return
	}

	// Отправляем успешный ответ о удалении
	c.JSON(http.StatusAccepted, gin.H{"message": "stock was deleted"})
}
