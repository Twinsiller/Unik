package service

import (
	"Magazine_backend/internal/database"
	"Magazine_backend/internal/models"
	"Magazine_backend/pkg/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"gorm.io/gorm"
)

// @Summary		Get products
// @Security		ApiKeyAuth
// @Description	Retrieve a list of products for a specific account by account ID with pagination
// @Tags			products
// @Accept			json
// @Produce		json
// @Param			page	query		int		false	"Page number (default: 1)"
// @Param			limit	query		int		false	"Number of products per page (default: 5)"
// @Success		200		{array}		models.Product
// @Failure		400		{object}	errorResponse
// @Failure		404		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/v1/products [get]
func GetProducts(c *gin.Context) {
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

	var products []models.Product

	// Использование GORM для выборки с лимитом и смещением
	err := database.DbPostgres.Limit(limit).Offset(offset).Preload("Manufacturer").Preload("Category").Preload("OrderItems").Preload("PurchaseItems").Preload("Stocks").Preload("StockMovements").Find(&products).Error
	if err != nil {
		utils.Logger.Error("Неудачный запрос|(product_handler.go|GetProducts|):", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	var count int64
	database.DbPostgres.Model(&models.Product{}).Count(&count)

	results := struct {
		Items []models.Product `json:"items"`
		Count int64            `json:"count"`
	}{
		Items: products,
		Count: count,
	}

	c.JSON(http.StatusOK, results)

	utils.Logger.Printf("GetProducts сработал: %v", results)
}

// @Summary		Get product by ID
// @Security		ApiKeyAuth
// @Description	Retrieve a specific product by its ID
// @Tags			products
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Product ID"
// @Success		200	{object}	models.Product
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/products/{id} [get]
func GetProductById(c *gin.Context) {
	// Получаем параметр id из запроса
	id := c.Param("id")

	// Использование GORM для поиска профиля по ID
	var product models.Product
	err := database.DbPostgres.Preload("Manufacturer").Preload("Category").Preload("OrderItems").Preload("PurchaseItems").Preload("Stocks").Preload("StockMovements").First(&product, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(product_handler.go|GetProductById|):", err)
		return
	}

	// Возвращаем профиль в ответе
	c.JSON(http.StatusOK, product)
}

// @Summary		Create a new product
// @Description	Creates a new product by accepting product details in the request body
// @Tags			sign
// @Accept			json
// @Produce		json
// @Param			product	body		models.Product	true	"Product data"
// @Success		201	{object}	models.Product
// @Failure		400	{object}	errorResponse
// @Failure		409	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/products [post]
func CreateProduct(c *gin.Context) {
	req := models.CreateProduct{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(product_handler.go|CreateProduct|)|:", err)
		return
	}

	product := models.Product{
		Name:           req.Name,
		CategoryID:     req.CategoryID,
		ManufacturerID: req.ManufacturerID,
		Unit:           req.Unit,
		Barcode:        req.Barcode,
		DefaultPrice:   req.DefaultPrice,
		MinStock:       req.MinStock,
	}

	if err := database.DbPostgres.Create(&product).Error; err != nil {
		utils.Logger.Error("Insert isn't done(product_handler.go|CreateProduct|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	if err := database.DbPostgres.Preload("Manufacturer").Preload("Category").Preload("OrderItems").Preload("PurchaseItems").Preload("Stocks").Preload("StockMovements").Last(&product).Error; err != nil {
		utils.Logger.Error("Сan't receive product(product_handler.go|CreateProduct|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		database.DbPostgres.Delete(&product)
		return
	}
	c.JSON(http.StatusCreated, product)
}

func UploadPhotoProduct(c *gin.Context) {
	if database.DbMongo == nil {
		c.JSON(500, gin.H{"error": "mongo not initialized"})
		utils.Logger.Error("Mongo not initialized(product_handler.go|UploadPhotoProduct|):")
		return
	}

	// 1. id из URL
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.Logger.Error("Invalid product id(product_handler.go|UploadPhotoProduct|):", err)
		c.JSON(400, gin.H{"error": "invalid product id"})
		return
	}

	// 2. файл (multipart)
	file, err := c.FormFile("file")
	if err != nil {
		utils.Logger.Error("File is required(product_handler.go|UploadPhotoProduct|):", err)
		c.JSON(400, gin.H{"error": "file is required"})
		return
	}

	f, err := file.Open()
	if err != nil {
		utils.Logger.Error("Cannot open file(product_handler.go|UploadPhotoProduct|):", err)
		c.JSON(500, gin.H{"error": "cannot open file"})
		return
	}
	defer f.Close()

	// 3. загрузка в Mongo GridFS
	photoID, err := database.UploadImage(
		database.DbMongo.Database(database.DBMongoName),
		file.Filename,
		f,
	)
	if err != nil {
		utils.Logger.Error("Upload failed(product_handler.go|UploadPhotoProduct|):", err)
		c.JSON(500, gin.H{"error": "upload failed"})
		return
	}

	// 4. сохранить photoID в Postgres
	if err := database.DbPostgres.Model(&models.Product{}).
		Where("id = ?", id).
		Update("photo_id", photoID).Error; err != nil {
		utils.Logger.Error("Db update failed(product_handler.go|UploadPhotoProduct|):", err)
		c.JSON(500, gin.H{"error": "db update failed"})
		return
	}

	c.JSON(200, gin.H{"photo_id": photoID})
}

func GetPhotoProduct(c *gin.Context) {
	// 1. id товара
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid product id"})
		return
	}

	var product models.Product
	// 2. получить photo_id из Postgres
	database.DbPostgres.First(&product, id)
	if product.PhotoID == nil {
		c.JSON(404, gin.H{"error": "photo not found"})
		return
	}

	// 3. ObjectID
	objID, err := primitive.ObjectIDFromHex(*product.PhotoID)
	if err != nil {
		c.JSON(500, gin.H{"error": "invalid photo id"})
		return
	}

	// 4. GridFS
	bucket, err := gridfs.NewBucket(
		database.DbMongo.Database(database.DBMongoName),
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "gridfs error"})
		return
	}

	// 5. headers
	c.Header("Content-Type", "image/jpeg") // или динамически
	c.Status(200)

	// 6. stream
	if _, err := bucket.DownloadToStream(objID, c.Writer); err != nil {
		c.JSON(500, gin.H{"error": "download failed"})
		return
	}
}

// @Summary		Update an existing product
// @Security		ApiKeyAuth
// @Description	Update an existing product's information by product ID
// @Tags			products
// @Accept			json
// @Produce		json
// @Param			id		path		int			true	"Product ID"
// @Param			product	body		models.Product	true	"Updated product data"
// @Success		202	{object}	models.Product
// @Failure		400	{object}	errorResponse
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	req := models.UpdateProduct{}

	// Парсим JSON из тела запроса
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		utils.Logger.Error("Data is bad|(product_handler.go|UpdateProduct|)|:", err)
		return
	}

	// Проверка, есть ли такой id
	var existing models.Product
	if err := database.DbPostgres.Where("id = ?", id).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		utils.Logger.Error("DB error when checking product (product_handler.go|UpdateProduct):", err)
		return
	}

	// Обновляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Model(&models.Product{}).Where("id = ?", id).Updates(req).Error; err != nil {
		utils.Logger.Error("Update isn't done(product_handler.go|UpdateProduct|):", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Использование GORM для поиска профиля по ID
	var product models.Product

	if err := database.DbPostgres.Preload("Manufacturer").Preload("Category").Preload("OrderItems").Preload("PurchaseItems").Preload("Stocks").Preload("StockMovements").First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		utils.Logger.Error("Неудачный запрос|(product_handler.go|UpdateProduct|):", err)
		return
	}

	// Отправляем успешный ответ с обновленным профилем
	c.JSON(http.StatusAccepted, product)
}

// @Summary		Delete a product by ID
// @Security		ApiKeyAuth
// @Description	Delete a product from the system by its ID
// @Tags			product
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"Product ID"
// @Success		202	{object}	string
// @Failure		404	{object}	errorResponse
// @Failure		500	{object}	errorResponse
// @Router			/v1/products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	// Удаляем профиль по ID с использованием GORM
	if err := database.DbPostgres.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product wasn't deleted"})
		utils.Logger.Error("Delete isn't done(product_handler.go|DeleteProduct|):", err)
		return
	}

	// Отправляем успешный ответ о удалении
	c.JSON(http.StatusAccepted, gin.H{"message": "product was deleted"})
}
