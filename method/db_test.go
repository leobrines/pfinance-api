package method_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/leobrines/pfinance-api/method"
)

var db *method.DB

func TestMain(m *testing.M) {
	redisServer, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})

	db = &method.DB{
		Redis: redisClient,
	}

	exit := m.Run()
	os.Exit(exit)
}

func TestGetExistingMethodByID(t *testing.T) {
	givenDatabaseWithRedisForTesting()

	method, err := db.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, method.Id)
}

func TestGetNonExistingMethodByID(t *testing.T) {
	givenDatabaseWithRedisForTesting()

	method, err := db.GetByID(2)

	assert.Error(t, err)
	assert.Nil(t, method)
}

func givenDatabaseWithRedisForTesting() {
	db.Redis.FlushAll()
	db.Redis.RPush("method", "{\"user_id\":1,\"currency_id\":1,\"balance\":100}")
}
