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

// @Summary		Get manufacturers
// @Security		ApiKeyAuth
// @Description	Retrieve a list of manufacturers for a specific account by account ID with pagination
// @Tags			manufacturers
// @Accept			json
// @Produce		json
// @Param			page	query		int		false	"Page number (default: 1)"
// @Param			limit	query		int		false	"Number of manufacturers per page (default: 5)"
// @Success		200		{array}		models.Manufacturer
// @Failure		400		{object}	errorResponse
// @Failure		404		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/v1/manufacturers [get]
func GetManufacturers(c *gin.Context) {
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

	var mans []models.Manufacturer

	// Использование GORM для выборки с лимитом и смещением
	err := database.DbPostgres.Limit(limit).Offset(offset).Preload("Products").Find(&mans).Error
	if err != nil {
		utils.Logger.Error("Неудачный запрос|(manufacturer_handler.go|GetUsers|):", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, mans)

	utils.Logger.Printf("GetManufacturers сработал: %v", mans)
}

// @Summary		Get manufacturer by ID
// @Security		ApiKeyAuth
// @Description	Retrieve a specific manufacturer by its ID
// @Tags			manufactures
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Account ID"
// @Success		200	{object}	models.Manufacturer
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/manufactures/{id} [get]
func GetManufacturerById(c *gin.Context) {
	// Получаем параметр id из запроса
	id := c.Param("id")

	// Использование GORM для поиска профиля по ID
	var man models.Manufacturer
	err := database.DbPostgres.Preload("Products").First(&man, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "manufacturer not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(manufacturer_handler.go|GetManufactureById|):", err)
		return
	}

	// Возвращаем профиль в ответе
	c.JSON(http.StatusOK, man)
}

// @Summary		Create a new manufacturer
// @Description	Creates a new manufacturer by accepting manufacturer details in the request body
// @Tags			sign
// @Accept			json
// @Produce		json
// @Param			manufacturer	body		models.Manufacturer	true	"Manufacturer data"
// @Success		201	{object}	models.Manufacturer
// @Failure		400	{object}	errorResponse
// @Failure		409	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/manufactures [post]
func CreateManufacturer(c *gin.Context) {
	req := models.CreateManufacturer{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(manufacturer_handler.go|CreateManufacturer|)|:", err)
		return
	}

	man := models.Manufacturer{
		Name:        req.Name,
		Country:     req.Country,
		ContactInfo: req.ContactInfo,
	}

	if err := database.DbPostgres.Create(&man).Error; err != nil {
		utils.Logger.Error("Insert isn't done(manufacturer_handler.go|CreateManufacturer|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	if err := database.DbPostgres.Preload("Products").Last(&man).Error; err != nil {
		utils.Logger.Error("Сan't receive manufacturer(manufacturer_handler.go|CreateManufacturer|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		database.DbPostgres.Delete(&man)
		return
	}
	c.JSON(http.StatusCreated, man)
}

// @Summary		Update an existing manufacturer
// @Security		ApiKeyAuth
// @Description	Update an existing manufacturer's information by manufacturer ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id		path		int			true	"Manufacturer ID"
// @Param			manufacturer	body		models.UsManufacturerer	true	"Updated manufacturer data"
// @Success		202	{object}	models.Manufacturer
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/users/{id} [put]
func UpdateManufacturer(c *gin.Context) {
	id := c.Param("id")
	req := models.UpdateManufacturer{}

	// Парсим JSON из тела запроса
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(manufacturer_handler.go|UpdateManufacturer|)|:", err)
		return
	}

	// Проверка, есть ли такой id
	var existing models.Manufacturer
	if err := database.DbPostgres.Where("id = ?", id).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "manufacturer not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		utils.Logger.Error("DB error when checking manufacturer (manufacturer_handler.go|UpdateManufacturer):", err)
		return
	}

	// Обновляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Model(&models.Manufacturer{}).Where("id = ?", id).Updates(req).Error; err != nil {
		utils.Logger.Error("Update isn't done(manufacturer_handler.go|UpdateManufacturer|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Использование GORM для поиска профиля по ID
	var manufacturer models.Manufacturer

	if err := database.DbPostgres.Preload("Products").First(&manufacturer, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "manufacturer not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(manufacturer_handler.go|UpdateManufacturer|):", err)
		return
	}

	// Отправляем успешный ответ с обновленным профилем
	c.JSON(http.StatusAccepted, manufacturer)
}

// @Summary		Delete a manufacturer by ID
// @Security		ApiKeyAuth
// @Description	Delete a manufacturer from the system by its ID
// @Tags			manufacturer
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"Manufacturer ID"
// @Success		202	{object}	string
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/manufacturers/{id} [delete]
func DeleteManufacturer(c *gin.Context) {
	id := c.Param("id")

	// Удаляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Delete(&models.Manufacturer{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "manufacturer wasn't deleted"})
		utils.Logger.Error("Delete isn't done(manufacturer_handler.go|DeleteManufacturer|):", err)
		return
	}

	// Отправляем успешный ответ о удалении
	c.JSON(http.StatusAccepted, gin.H{"message": "manufacturer was deleted"})
}
