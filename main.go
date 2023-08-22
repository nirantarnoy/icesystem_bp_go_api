package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tarlek.com/icesystem/config"
	"tarlek.com/icesystem/controller"
	"tarlek.com/icesystem/middleware"
	"tarlek.com/icesystem/repository"
	"tarlek.com/icesystem/service"
)

var (
	db         *gorm.DB           = config.SetupDB()
	jwtService service.JWTService = service.NewJWTService()

	userRepository repository.UserRepository = repository.NewUserRepository(db)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService, jwtService)

	orderRepository repository.OrderRepository = repository.NewOrderRepository(db)
	orderService    service.OrderService       = service.NewOrderService(orderRepository)
	orderController controller.OrderController = controller.NewOrderController(orderService, jwtService)

	posCloseRepository repository.PosCloseRepository = repository.NewPosCloseRepository(db)
	posCloseService    service.PosCloseService       = service.NewPosCloseService(posCloseRepository)
	posCloseController controller.PoscloseController = controller.NewPosCloseController(posCloseService)
)

func main() {
	defer config.CloseDBConn(db)

	server := gin.Default()

	//cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}
	// corsConfig.AllowHeaders = []string{"Content-Type", "application/json"}
	corsConfig.AllowCredentials = true
	server.Use(cors.New(corsConfig))

	authRoute := server.Group("api/auth", authController.Login)
	{
		authRoute.POST("/login")
	}

	userRoute := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoute.GET("/profile/:id", userController.Profile)
	}

	// orderRoute := server.Group("api/order", middleware.AuthorizeJWT(jwtService))
	// {
	// 	orderRoute.POST("/addorder", orderController.CreateOrder)
	// 	orderRoute.GET("/getlastno", orderController.GetLastNo)
	// 	// orderRoute.GET("/listorderbyemp/:id", orderController.ListorderByEmpId)
	// 	// orderRoute.POST("/closeorder", orderController.Closeorder)
	// 	// orderRoute.POST("/listorderopen/:id", orderController.ListorderOpen)
	// 	// orderRoute.GET("/getorderphoto/:id", orderController.GetorderPhoto)
	// 	// orderRoute.GET("/getorderclosephoto/:id", orderController.GetorderClosePhoto)

	// }
	orderRoute := server.Group("api/order")
	{
		orderRoute.POST("/addorder", orderController.CreateOrder)
		orderRoute.GET("/getlastno", orderController.GetLastNo)
		orderRoute.POST("/closeorder", orderController.CloseOrder)
		orderRoute.POST("/orderlist", orderController.CustomerOrder)
		// orderRoute.GET("/listorderbyemp/:id", orderController.ListorderByEmpId)
		// orderRoute.POST("/closeorder", orderController.Closeorder)
		// orderRoute.POST("/listorderopen/:id", orderController.ListorderOpen)
		// orderRoute.GET("/getorderphoto/:id", orderController.GetorderPhoto)
		// orderRoute.GET("/getorderclosephoto/:id", orderController.GetorderClosePhoto)

	}

	poscloseRoute := server.Group("api/pos")
	{
		poscloseRoute.POST("/posclose", posCloseController.PosClose)
	}

	server.Run(":1224")
}
