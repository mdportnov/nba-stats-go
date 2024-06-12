package repository

import (
	"context"
	db "github.com/mdportnov/common/db/sqlc"
	"github.com/mdportnov/common/dto"
)

func SaveStat(ctx context.Context, q *db.Queries, stat *db.Stat) error {
	return q.SaveStat(ctx, db.SaveStatParams{
		PlayerID:      stat.PlayerID,
		TeamID:        stat.TeamID,
		Points:        stat.Points,
		Rebounds:      stat.Rebounds,
		Assists:       stat.Assists,
		Steals:        stat.Steals,
		Blocks:        stat.Blocks,
		Fouls:         stat.Fouls,
		Turnovers:     stat.Turnovers,
		MinutesPlayed: stat.MinutesPlayed,
	})
}

func GetPlayerSeasonAverage(ctx context.Context, q *db.Queries, playerID int) (dto.PlayerAverage, error) {
	result, err := q.GetPlayerSeasonAverage(ctx, int32(playerID))
	if err != nil {
		return dto.PlayerAverage{}, err
	}
	return dto.PlayerAverage{
		Points:        result.Points,
		Rebounds:      result.Rebounds,
		Assists:       result.Assists,
		Steals:        result.Steals,
		Blocks:        result.Blocks,
		Fouls:         result.Fouls,
		Turnovers:     result.Turnovers,
		MinutesPlayed: result.MinutesPlayed,
	}, nil
}

func GetTeamSeasonAverage(ctx context.Context, q *db.Queries, teamID int) (dto.TeamAverage, error) {
	result, err := q.GetTeamSeasonAverage(ctx, int32(teamID))
	if err != nil {
		return dto.TeamAverage{}, err
	}
	return dto.TeamAverage{
		Points:        result.Points,
		Rebounds:      result.Rebounds,
		Assists:       result.Assists,
		Steals:        result.Steals,
		Blocks:        result.Blocks,
		Fouls:         result.Fouls,
		Turnovers:     result.Turnovers,
		MinutesPlayed: result.MinutesPlayed,
	}, nil
}

func GetAllPlayerIDs(ctx context.Context, q *db.Queries) ([]int32, error) {
	ids, err := q.GetAllPlayerIDs(ctx)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func GetAllTeamIDs(ctx context.Context, q *db.Queries) ([]int32, error) {
	ids, err := q.GetAllTeamIDs(ctx)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
