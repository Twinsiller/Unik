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

func GetSuppliers(c *gin.Context) {
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

	var sups []models.Supplier

	err := database.DbPostgres.Limit(limit).Offset(offset).Preload("Purchases").Find(&sups).Error
	if err != nil {
		utils.Logger.Error("Неудачный запрос|(supply_handler.go|GetSuppliers|):", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, sups)

	utils.Logger.Printf("GetSuppliers сработал: %v", sups)
}

func GetSupplierById(c *gin.Context) {
	id := c.Param("id")

	var sup models.Supplier

	err := database.DbPostgres.Preload("Purchases").First(&sup, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "sup not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(supply_handler.go|GetSupplierById|):", err)
		return
	}

	c.JSON(http.StatusOK, sup)
}

func CreateSupplier(c *gin.Context) {
	req := models.CreateSupplier{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(supply_handler.go|CreateSupplier|)|:", err)
		return
	}

	sup := models.Supplier{
		Name:        req.Name,
		ContactInfo: req.ContactInfo,
		Preferred:   req.Preferred,
	}

	if err := database.DbPostgres.Create(&sup).Error; err != nil {
		utils.Logger.Error("Insert isn't done(supply_handler.go|CreateSupplier|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	if err := database.DbPostgres.Preload("Purchases").Last(&sup).Error; err != nil {
		utils.Logger.Error("Сan't receive sup(supply_handler.go|CreateSupplier|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		database.DbPostgres.Delete(&sup)
		return
	}
	c.JSON(http.StatusCreated, sup)
}

func UpdateSupplier(c *gin.Context) {
	id := c.Param("id")

	req := models.UpdateSupplier{}

	// Парсим JSON из тела запроса
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(supply_handler.go|UpdateSupplier|)|:", err)
		return
	}

	// Проверка, есть ли такой id
	// existing := models.Supplier{}
	var existing models.Supplier
	if err := database.DbPostgres.Where("id = ?", id).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "sup not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		utils.Logger.Error("DB error when checking sup (supply_handler.go|UpdateSupplier):", err)
		return
	}

	// Обновляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Model(&models.Supplier{}).Where("id = ?", id).Updates(req).Error; err != nil {
		utils.Logger.Error("Update isn't done(supply_handler.go|UpdateSupplier|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	var sup models.Supplier

	err := database.DbPostgres.Preload("Purchases").First(&sup, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "sup not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(supply_handler.go|UpdateSupplier|):", err)
		return
	}

	// Отправляем успешный ответ с обновленным профилем
	c.JSON(http.StatusOK, sup)
}

func DeleteSupplier(c *gin.Context) {
	id := c.Param("id")

	if err := database.DbPostgres.Where("id = ?", id).Delete(&models.Supplier{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "sup wasn't deleted"})
		utils.Logger.Error("Delete isn't done(supply_handler.go|DeleteSupplier|):", err)
		return
	}

	// Отправляем успешный ответ о удалении
	c.JSON(http.StatusAccepted, gin.H{"message": "sup was deleted"})
}
