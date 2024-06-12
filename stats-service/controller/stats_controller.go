package controller

import (
	"github.com/gin-gonic/gin"
	db "github.com/mdportnov/common/db/sqlc"
	"nba-stats/db"
	"nba-stats/kafka"
	"nba-stats/repository"
	"net/http"
	"strconv"
)

var q = db.New(db.DB)

func LogStats(c *gin.Context) {
	var stat db.Stat
	if err := c.ShouldBindJSON(&stat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateStat(stat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.SaveStat(c, q, &stat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	kafka.ProduceStatMessage(stat)

	c.JSON(http.StatusOK, gin.H{"message": "Stat logged successfully"})
}

func GetPlayerSeasonAverage(c *gin.Context) {
	playerID, _ := strconv.Atoi(c.Param("player_id"))
	average, err := repository.GetPlayerSeasonAverage(c, q, playerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, average)
}

func GetTeamSeasonAverage(c *gin.Context) {
	teamID, _ := strconv.Atoi(c.Param("team_id"))
	average, err := repository.GetTeamSeasonAverage(c, q, teamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, average)
}

func GetAllPlayerIDs(c *gin.Context) {
	ids, err := repository.GetAllPlayerIDs(c, q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get player ids"})
		return
	}
	c.JSON(http.StatusOK, ids)
}

func GetAllTeamIDs(c *gin.Context) {
	ids, err := repository.GetAllTeamIDs(c, q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get team ids"})
		return
	}
	c.JSON(http.StatusOK, ids)
}
