package controller

import (
	"errors"
	"github.com/mdportnov/common/db/sqlc"
)

func ValidateStat(stat db.Stat) error {
	if stat.Points < 0 || stat.Rebounds < 0 || stat.Assists < 0 || stat.Steals < 0 || stat.Blocks < 0 || stat.Turnovers < 0 {
		return errors.New("points, rebounds, assists, steals, blocks, and turnovers must be positive integers")
	}
	if stat.Fouls < 0 || stat.Fouls > 6 {
		return errors.New("fouls must be between 0 and 6")
	}
	if stat.MinutesPlayed < 0 || stat.MinutesPlayed > 48 {
		return errors.New("minutes played must be between 0 and 48")
	}
	return nil
}
