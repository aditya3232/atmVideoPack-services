package routes

import (
	"github.com/aditya3232/atmVideoPack-services.git/connection"
	"github.com/aditya3232/atmVideoPack-services.git/handler"
	"github.com/aditya3232/atmVideoPack-services.git/model/get_human_detection_from_elastic"
	"github.com/aditya3232/atmVideoPack-services.git/model/tb_tid"
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	// Initialize repositories
	tbTidRepository := tb_tid.NewRepository(connection.DatabaseMysql())
	elasticHumanDetectionIndexRepository := get_human_detection_from_elastic.NewRepository(connection.ElasticSearch())

	// Initialize services
	tbTidService := tb_tid.NewService(tbTidRepository)
	elasticHumanDetectionIndexService := get_human_detection_from_elastic.NewService(elasticHumanDetectionIndexRepository)

	// Initialize handlers
	tbTidHandler := handler.NewTbTidHandler(tbTidService)
	elasticHumanDetectionIndexHandler := handler.NewGetHumanDetectionFromElasticHandler(elasticHumanDetectionIndexService)

	// Configure routes
	api := router.Group("/api/atmvideopack/v1")

	// tbTidRoutes := api.Group("/device", middleware.ApiKeyMiddleware(config.CONFIG.API_KEY))
	tbTidRoutes := api.Group("/device")
	elasticHumanDetectionIndexRoutes := api.Group("/humandetection")

	configureTbTidRoutes(tbTidRoutes, tbTidHandler)
	configureElasticHumanDetectionIndexRoutes(elasticHumanDetectionIndexRoutes, elasticHumanDetectionIndexHandler)

}

func configureTbTidRoutes(group *gin.RouterGroup, handler *handler.TbTidHandler) {
	group.POST("/create", handler.CreateTbTid)
	group.POST("/getonebyid", handler.GetOneByID)
	group.GET("/getall", handler.GetAllTbEntries)

	// GetStreamVideo
	// group.GET("/getstreamvideo", handler.GetStreamVideo)
	group.GET("/getstreamvideo/:id", handler.GetStreamVideo)

}

func configureElasticHumanDetectionIndexRoutes(group *gin.RouterGroup, handler *handler.GetHumanDetectionFromElasticHandler) {
	group.POST("/getall", handler.FindAll)
}
