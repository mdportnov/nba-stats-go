package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mdportnov/common/util"
	"log"
	"nba-stats/controller"
	"nba-stats/redis"
	"nba-stats/repository"
)

func main() {
	router := gin.Default()

	redis.InitRedis()
	//kafka.SetupConsumer()

	// Recalculation and caching of aggregated data at startup
	repository.RecalculateAndCacheAll()

	router.GET("/average/player/:player_id", controller.GetPlayerAverage)
	router.GET("/average/team/:team_id", controller.GetTeamAverage)

	port := util.GetEnv("PORT", "8081")

	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
