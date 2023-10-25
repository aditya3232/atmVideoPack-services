package cron

import (
	"sync"

	"github.com/aditya3232/atmVideoPack-services.git/connection"
	"github.com/aditya3232/atmVideoPack-services.git/helper"
	log_function "github.com/aditya3232/atmVideoPack-services.git/log"
	"github.com/aditya3232/atmVideoPack-services.git/model/del_old_log_from_elastic"
	"github.com/aditya3232/atmVideoPack-services.git/model/del_old_log_human_detection_from_elastic"
	"github.com/aditya3232/atmVideoPack-services.git/model/del_old_log_status_mc_detection_from_elastic"
	"github.com/aditya3232/atmVideoPack-services.git/model/del_old_log_vandal_detection_from_elastic"

	"github.com/robfig/cron/v3"
)

func init() {
	log_function.Info("cronjob started")
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer helper.RecoverPanic() // Menambahkan recover di dalam goroutine
		defer wg.Done()
		// cron := cron.New()
		cron := cron.New(cron.WithChain(
			cron.SkipIfStillRunning(cron.DefaultLogger),
		))

		// hapus tiap tengah malam
		cron.AddFunc("0 0 * * *", func() {
			// cron.AddFunc("@every 5s", func() {
			delOldLogFromElasticRepository := del_old_log_from_elastic.NewRepository(connection.ElasticSearch())
			delOldLogFromElasticService := del_old_log_from_elastic.NewService(delOldLogFromElasticRepository)

			err := delOldLogFromElasticService.DelOneMonthOldLogs()
			if err != nil {
				log_function.Error("Error delete log:", err)
			}

			log_function.Info("delete log in elastic berhasil dilakukan")
		})

		cron.Start()
	}()

	wg.Add(1)
	go func() {
		defer helper.RecoverPanic() // Menambahkan recover di dalam goroutine
		defer wg.Done()
		cron := cron.New(cron.WithChain(
			cron.SkipIfStillRunning(cron.DefaultLogger),
		))

		cron.AddFunc("0 0 * * *", func() {
			// cron.AddFunc("@every 5s", func() {
			delOldHumanDetectionFromElasticRepository := del_old_log_human_detection_from_elastic.NewRepository(connection.ElasticSearch())
			delOldHumanDetectionFromElasticService := del_old_log_human_detection_from_elastic.NewService(delOldHumanDetectionFromElasticRepository)

			err := delOldHumanDetectionFromElasticService.DelOneMonthOldHumanDetectionLogs()
			if err != nil {
				log_function.Error("Error delete human detection:", err)
			}

			log_function.Info("delete human detection in elastic berhasil dilakukan")
		})

		cron.Start()
	}()

	wg.Add(1)
	go func() {
		defer helper.RecoverPanic() // Menambahkan recover di dalam goroutine
		defer wg.Done()
		cron := cron.New(cron.WithChain(
			cron.SkipIfStillRunning(cron.DefaultLogger),
		))

		cron.AddFunc("0 0 * * *", func() {
			// cron.AddFunc("@every 5s", func() {
			delOldVandalDetectionFromElasticRepository := del_old_log_vandal_detection_from_elastic.NewRepository(connection.ElasticSearch())
			delOldVandalDetectionFromElasticService := del_old_log_vandal_detection_from_elastic.NewService(delOldVandalDetectionFromElasticRepository)

			err := delOldVandalDetectionFromElasticService.DelOneMonthOldVandalDetectionLogs()
			if err != nil {
				log_function.Error("Error delete vandal detection:", err)
			}

			log_function.Info("delete vandal detection in elastic berhasil dilakukan")
		})

		cron.Start()
	}()

	wg.Add(1)
	go func() {
		defer helper.RecoverPanic() // Menambahkan recover di dalam goroutine
		defer wg.Done()
		cron := cron.New(cron.WithChain(
			cron.SkipIfStillRunning(cron.DefaultLogger),
		))

		cron.AddFunc("0 0 * * *", func() {
			// cron.AddFunc("@every 5s", func() {
			delOldStatusMcDetectionFromElasticRepository := del_old_log_status_mc_detection_from_elastic.NewRepository(connection.ElasticSearch())
			delOldStatusMcDetectionFromElasticService := del_old_log_status_mc_detection_from_elastic.NewService(delOldStatusMcDetectionFromElasticRepository)

			err := delOldStatusMcDetectionFromElasticService.DelOneMonthOldStatusMcDetectionLogs()
			if err != nil {
				log_function.Error("Error delete status mc detection:", err)
			}

			log_function.Info("delete status mc detection in elastic berhasil dilakukan")
		})

		cron.Start()
	}()

	// wait
	wg.Wait()
}
