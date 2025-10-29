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
