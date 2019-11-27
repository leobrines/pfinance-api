package currency_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/leobrines/pfinance-api/currency"
)

var db *currency.DB

func TestMain(m *testing.M) {
	redisServer, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})

	db = &currency.DB{
		Redis: redisClient,
	}

	exit := m.Run()
	os.Exit(exit)
}

func TestGetExistingCurrencyByID(t *testing.T) {
	givenDatabaseWithRedisForTesting()

	currency, err := db.GetByID(1)

	assert.Equal(t, 1, currency.Id)
	assert.NoError(t, err)
}

func TestGetNonExistingCurrencyByID(t *testing.T) {
	givenDatabaseWithRedisForTesting()

	currency, err := db.GetByID(2)

	assert.Nil(t, currency)
	assert.Error(t, err)
}

func givenDatabaseWithRedisForTesting() {
	db.Redis.FlushAll()
	db.Redis.RPush("currency", "USD")
}
