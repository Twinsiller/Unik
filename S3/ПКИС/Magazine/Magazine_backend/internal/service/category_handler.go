package service

import (
	"Magazine_backend/internal/database"
	"Magazine_backend/internal/models"
	"Magazine_backend/pkg/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary		Get categories
// @Security		ApiKeyAuth
// @Description	Retrieve a list of categories for a specific account by account ID with pagination
// @Tags			categories
// @Accept			json
// @Produce		json
// @Param			page	query		int		false	"Page number (default: 1)"
// @Param			limit	query		int		false	"Number of categories per page (default: 5)"
// @Success		200		{array}		models.Category
// @Failure		400		{object}	errorResponse
// @Failure		404		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/v1/categories [get]
func GetCategories(c *gin.Context) {
	// // Получение параметров из запроса
	// page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))   // Номер страницы, по умолчанию 1
	// limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5")) // Количество элементов на странице, по умолчанию 5

	// if page < 1 {
	// 	page = 1
	// }
	// if limit < 1 {
	// 	limit = 5
	// }

	// offset := (page - 1) * limit // Вычисление смещения

	var categories []models.Category

	// Использование GORM для выборки с лимитом и смещением
	err := database.DbPostgres.Find(&categories).Error //.Limit(limit).Offset(offset).Preload("Products")
	if err != nil {
		utils.Logger.Error("Неудачный запрос|(categories_handler.go|GetCategories|):", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, categories)

	utils.Logger.Printf("GetCategoriess сработал: %v", categories)
}

// @Summary		Get category by ID
// @Security		ApiKeyAuth
// @Description	Retrieve a specific category by its ID
// @Tags			manufactures
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Category ID"
// @Success		200	{object}	models.Category
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/manufactures/{id} [get]
func GetCategoryById(c *gin.Context) {
	// Получаем параметр id из запроса
	id := c.Param("id")

	// Использование GORM для поиска профиля по ID
	var category models.Category
	err := database.DbPostgres.Preload("Products").First(&category, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(category_handler.go|GetCategoryById|):", err)
		return
	}

	// Возвращаем профиль в ответе
	c.JSON(http.StatusOK, category)
}

// @Summary		Create a new category
// @Description	Creates a new category by accepting category details in the request body
// @Tags			sign
// @Accept			json
// @Produce		json
// @Param			category	body		models.Category	true	"Category data"
// @Success		201	{object}	models.Category
// @Failure		400	{object}	errorResponse
// @Failure		409	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/manufactures [post]
func CreateCategory(c *gin.Context) {
	req := models.CreateCategory{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(category_handler.go|CreateManufacturer|)|:", err)
		return
	}

	category := models.Category{
		Name: req.Name,
	}

	if err := database.DbPostgres.Create(&category).Error; err != nil {
		utils.Logger.Error("Insert isn't done(category_handler.go|CreateManufacturer|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	if err := database.DbPostgres.Preload("Products").Last(&category).Error; err != nil {
		utils.Logger.Error("Сan't receive category(category_handler.go|CreateManufacturer|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		database.DbPostgres.Delete(&category)
		return
	}
	c.JSON(http.StatusCreated, category)
}

// @Summary		Update an existing category
// @Security		ApiKeyAuth
// @Description	Update an existing category's information by category ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id		path		int			true	"Category ID"
// @Param			category	body		models.Category	true	"Updated category data"
// @Success		202	{object}	models.Category
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/users/{id} [put]
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	req := models.UpdateCategory{}

	// Парсим JSON из тела запроса
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(category_handler.go|UpdateCategory|)|:", err)
		return
	}

	// Проверка, есть ли такой id
	var existing models.Category
	if err := database.DbPostgres.Where("id = ?", id).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		utils.Logger.Error("DB error when checking category (category_handler.go|UpdateCategory):", err)
		return
	}

	// Обновляем категорию по ID с использованием GORM
	if err := database.DbPostgres.Model(&models.Category{}).Where("id = ?", id).Updates(req).Error; err != nil {
		utils.Logger.Error("Update isn't done(category_handler.go|UpdateCategory|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Использование GORM для поиска профиля по ID
	var category models.Category

	if err := database.DbPostgres.Preload("Products").First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(category_handler.go|UpdateCategory|):", err)
		return
	}

	// Отправляем успешный ответ с обновленным профилем
	c.JSON(http.StatusAccepted, category)
}

// @Summary		Delete a category by ID
// @Security		ApiKeyAuth
// @Description	Delete a category from the system by its ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"Category ID"
// @Success		202	{object}	string
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	// Удаляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Delete(&models.Category{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "category wasn't deleted"})
		utils.Logger.Error("Delete isn't done(category_handler.go|DeleteCategory|):", err)
		return
	}

	// Отправляем успешный ответ о удалении
	c.JSON(http.StatusAccepted, gin.H{"message": "category was deleted"})
}
