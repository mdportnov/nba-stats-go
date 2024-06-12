-- name: SaveStat :exec
INSERT INTO stats (player_id, team_id, points, rebounds, assists, steals, blocks, fouls, turnovers, minutes_played,
                   created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW());

-- name: GetPlayerSeasonAverage :one
SELECT AVG(points)         as points,
       AVG(rebounds)       as rebounds,
       AVG(assists)        as assists,
       AVG(steals)         as steals,
       AVG(blocks)         as blocks,
       AVG(fouls)          as fouls,
       AVG(turnovers)      as turnovers,
       AVG(minutes_played) as minutes_played
FROM stats
WHERE player_id = $1;

-- name: GetTeamSeasonAverage :one
SELECT AVG(points)         as points,
       AVG(rebounds)       as rebounds,
       AVG(assists)        as assists,
       AVG(steals)         as steals,
       AVG(blocks)         as blocks,
       AVG(fouls)          as fouls,
       AVG(turnovers)      as turnovers,
       AVG(minutes_played) as minutes_played
FROM stats
WHERE team_id = $1;

-- name: GetAllPlayerIDs :many
SELECT DISTINCT player_id
FROM stats;

-- name: GetAllTeamIDs :many
SELECT DISTINCT team_id
FROM stats;