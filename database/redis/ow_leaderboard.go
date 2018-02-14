package redis

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
)

type OWLeaderboardOperator interface {
	RetrieveLeaderboard() ([]redis.Z, error)
	ZAdd(string, int) error
}

var owLeaderboardOperator OWLeaderboardOperator

type OWLeaderboardOperatorImpl struct{}

func GetOWLeaderboardOperator() OWLeaderboardOperator {
	if owLeaderboardOperator == nil {
		return &OWLeaderboardOperatorImpl{}
	}
	return owLeaderboardOperator
}

func (*OWLeaderboardOperatorImpl) RetrieveLeaderboard() ([]redis.Z, error) {
	if err := shouldPerformOnRedis(); err != nil {
		log.Error().
			Str("method", "should_perform_on_redis").
			Msgf("[REDIS] Error: %s", err.Error())
		return nil, err
	}

	key := os.Getenv("REDIS_SORTED_SET_KEY")
	result, err := redisClient.ZRevRangeByScoreWithScores(key, redis.ZRangeBy{Max: "+inf", Min: "-inf"}).Result()
	if err != nil {
		log.Error().
			Str("method", "redis_retrieve_leaderboard").
			Msgf("[REDIS] Error: %s", err.Error())
		return nil, err
	}

	return result, nil
}

func (*OWLeaderboardOperatorImpl) ZAdd(battletag string, rank int) error {
	if err := shouldPerformOnRedis(); err != nil {
		log.Error().
			Str("method", "should_perform_on_redis").
			Msgf("[REDIS] Error: %s", err.Error())
		return err
	}

	err := redisClient.ZAdd(os.Getenv("REDIS_SORTED_SET_KEY"), redis.Z{Member: battletag, Score: float64(rank)})
	if err.Err() != nil {
		log.Error().
			Str("method", "zadd").
			Msgf("[REDIS] Error: %v", err.Err())
		return err.Err()
	}

	return nil
}
