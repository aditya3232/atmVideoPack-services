package routes

import (
	"github.com/aditya3232/atmVideoPack-services.git/config"
	"github.com/aditya3232/atmVideoPack-services.git/connection"
	"github.com/aditya3232/atmVideoPack-services.git/handler"
	"github.com/aditya3232/atmVideoPack-services.git/middleware"
	"github.com/aditya3232/atmVideoPack-services.git/model/download_playback"
	"github.com/aditya3232/atmVideoPack-services.git/model/get_download_playback_from_elastic"
	"github.com/aditya3232/atmVideoPack-services.git/model/get_human_detection_from_elastic"
	"github.com/aditya3232/atmVideoPack-services.git/model/get_status_mc_detection_from_elastic"
	"github.com/aditya3232/atmVideoPack-services.git/model/get_vandal_detection_from_elastic"
	"github.com/aditya3232/atmVideoPack-services.git/model/streaming_cctv"
	"github.com/aditya3232/atmVideoPack-services.git/model/tb_tid"
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	// Initialize repositories
	tbTidRepository := tb_tid.NewRepository(connection.DatabaseMysql())
	elasticHumanDetectionIndexRepository := get_human_detection_from_elastic.NewRepository(connection.ElasticSearch())
	elasticVandalDetectionIndexRepository := get_vandal_detection_from_elastic.NewRepository(connection.ElasticSearch())
	elasticStatusMcDetectionRepository := get_status_mc_detection_from_elastic.NewRepository(connection.ElasticSearch())
	elasticDownloadPlaybackRepository := get_download_playback_from_elastic.NewRepository(connection.ElasticSearch())
	downloadPlaybackRepository := download_playback.NewRepository(connection.DatabaseMysql())
	streamingCctvRepository := streaming_cctv.NewRepository(connection.DatabaseMysql())

	// Initialize services
	tbTidService := tb_tid.NewService(tbTidRepository)
	elasticHumanDetectionIndexService := get_human_detection_from_elastic.NewService(elasticHumanDetectionIndexRepository)
	elasticVandalDetectionIndexService := get_vandal_detection_from_elastic.NewService(elasticVandalDetectionIndexRepository)
	elasticStatusMcDetectionService := get_status_mc_detection_from_elastic.NewService(elasticStatusMcDetectionRepository)
	elasticDownloadPlaybackService := get_download_playback_from_elastic.NewService(elasticDownloadPlaybackRepository)
	downloadPlaybackService := download_playback.NewService(downloadPlaybackRepository, tbTidRepository)
	streamingCctvService := streaming_cctv.NewService(streamingCctvRepository, tbTidRepository)

	// Initialize handlers
	tbTidHandler := handler.NewTbTidHandler(tbTidService)
	elasticHumanDetectionIndexHandler := handler.NewGetHumanDetectionFromElasticHandler(elasticHumanDetectionIndexService)
	elasticVandalDetectionIndexHandler := handler.NewGetVandalDetectionFromElasticHandler(elasticVandalDetectionIndexService)
	elasticStatusMcDetectionHandler := handler.NewGetStatusMcDetectionFromElasticHandler(elasticStatusMcDetectionService)
	elasticDownloadPlaybackHandler := handler.NewGetDownloadPlaybackFromElasticHandler(elasticDownloadPlaybackService)
	downloadPlaybackHandler := handler.NewDownloadPlaybackHandler(downloadPlaybackService)
	streamingCctvHandler := handler.NewStreamingCctvHandler(streamingCctvService)

	// Configure routes
	api := router.Group("/api/atmvideopack/v1")

	tbTidRoutes := api.Group("/device", middleware.ApiKeyMiddleware(config.CONFIG.API_KEY))
	streamingCctvRoutes := api.Group("/stream")
	elasticHumanDetectionIndexRoutes := api.Group("/humandetection", middleware.ApiKeyMiddleware(config.CONFIG.API_KEY))
	elasticVandalDetectionIndexRoutes := api.Group("/vandaldetection", middleware.ApiKeyMiddleware(config.CONFIG.API_KEY))
	elasticStatusMcDetectionRoutes := api.Group("/statusmcdetection", middleware.ApiKeyMiddleware(config.CONFIG.API_KEY))
	elasticDownloadPlaybackRoutes := api.Group("/downloadplayback", middleware.ApiKeyMiddleware(config.CONFIG.API_KEY))
	downloadPlaybackRoutes := api.Group("/downloadvideoplayback")

	configureTbTidRoutes(tbTidRoutes, tbTidHandler)
	configureStreamingCctvRoutes(streamingCctvRoutes, streamingCctvHandler)
	configureElasticHumanDetectionIndexRoutes(elasticHumanDetectionIndexRoutes, elasticHumanDetectionIndexHandler)
	configureElasticVandalDetectionIndexRoutes(elasticVandalDetectionIndexRoutes, elasticVandalDetectionIndexHandler)
	configureElasticStatusMcDetectionIndexRoutes(elasticStatusMcDetectionRoutes, elasticStatusMcDetectionHandler)
	configureElasticDownloadPlaybackIndexRoutes(elasticDownloadPlaybackRoutes, elasticDownloadPlaybackHandler)
	configureDownloadPlaybackRoutes(downloadPlaybackRoutes, downloadPlaybackHandler)

}

func configureTbTidRoutes(group *gin.RouterGroup, handler *handler.TbTidHandler) {
	group.POST("/create", handler.CreateTbTid)
	group.POST("/getonebyid", handler.GetOneByID)
	group.GET("/getall", handler.GetAllTbEntries)
}

func configureStreamingCctvRoutes(group *gin.RouterGroup, handler *handler.StreamingCctvHandler) {
	group.GET("/cctv/:tid", handler.StreamingCctv)
}

func configureElasticHumanDetectionIndexRoutes(group *gin.RouterGroup, handler *handler.GetHumanDetectionFromElasticHandler) {
	group.POST("/getall", handler.FindAll)
}

func configureElasticVandalDetectionIndexRoutes(group *gin.RouterGroup, handler *handler.GetVandalDetectionFromElasticHandler) {
	group.POST("/getall", handler.FindAll)
}

func configureElasticStatusMcDetectionIndexRoutes(group *gin.RouterGroup, handler *handler.GetStatusMcDetectionFromElasticHandler) {
	group.POST("/getall", handler.FindAll)
	group.POST("/getdeviceupanddown", handler.FindDeviceUpAndDown)
}

func configureElasticDownloadPlaybackIndexRoutes(group *gin.RouterGroup, handler *handler.GetDownloadPlaybackFromElasticHandler) {
	group.POST("/getall", handler.FindAll)
}

func configureDownloadPlaybackRoutes(group *gin.RouterGroup, handler *handler.DownloadPlaybackHandler) {
	group.GET("/:videos/:tid/:yyyymmdd/:filename", handler.DownloadPlayback)
}
