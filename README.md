# NBA Stats Service

This project is designed to create a scalable system for logging and calculating NBA player and team statistics. The
system can run either on-premises or on AWS and includes two main microservices: `StatsService`
and `AggregationService`. The project uses Go with Gin, PostgreSQL, Kafka, and Redis, Docker Compose.

## Features

- Log NBA player statistics
- Calculate aggregate statistics (season average per player and team)
- Highly available and scalable architecture
- Support for batches of tens or hundreds of requests concurrently
- Up-to-date data available immediately after writing
- Maintainable and supports frequent updates

## How Processes Work

### StatsService

The `StatsService` is responsible for logging player statistics into the PostgreSQL database and publishing these events
to Kafka.

#### Logging Player Statistics:

- The service exposes a POST endpoint `/log` to accept player statistics in JSON format.
- The received statistics are validated and then saved to the PostgreSQL database.
- After saving the statistics, a message is published to a Kafka topic to notify other services about the new data.

#### Retrieving Data:

- The service exposes GET endpoints to retrieve player and team statistics, as well as their IDs.
- Additionally, endpoints are available to calculate and return the season average statistics for a player or a team.

### AggregationService

The `AggregationService` is responsible for consuming events from Kafka, updating aggregate statistics, and caching
these statistics in Redis for fast access.

#### Consuming Kafka Events:

- The service subscribes to the Kafka topic where `StatsService` publishes new statistics events.
- Upon receiving a new event, it updates the aggregate statistics for the respective player and team.

#### Updating and Caching Aggregates:

- The aggregate statistics are calculated using the data received from the Kafka events.
- These statistics are then cached in Redis for quick retrieval.
- The service can also recalculate all aggregate statistics from the database upon startup to ensure the cache is
  up-to-date.

#### Serving Aggregate Data:

- The service exposes GET endpoints to retrieve the cached aggregate statistics for players and teams.
- If the requested data is not in the cache, the service can recompute the aggregates from the database.

## Data Flow

### Logging Statistics:

1. A client sends a POST request to `StatsService` with player statistics.
2. `StatsService` validates and saves the statistics to PostgreSQL.
3. `StatsService` publishes an event to Kafka about the new statistics.

### Processing Statistics:

1. `AggregationService` consumes the event from Kafka.
2. `AggregationService` updates the aggregate statistics for the player and team.
3. The updated statistics are cached in Redis.

### Retrieving Aggregate Data:

1. A client sends a GET request to `AggregationService` to retrieve player or team aggregate statistics.
2. `AggregationService` checks Redis for cached data.
3. If the data is not cached, `AggregationService` retrieves raw statistics from `StatsService`, calculates the
   aggregates, and caches the result.

## API Endpoints

### StatsService

- `POST /log`: Log NBA player statistics
- `GET /player/:player_id/stats`: Get player stats
- `GET /team/:team_id/stats`: Get team stats
- `GET /player/ids`: Get all player IDs
- `GET /team/ids`: Get all team IDs
- `GET /average/player/:player_id`: Get player season average
- `GET /average/team/:team_id`: Get team season average

### AggregationService

- `GET /average/player/:player_id`: Get cached player season average
- `GET /average/team/:team_id`: Get cached team season average

## Configuration

Configuration is managed through environment variables. The following environment variables are required:

- `DB_USER`: Database user
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `DB_HOST`: Database host
- `DB_PORT`: Database port
- `KAFKA_BROKER`: Kafka broker address
- `REDIS_ADDR`: Redis address

### Running with Docker Compose

1. Create a `.env` file with the necessary environment variables.
2. Run the following command to start the services:

```sh
docker-compose up --build
```
