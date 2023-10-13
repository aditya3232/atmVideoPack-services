package routes

import (
	"github.com/aditya3232/gatewatchApp-services.git/config"
	"github.com/aditya3232/gatewatchApp-services.git/connection"
	"github.com/aditya3232/gatewatchApp-services.git/handler"
	"github.com/aditya3232/gatewatchApp-services.git/middleware"
	"github.com/aditya3232/gatewatchApp-services.git/model/tb_tid"
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	// Initialize repositories
	tbTidRepository := tb_tid.NewRepository(connection.DatabaseMysql())

	// Initialize services
	tbTidService := tb_tid.NewService(tbTidRepository)

	// Initialize handlers
	tbTidHandler := handler.NewTbTidHandler(tbTidService)

	// Configure routes
	api := router.Group("/api/atmvideopack/v1")

	tbTidRoutes := api.Group("/device", middleware.ApiKeyMiddleware(config.CONFIG.API_KEY))

	configureTbTidRoutes(tbTidRoutes, tbTidHandler)

}

func configureTbTidRoutes(group *gin.RouterGroup, handler *handler.TbTidHandler) {
	group.POST("/create", handler.CreateTbTid)
}
