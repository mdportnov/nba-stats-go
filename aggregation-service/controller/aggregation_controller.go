package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mdportnov/common/dto"
	"nba-stats/redis"
	"net/http"
	"strconv"
)

func GetPlayerAverage(c *gin.Context) {
	playerID, _ := strconv.Atoi(c.Param("player_id"))
	cachedData, err := redis.GetCache(fmt.Sprintf("player:%d:average", playerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get player average"})
		return
	}
	var playerAverage map[string]float64
	json.Unmarshal([]byte(cachedData), &playerAverage)
	c.JSON(http.StatusOK, playerAverage)
}

func GetTeamAverage(c *gin.Context) {
	teamID, _ := strconv.Atoi(c.Param("team_id"))
	cachedData, err := redis.GetCache(fmt.Sprintf("team:%d:average", teamID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get team average"})
		return
	}
	var teamAverage dto.TeamAverage
	json.Unmarshal([]byte(cachedData), &teamAverage)
	c.JSON(http.StatusOK, teamAverage)
}
