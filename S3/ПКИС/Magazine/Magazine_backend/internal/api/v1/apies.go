package v1

import (
	"Magazine_backend/internal/service"
	"Magazine_backend/pkg/utils"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Apies() {
	router := gin.Default()
	// Маршрут для Swagger UI
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Создание нового профиля
	router.POST("/register", service.CreateUser)

	routerv1 := router.Group("/v1")

	users := routerv1.Group("/users")
	{
		// Получение всех профилей
		users.GET("", service.GetUsers)

		// Получение поста по ID
		users.GET("/:id", service.GetUserById)

		// Обновление существующего профиля
		users.PUT("/:id", service.UpdateUser)

		// Удаление профиля
		users.DELETE("/:id", service.DeleteUserById)
	}

	orders := routerv1.Group("/orders")
	{
		// Получение всех заказов
		orders.GET("", service.GetOrders)

		// Получение заказа по ID
		orders.GET("/:id", service.GetOrderById)

		// Создание нового заказа
		orders.POST("", service.CreateOrder)

		// // Обновление существующего заказа
		orders.PUT("/:id", service.UpdateOrder)

		// // Удаление заказа
		orders.DELETE("/:id", service.DeleteOrder)
	}

	manufacturers := routerv1.Group("/manufacturers")
	{
		// Получение всех заказов
		manufacturers.GET("", service.GetManufacturers)

		// Получение заказа по ID
		manufacturers.GET("/:id", service.GetManufacturerById)

		// Создание нового заказа
		manufacturers.POST("", service.CreateManufacturer)

		// // Обновление существующего заказа
		manufacturers.PUT("/:id", service.UpdateManufacturer)

		// // Удаление заказа
		manufacturers.DELETE("/:id", service.DeleteManufacturer)
	}

	categories := routerv1.Group("/categories")
	{
		// Получение всех заказов
		categories.GET("", service.GetCategories)

		// Получение заказа по ID
		categories.GET("/:id", service.GetCategoryById)

		// Создание нового заказа
		categories.POST("", service.CreateCategory)

		// // Обновление существующего заказа
		categories.PUT("/:id", service.UpdateCategory)

		// // Удаление заказа
		categories.DELETE("/:id", service.DeleteCategory)
	}

	suppliers := routerv1.Group("/suppliers")
	{
		// Получение всех заказов
		suppliers.GET("", service.GetSuppliers)

		// Получение заказа по ID
		suppliers.GET("/:id", service.GetSupplierById)

		// Создание нового заказа
		suppliers.POST("", service.CreateSupplier)

		// // Обновление существующего заказа
		suppliers.PUT("/:id", service.UpdateSupplier)

		// // Удаление заказа
		suppliers.DELETE("/:id", service.DeleteSupplier)
	}

	purchases := routerv1.Group("/purchases")
	{
		// Получение всех заказов
		purchases.GET("", service.GetPurchases)

		// Получение заказа по ID
		purchases.GET("/:id", service.GetPurchaseById)

		// Создание нового заказа
		purchases.POST("", service.CreatePurchase)

		// // Обновление существующего заказа
		purchases.PUT("/:id", service.UpdatePurchase)

		// // Удаление заказа
		purchases.DELETE("/:id", service.DeletePurchase)
	}

	products := routerv1.Group("/products")
	{
		// Получение всех заказов
		products.GET("", service.GetProducts)

		// Получение заказа по ID
		products.GET("/:id", service.GetProductById)

		// Создание нового заказа
		products.POST("", service.CreateProduct)

		// // Обновление существующего заказа
		products.PUT("/:id", service.UpdateProduct)

		// // Удаление заказа
		products.DELETE("/:id", service.DeleteProduct)
	}

	purchaseItems := routerv1.Group("/purchaseItems")
	{
		// Получение всех заказов
		purchaseItems.GET("", service.GetPurchaseItems)

		// Получение заказа по ID
		purchaseItems.GET("/:id", service.GetPurchaseItemById)

		// Создание нового заказа
		purchaseItems.POST("", service.CreatePurchaseItem)

		// // Обновление существующего заказа
		purchaseItems.PUT("/:id", service.UpdatePurchaseItem)

		// // Удаление заказа
		purchaseItems.DELETE("/:id", service.DeletePurchaseItem)
	}

	orderItems := routerv1.Group("/orderItems")
	{
		// Получение всех заказов
		orderItems.GET("", service.GetOrderItems)

		// Получение заказа по ID
		orderItems.GET("/:id", service.GetOrderItemById)

		// Создание нового заказа
		orderItems.POST("", service.CreateOrderItem)
		// // Обновление существующего заказа
		orderItems.PUT("/:id", service.UpdateOrderItem)

		// // Удаление заказа
		orderItems.DELETE("/:id", service.DeleteOrderItem)
	}

	stocks := routerv1.Group("/stocks")
	{
		// Получение всех заказов
		stocks.GET("", service.GetStocks)

		// Получение заказа по ID
		stocks.GET("/:id", service.GetStockById)

		// Создание нового заказа
		stocks.POST("", service.CreateStock)
		// // Обновление существующего заказа
		stocks.PUT("/:id", service.UpdateStock)

		// // Удаление заказа
		stocks.DELETE("/:id", service.DeleteStock)
	}

	stockMovements := routerv1.Group("/stocks")
	{
		// Получение всех заказов
		stockMovements.GET("", service.GetStockMovements)

		// Получение заказа по ID
		stockMovements.GET("/:id", service.GetStockMovementById)

		// Создание нового заказа
		stockMovements.POST("", service.CreateStockMovement)

		// // Обновление существующего заказа
		stockMovements.PUT("/:id", service.UpdateStockMovement)

		// // Удаление заказа
		stockMovements.DELETE("/:id", service.DeleteStockMovement)
	}

	//router.Run(":8080") // Это прошлое

	// Создаём кастомный сервер
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Запуск сервера в горутине
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Logger.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	// Канал для сигналов завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Блокировка до получения сигнала
	sig := <-quit // программа здесь "ждёт" сигнал

	utils.Logger.Warn("Завершение работы сервера...")

	if sig == os.Interrupt {
		utils.Logger.Info("Пойман сигнал (Ctrl + C):", sig)
	} else {
		utils.Logger.Info("Вызов из вне (Shutdown()):", sig)
	}

	/* if err := service.DumpDataToFile(); err != nil {
		utils.Logger.Error("Ошибка при выгрузке данных:", err)
	} */

	// Таймаут для graceful shutdown
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	// Остановка сервера
	if err := srv.Shutdown(ctx); err != nil {
		utils.Logger.Fatal("Server forced to shutdown:", err)
	}

	utils.Logger.Println("Server exiting")
}
