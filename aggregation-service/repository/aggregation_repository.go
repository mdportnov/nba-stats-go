package repository

import (
	"encoding/json"
	"fmt"
	db "github.com/mdportnov/common/db/sqlc"
	"github.com/mdportnov/common/dto"
	"log"
	"nba-stats/redis"
	"net/http"
)

func RecalculateAndCacheAll() {
	playerIDs, err := getAllPlayerIDs()
	if err != nil {
		log.Printf("Error getting player IDs: %s", err)
		return
	}

	for _, playerID := range playerIDs {
		stats, err := getPlayerStats(playerID)
		if err != nil {
			log.Printf("Error getting stats for player %d: %s", playerID, err)
			continue
		}
		playerAvg, err := CalculatePlayerAverage(stats)
		if err == nil {
			jsonAvg, _ := json.Marshal(playerAvg)
			err = redis.SetCache(fmt.Sprintf("player:%d:average", playerID), jsonAvg)
			if err != nil {
				log.Printf("Error setting cache for player %d: %s", playerID, err)
			} else {
				log.Printf("Successfully cached average for player %d: %s", playerID, string(jsonAvg))
			}
		} else {
			log.Printf("Error calculating average for player %d: %s", playerID, err)
		}
	}

	teamIDs, err := getAllTeamIDs()
	if err != nil {
		log.Printf("Error getting team IDs: %s", err)
		return
	}

	for _, teamID := range teamIDs {
		stats, err := getTeamStats(teamID)
		if err != nil {
			log.Printf("Error getting stats for team %d: %s", teamID, err)
			continue
		}
		teamAvg, err := CalculateTeamAverage(stats)
		if err == nil {
			jsonAvg, _ := json.Marshal(teamAvg)
			err = redis.SetCache(fmt.Sprintf("team:%d:average", teamID), jsonAvg)
			if err != nil {
				log.Printf("Error setting cache for team %d: %s", teamID, err)
			} else {
				log.Printf("Successfully cached average for team %d: %s", teamID, string(jsonAvg))
			}
		} else {
			log.Printf("Error calculating average for team %d: %s", teamID, err)
		}
	}
}

func UpdateAndCacheAggregatedData(stat db.Stat) {
	// Update aggregate data for the players
	playerStats, err := getPlayerStats(int(stat.PlayerID))
	if err == nil {
		playerAvg, err := CalculatePlayerAverage(playerStats)
		if err == nil {
			jsonAvg, _ := json.Marshal(playerAvg)
			redis.SetCache(fmt.Sprintf("player:%d:average", stat.PlayerID), jsonAvg)
		}
	}

	// Update aggregate data for the team
	teamStats, err := getTeamStats(int(stat.TeamID))
	if err == nil {
		teamAvg, err := CalculateTeamAverage(teamStats)
		if err == nil {
			jsonAvg, _ := json.Marshal(teamAvg)
			redis.SetCache(fmt.Sprintf("team:%d:average", stat.TeamID), jsonAvg)
		}
	}
}

func CalculatePlayerAverage(stats []db.Stat) (dto.PlayerAverage, error) {
	var totalPoints, totalRebounds, totalAssists, totalSteals, totalBlocks, totalFouls, totalTurnovers, totalMinutes float64

	for _, stat := range stats {
		totalPoints += float64(stat.Points)
		totalRebounds += float64(stat.Rebounds)
		totalAssists += float64(stat.Assists)
		totalSteals += float64(stat.Steals)
		totalBlocks += float64(stat.Blocks)
		totalFouls += float64(stat.Fouls)
		totalTurnovers += float64(stat.Turnovers)
		totalMinutes += stat.MinutesPlayed
	}

	count := float64(len(stats))
	return dto.PlayerAverage{
		Points:        totalPoints / count,
		Rebounds:      totalRebounds / count,
		Assists:       totalAssists / count,
		Steals:        totalSteals / count,
		Blocks:        totalBlocks / count,
		Fouls:         totalFouls / count,
		Turnovers:     totalTurnovers / count,
		MinutesPlayed: totalMinutes / count,
	}, nil
}

func CalculateTeamAverage(stats []db.Stat) (dto.TeamAverage, error) {
	var totalPoints, totalRebounds, totalAssists, totalSteals, totalBlocks, totalFouls, totalTurnovers, totalMinutes float64

	for _, stat := range stats {
		totalPoints += float64(stat.Points)
		totalRebounds += float64(stat.Rebounds)
		totalAssists += float64(stat.Assists)
		totalSteals += float64(stat.Steals)
		totalBlocks += float64(stat.Blocks)
		totalFouls += float64(stat.Fouls)
		totalTurnovers += float64(stat.Turnovers)
		totalMinutes += stat.MinutesPlayed
	}

	count := float64(len(stats))
	return dto.TeamAverage{
		Points:        totalPoints / count,
		Rebounds:      totalRebounds / count,
		Assists:       totalAssists / count,
		Steals:        totalSteals / count,
		Blocks:        totalBlocks / count,
		Fouls:         totalFouls / count,
		Turnovers:     totalTurnovers / count,
		MinutesPlayed: totalMinutes / count,
	}, nil
}

func getAllPlayerIDs() ([]int, error) {
	resp, err := http.Get("http://localhost:8080/player/ids")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ids []int
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, err
	}
	return ids, nil
}

func getAllTeamIDs() ([]int, error) {
	resp, err := http.Get("http://stats-service:8080/team/ids")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ids []int
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, err
	}
	return ids, nil
}

func getPlayerStats(playerID int) ([]db.Stat, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/player/%d/stats", playerID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var stats []db.Stat
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, err
	}
	return stats, nil
}

func getTeamStats(teamID int) ([]db.Stat, error) {
	resp, err := http.Get(fmt.Sprintf("http://stats-service:8080/team/%d/stats", teamID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var stats []db.Stat
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, err
	}
	return stats, nil
}
