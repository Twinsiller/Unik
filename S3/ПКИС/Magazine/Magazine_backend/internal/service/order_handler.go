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

func GetOrders(c *gin.Context) {
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

	var orders []models.Order

	err := database.DbPostgres.Limit(limit).Offset(offset).Preload("Cashier").Find(&orders).Error
	if err != nil {
		utils.Logger.Error("Неудачный запрос|(order_handler.go|GetOrders|):", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, orders)

	utils.Logger.Printf("GetOrders сработал: %v", orders)
}

func GetOrderById(c *gin.Context) {
	id := c.Param("id")

	var order models.Order

	err := database.DbPostgres.Preload("Cashier").First(&order, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(order_handler.go|GetOrderById|):", err)
		return
	}

	c.JSON(http.StatusOK, order)
}

func CreateOrder(c *gin.Context) {
	print("Yep")
	req := models.CreateOrder{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(order_handler.go|CreateOrder|)|:", err)
		return
	}
	print("Yep2")
	order := models.Order{
		CashierID:     req.CashierID,
		TotalAmount:   req.TotalAmount,
		PaymentMethod: req.PaymentMethod,
		ShiftID:       req.ShiftID,
	}

	if err := database.DbPostgres.Create(&order).Error; err != nil {
		utils.Logger.Error("Insert isn't done(order_handler.go|CreateOrder|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	if err := database.DbPostgres.Preload("Cashier").Last(&order).Error; err != nil {
		utils.Logger.Error("Сan't receive order(order_handler.go|CreateOrder|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		database.DbPostgres.Delete(&order)
		return
	}
	c.JSON(http.StatusCreated, order)
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	req := models.UpdateOrder{}

	// Парсим JSON из тела запроса
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(order_handler.go|UpdateOrder|)|:", err)
		return
	}

	// Проверка, есть ли такой id
	// existing := models.Order{}
	var existing models.Order
	if err := database.DbPostgres.Where("id = ?", id).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		utils.Logger.Error("DB error when checking order (order_handler.go|UpdateOrder):", err)
		return
	}

	// Обновляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Model(&models.Order{}).Where("id = ?", id).Updates(req).Error; err != nil {
		utils.Logger.Error("Update isn't done(order_handler.go|UpdateOrder|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	var order models.Order

	err := database.DbPostgres.Preload("Cashier").First(&order, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(order_handler.go|UpdateOrder|):", err)
		return
	}

	// Отправляем успешный ответ с обновленным профилем
	c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	if err := database.DbPostgres.Where("id = ?", id).Delete(&models.Order{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order wasn't deleted"})
		utils.Logger.Error("Delete isn't done(order_handler.go|DeleteOrder|):", err)
		return
	}

	// Отправляем успешный ответ о удалении
	c.JSON(http.StatusAccepted, gin.H{"message": "order was deleted"})
}
