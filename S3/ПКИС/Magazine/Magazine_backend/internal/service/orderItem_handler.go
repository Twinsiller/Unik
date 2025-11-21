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

// @Summary		Get orderItems
// @Security		ApiKeyAuth
// @Description	Retrieve a list of orderItems for a specific account by account ID with pagination
// @Tags			orderItems
// @Accept			json
// @Produce		json
// @Param			page	query		int		false	"Page number (default: 1)"
// @Param			limit	query		int		false	"Number of orderItems per page (default: 5)"
// @Success		200		{array}		models.OrderItem
// @Failure		400		{object}	errorResponse
// @Failure		404		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/v1/orderItems [get]
func GetOrderItems(c *gin.Context) {
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

	var orderItems []models.OrderItem

	// Использование GORM для выборки с лимитом и смещением
	err := database.DbPostgres.Limit(limit).Offset(offset).Preload("Order").Preload("Product").Find(&orderItems).Error
	if err != nil {
		utils.Logger.Error("Неудачный запрос|(orderItem_handler.go|GetOrderItems|):", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, orderItems)

	utils.Logger.Printf("GetCategoriess сработал: %v", orderItems)
}

// @Summary		Get orderItem by ID
// @Security		ApiKeyAuth
// @Description	Retrieve a specific orderItem by its ID
// @Tags			manufactures
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"OrderItem ID"
// @Success		200	{object}	models.OrderItem
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/manufactures/{id} [get]
func GetOrderItemById(c *gin.Context) {
	// Получаем параметр id из запроса
	id := c.Param("id")

	// Использование GORM для поиска профиля по ID
	var orderItem models.OrderItem
	err := database.DbPostgres.Preload("Order").Preload("Product").First(&orderItem, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "orderItem not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(orderItem_handler.go|GetOrderItemById|):", err)
		return
	}

	// Возвращаем профиль в ответе
	c.JSON(http.StatusOK, orderItem)
}

// @Summary		Create a new orderItem
// @Description	Creates a new orderItem by accepting orderItem details in the request body
// @Tags			sign
// @Accept			json
// @Produce		json
// @Param			orderItem	body		models.OrderItem	true	"OrderItem data"
// @Success		201	{object}	models.OrderItem
// @Failure		400	{object}	errorResponse
// @Failure		409	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/orderitems [post]
func CreateOrderItem(c *gin.Context) {
	req := models.CreateOrderItem{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(orderItem_handler.go|CreateOrderItem|)|:", err)
		return
	}

	orderItem := models.OrderItem{
		OrderID:      req.OrderID,
		ProductID:    req.ProductID,
		Quantity:     req.Quantity,
		PricePerUnit: req.PricePerUnit,
		Discount:     req.Discount,
		VatRate:      req.VatRate,
	}

	if err := database.DbPostgres.Create(&orderItem).Error; err != nil {
		utils.Logger.Error("Insert isn't done(orderItem_handler.go|CreateOrderItem|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	if err := database.DbPostgres.Preload("Order").Preload("Product").Last(&orderItem).Error; err != nil {
		utils.Logger.Error("Сan't receive orderItem(orderItem_handler.go|CreateOrderItem|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		database.DbPostgres.Delete(&orderItem)
		return
	}
	c.JSON(http.StatusCreated, orderItem)
}

// @Summary		Update an existing orderItem
// @Security		ApiKeyAuth
// @Description	Update an existing orderItem's information by orderItem ID
// @Tags			orderitems
// @Accept			json
// @Produce		json
// @Param			id		path		int			true	"OrderItem ID"
// @Param			orderItem	body		models.OrderItem	true	"Updated orderItem data"
// @Success		202	{object}	models.OrderItem
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/orderitems/{id} [put]
func UpdateOrderItem(c *gin.Context) {
	id := c.Param("id")
	req := models.UpdateOrderItem{}

	// Парсим JSON из тела запроса
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(orderItem_handler.go|UpdateOrderItem|)|:", err)
		return
	}

	// Проверка, есть ли такой id
	var existing models.OrderItem
	if err := database.DbPostgres.Where("id = ?", id).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "orderItem not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		utils.Logger.Error("DB error when checking orderItem (orderItem_handler.go|UpdateOrderItem):", err)
		return
	}

	// Обновляем категорию по ID с использованием GORM
	if err := database.DbPostgres.Model(&models.OrderItem{}).Where("id = ?", id).Updates(req).Error; err != nil {
		utils.Logger.Error("Update isn't done(orderItem_handler.go|UpdateOrderItem|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Использование GORM для поиска профиля по ID
	var orderItem models.OrderItem

	if err := database.DbPostgres.Preload("Order").Preload("Product").First(&orderItem, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "orderItem not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(orderItem_handler.go|UpdateOrderItem|):", err)
		return
	}

	// Отправляем успешный ответ с обновленным профилем
	c.JSON(http.StatusAccepted, orderItem)
}

// @Summary		Delete a orderItem by ID
// @Security		ApiKeyAuth
// @Description	Delete a orderItem from the system by its ID
// @Tags			orderitems
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"OrderItem ID"
// @Success		202	{object}	string
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/orderItems/{id} [delete]
func DeleteOrderItem(c *gin.Context) {
	id := c.Param("id")

	// Удаляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Delete(&models.OrderItem{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "orderItem wasn't deleted"})
		utils.Logger.Error("Delete isn't done(orderItem_handler.go|DeleteOrderItem|):", err)
		return
	}

	// Отправляем успешный ответ о удалении
	c.JSON(http.StatusAccepted, gin.H{"message": "orderItem was deleted"})
}
