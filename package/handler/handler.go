package handler

import (
	// "fmt"
	"fmt"
	// "norbekov/docs"
	"norbekov/docs"
	"norbekov/package/service"
	"norbekov/util/logrus"

	"norbekov/configs"
	_ "norbekov/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
	logrus   *logrus.Logger
	config   *configs.Configs
}

func NewHandler(services *service.Service, logrus *logrus.Logger, config *configs.Configs) *Handler {
	return &Handler{services: services, logrus: logrus, config: config}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	config := handler.config
	fmt.Println(config)

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
	{
		home.POST("/create", handler.CreateHomePost)
		home.PATCH("/upload-img/:id", handler.uploadHomeImage)
		home.PUT("/update/:id", handler.updateHome)
		router.Group("/home").GET("/get", handler.GetHomeById)
		home.DELETE("/delete", handler.DeleteHome)
	}
	news := api.Group("/news")
	{
		news.POST("/create", handler.CreateNewsPost)
		news.PATCH("/upload-img/:id", handler.uploadNewsImage)
		news.PUT("/update/:id", handler.updateNews)
		news.GET("/get", handler.GetNewsById)
		news.DELETE("/delete", handler.DeleteNews)
	}
	service := api.Group("/service")
	{
		service.POST("/create", handler.CreateServicePost)
		service.PATCH("/upload-img/:id", handler.uploadServiceImage)
		service.PUT("/update/:id", handler.UpdateService)
		service.GET("/get", handler.GetServiceById)
		service.DELETE("/delete", handler.DeleteService)
	}
	table := api.Group("/table")
	{
		table.POST("/create", handler.CreateTablePost)
		table.PATCH("/upload-img/:id", handler.uploadTableImage)
		table.PUT("/update/:id", handler.UpdateTable)
		table.GET("/get", handler.GetTableById)
		table.DELETE("/delete", handler.DeleteTable)
	}

	contact := router.Group("/contact")
	contacts := api.Group("/contacts")
	{
		contact.POST("/create", handler.CreateContactPost)
		contacts.GET("/get", handler.GetAllContact)
	}

	return router
}
