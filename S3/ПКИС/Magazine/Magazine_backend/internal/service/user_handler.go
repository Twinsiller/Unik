package service

import (
	"Magazine_backend/internal/database"
	"Magazine_backend/internal/models"
	"Magazine_backend/pkg/utils"
	"errors"
	"net/http"
	"strconv"

	password "github.com/vzglad-smerti/password_hash"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary		Get users
// @Security		ApiKeyAuth
// @Description	Retrieve a list of users for a specific account by account ID with pagination
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			page	query		int		false	"Page number (default: 1)"
// @Param			limit	query		int		false	"Number of users per page (default: 5)"
// @Success		200		{array}		models.User
// @Failure		400		{object}	errorResponse
// @Failure		404		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/v1/users [get]
func GetUsers(c *gin.Context) {
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

	var users []models.User

	// Использование GORM для выборки с лимитом и смещением
	err := database.DbPostgres.Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		utils.Logger.Error("Неудачный запрос|(user_handler.go|GetUsers|):", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, users)

	utils.Logger.Printf("GetUsers сработал: %v", users)
}

// @Summary		Get user by ID
// @Security		ApiKeyAuth
// @Description	Retrieve a specific user by its ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Account ID"
// @Success		200	{object}	models.User
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/users/{id} [get]
func GetUserById(c *gin.Context) {
	// Получаем параметр id из запроса
	id := c.Param("id")

	// Использование GORM для поиска профиля по ID
	var user models.User
	err := database.DbPostgres.First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(user_handler.go|GetUserById|):", err)
		return
	}

	// Возвращаем профиль в ответе
	c.JSON(http.StatusOK, user)
}

// @Summary		Create a new user
// @Description	Creates a new user by accepting user details in the request body
// @Tags			sign
// @Accept			json
// @Produce		json
// @Param			user	body		models.User	true	"User data"
// @Success		201	{object}	models.User
// @Failure		400	{object}	errorResponse
// @Failure		409	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/register [post]
func CreateUser(c *gin.Context) {
	req := models.CreateUserRequest{}

	// Парсим JSON из тела запроса в структуру User
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(user_handler.go|CreateUser|)|:", err)
		return
	}

	// Проверка, есть ли уже такой nickname
	var existing models.User
	if err := database.DbPostgres.Where("nickname = ?", req.Nickname).First(&existing).Error; err == nil {
		// Нашли совпадение
		c.JSON(http.StatusConflict, gin.H{"message": "nickname already taken"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Ошибка при запросе
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		utils.Logger.Error("DB error when checking nickname (user_handler.go|CreateUser):", err)
		return
	}

	// Хеширование пароля
	hash, err := password.Hash(req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Problem with password hashing"})
		utils.Logger.Error("Hash wasn't working(user_handler.go|CreateUser|):", err)
		return
	}

	p := models.User{
		Nickname:     req.Nickname,
		HashPassword: hash,
		Role:         req.Role,
		Name:         req.Name,
	}

	// Создаем новый профиль в базе данных с использованием GORM
	if err := database.DbPostgres.Create(&p).Error; err != nil {
		utils.Logger.Error("Insert isn't done(user_handler.go|CreateUser|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Отправляем успешный ответ с созданным профилем
	c.JSON(http.StatusCreated, p)
}

// @Summary		Update an existing user
// @Security		ApiKeyAuth
// @Description	Update an existing user's information by user ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id		path		int			true	"User ID"
// @Param			user	body		models.User	true	"Updated user data"
// @Success		202	{object}	models.User
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	req := models.CreateUserRequest{}

	// Парсим JSON из тела запроса
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(user_handler.go|UpdateUser|)|:", err)
		return
	}

	// Проверка, есть ли такой id
	var existing models.User
	if err := database.DbPostgres.Where("id = ?", id).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		utils.Logger.Error("DB error when checking user (user_handler.go|UpdateUser):", err)
		return
	}

	var hash string
	var err error
	if req.Password != "" {
		// Хеширование пароля
		hash, err = password.Hash(req.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Problem with password hashing"})
			utils.Logger.Error("Hash wasn't working(user_handler.go|CreateUser|):", err)
			return
		}
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

// @Summary		Delete a user by ID
// @Security		ApiKeyAuth
// @Description	Delete a user from the system by its ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"User ID"
// @Success		202	{object}	string
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/users/{id} [delete]
func DeleteUserById(c *gin.Context) {
	id := c.Param("id")

	// Удаляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user wasn't deleted"})
		utils.Logger.Error("Delete isn't done(user_handler.go|DeleteUser|):", err)
		return
	}

	// Отправляем успешный ответ о удалении
	c.JSON(http.StatusAccepted, gin.H{"message": "user was deleted"})
}

// @Summary		Delete a user by Nickname
// @Security		ApiKeyAuth
// @Description	Delete a user from the system by its Nickname
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"User ID"
// @Success		202	{object}	string
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/users/{nickname} [delete]
func DeleteUserByNickname(c *gin.Context) {
	nickname := c.Param("nickname")

	// Удаляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Where("nickname = ?", nickname).Delete(&models.User{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user wasn't deleted"})
		utils.Logger.Error("Delete isn't done(user_handler.go|DeleteUser|):", err)
		return
	}

	// Отправляем успешный ответ о удалении
	c.JSON(http.StatusAccepted, gin.H{"message": "user was deleted"})
}
