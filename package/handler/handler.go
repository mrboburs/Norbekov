package handler

import (
	// "fmt"
	// "fmt"
	// "norbekov/docs"
	"github.com/mrboburs/Norbekov/docs"
	"github.com/mrboburs/Norbekov/package/service"
	"github.com/mrboburs/Norbekov/util/logrus"

	"github.com/mrboburs/Norbekov/configs"
	_ "github.com/mrboburs/Norbekov/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
	logrus   *logrus.Logger
}

func NewHandler(services *service.Service, logrus *logrus.Logger, config *configs.Configs) *Handler {
	return &Handler{services: services, logrus: logrus}
}

func (handler *Handler) InitRoutes() *gin.Engine {

	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	admin := router.Group("/admin")
	{
		admin.POST("/create", handler.CreateAdmin)
		admin.DELETE("/delete", handler.DeleteAdmin)
		admin.POST("/login", handler.LoginAdmin)
	}
	api := router.Group("/api", handler.userIdentity)
	home := api.Group("/home")
	homeGet := router.Group("/home")
	{
		home.POST("/", handler.CreateHomePost)
		home.PATCH("/:id", handler.uploadHomeImage)
		home.PUT("/:id", handler.updateHome)
		homeGet.GET("/", handler.GetHomeById)
		home.DELETE("/", handler.DeleteHome)
	}
	news := api.Group("/news")
	newsGet := router.Group("/news")
	{
		news.POST("/", handler.CreateNewsPost)
		news.PATCH("/:id", handler.uploadNewsImage)
		news.PUT("/:id", handler.updateNews)
		newsGet.GET("/", handler.GetNewsById)
		news.DELETE("/", handler.DeleteNews)
	}
	service := api.Group("/service")
	serviceGet := router.Group("/service")
	{
		service.POST("/", handler.CreateServicePost)
		service.PATCH("/:id", handler.uploadServiceImage)
		service.PUT("/:id", handler.UpdateService)
		serviceGet.GET("/", handler.GetServiceById)
		service.DELETE("/", handler.DeleteService)
	}
	table := api.Group("/table")
	tableGet := router.Group("/table")
	{
		table.POST("/", handler.CreateTablePost)
		table.PATCH("/:id", handler.uploadTableImage)
		table.PUT("/:id", handler.UpdateTable)
		tableGet.GET("/", handler.GetTableById)
		table.DELETE("/", handler.DeleteTable)
	}

	contact := router.Group("/contact")
	contacts := api.Group("/contacts")
	{
		contact.POST("/", handler.CreateContactPost)
		contacts.GET("/", handler.GetAllContact)
	}

	return router
}
