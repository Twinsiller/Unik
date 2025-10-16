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

	err := database.DbPostgres.Limit(limit).Offset(offset).Find(&orders).Error
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

	err := database.DbPostgres.First(&order, id).Error
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
	req := models.Order{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(order_handler.go|CreateOrder|)|:", err)
		return
	}

	if err := database.DbPostgres.Create(&req).Error; err != nil {
		utils.Logger.Error("Insert isn't done(order_handler.go|CreateOrder|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, req)
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	req := models.Order{}

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
			c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		utils.Logger.Error("DB error when checking user (user_handler.go|UpdateUser):", err)
		return
	}

	p := models.User{
		Nickname:     req.Nickname,
		HashPassword: hash,
		Role:         req.Role,
		Name:         req.Name,
	}

	// Обновляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Model(&models.User{}).Where("id = ?", id).Updates(p).Error; err != nil {
		utils.Logger.Error("Update isn't done(user_handler.go|UpdateUser|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Использование GORM для поиска профиля по ID
	var user models.User

	if err := database.DbPostgres.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(user_handler.go|GetUserById|):", err)
		return
	}

	// Отправляем успешный ответ с обновленным профилем
	c.JSON(http.StatusAccepted, user)
}
