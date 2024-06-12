package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mdportnov/common/util"
	"log"
	"nba-stats/controller"
	"nba-stats/db"
	"nba-stats/kafka"
)

func main() {
	router := gin.Default()
	db.ConnectDatabase()
	kafka.SetupProducer()

	router.POST("/log", controller.LogStats)
	router.GET("/player/:player_id/stats", controller.GetPlayerSeasonAverage)
	router.GET("/team/:team_id/stats", controller.GetTeamSeasonAverage)
	router.GET("/player/ids", controller.GetAllPlayerIDs)
	router.GET("/team/ids", controller.GetAllTeamIDs)

	port := util.GetEnv("PORT", "8080")

	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
