package method_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/leobrines/pfinance-api/method"
)

var repo *method.Repo

func TestMain(m *testing.M) {
	redisServer, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})

	repo = &method.Repo{
		Redis: redisClient,
	}

	exit := m.Run()
	os.Exit(exit)
}

func TestGetExistingMethodByID(t *testing.T) {
	givenRepoWithRedisForTesting()

	method, err := repo.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, method.Id)
}

func TestGetNonExistingMethodByID(t *testing.T) {
	givenRepoWithRedisForTesting()

	method, err := repo.GetByID(2)

	assert.Error(t, err)
	assert.Nil(t, method)
}

func givenRepoWithRedisForTesting() {
	repo.Redis.FlushAll()
	repo.Redis.RPush("method", "{\"user_id\":1,\"currency_id\":1,\"balance\":100}")
}
