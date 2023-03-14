package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) initRouter() {
	router := gin.Default()
	router.Use(s.setTranslations())
	router.Use(s.setCors())
	// 注册Swagger文档路由
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 服务检查
	router.GET("/v1/api/ping", s.pingHandle)

	// 发送短信
	router.POST("/v1/api/sms", s.sendSMS)

	// 静态资源访问
	router.Static("/storage/uploads", "./storage/uploads")

	// auth 登录/认证模块
	authRouter := router.Group("/v1/api/auth")
	{
		authRouter.POST("/login", s.loginUser)
		authRouter.POST("/refresh", s.refreshToken)
	}

	// admin 管理员模块
	adminRouter := router.Group("/v1/api/admin").Use(s.authentication())
	{
		// 上传文件
		adminRouter.POST("/uploadFiles", s.uploadFiles).Use(s.authentication())

		adminRouter.POST("/profile", s.getProfile)
		adminRouter.POST("/users", s.createUser)
		adminRouter.GET("/users", s.getListUsers) // /v1/users?page_num=1&page_size=10
		adminRouter.POST("/users/:id", s.updateUser)
		adminRouter.GET("/users/:id", s.getUserById)

		adminRouter.GET("/roles", s.getListRoles)
		adminRouter.GET("/roles/:id", s.getRoleById)

		adminRouter.POST("/theaters", s.createTheater)
		adminRouter.GET("/theaters", s.listTheaters)
		adminRouter.GET("/theaters/:id", s.getTheater)
		adminRouter.PUT("/theaters/:id", s.updateTheater)
		adminRouter.DELETE("/theaters/:id", s.delTheater)

		adminRouter.POST("/halls", s.createHall)
		adminRouter.GET("/halls", s.listHalls)
		adminRouter.GET("/halls/:id", s.getHall)
		adminRouter.PUT("/halls/:id", s.updateHall)
		adminRouter.DELETE("/halls/:id", s.delHall)

		adminRouter.POST("/seats", s.createSeat)
		adminRouter.GET("/seats", s.listSeats)
		adminRouter.GET("/seats/:id", s.getSeat)
		adminRouter.PUT("/seats/:id", s.updateSeat)
		adminRouter.DELETE("/seats/:id", s.delSeat)

		adminRouter.POST("/movies", s.createMovie)
		adminRouter.GET("/movies", s.listMovies)
		adminRouter.GET("/movies/:id", s.getMovie)
		adminRouter.PUT("/movies/:id", s.updateMovie)
		adminRouter.DELETE("/movies/:id", s.delMovie)

		adminRouter.POST("/showtimes", s.createShowtime)
		adminRouter.GET("/showtimes", s.listShowtimes)
		adminRouter.GET("/showtimes/:id", s.getShowtime)
		adminRouter.PUT("/showtimes/:id", s.updateShowtime)
		adminRouter.DELETE("/showtimes/:id", s.delShowtime)

		adminRouter.POST("/tickets", s.createTicket)
		adminRouter.GET("/tickets", s.listTickets)
		adminRouter.GET("/tickets/:id", s.getTicket)
		adminRouter.PUT("/tickets/:id", s.updateTicket)
		adminRouter.DELETE("/tickets/:id", s.delTicket)

		adminRouter.POST("/payments", s.createPayment)
		adminRouter.GET("/payments", s.listPayments)
		adminRouter.GET("/payments/:id", s.getPayment)
		adminRouter.PUT("/payments/:id", s.updatePayment)
		adminRouter.DELETE("/payments/:id", s.delPayment)

	}

	s.router = router
}
