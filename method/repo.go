package method

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/leobrines/pfinance-api/domain"
	. "github.com/leobrines/pfinance-api/domain"
)

const key = "method"

type Repo struct {
	Redis *redis.Client
}

func (r *Repo) GetByID(id int) (*Method, error) {
	method := &domain.Method{}

	methodjson, err := r.Redis.LIndex(key, int64(id)-1).Result()

	if err != nil {
		return nil, fmt.Errorf("Method not found")
	}

	fmt.Println(methodjson)

	err = json.Unmarshal([]byte(methodjson), method)

	if err != nil {
		return nil, err //fmt.Errorf("Method isn't json")
	}

	method.Id = id

	return method, nil
}

func (r *Repo) Create(method *Method) (*Method, error) {
	methodjson, err := json.Marshal(method)
	if err != nil {
		return nil, err
	}
	id, err := r.Redis.RPush(key, methodjson).Result()
	if err != nil {
		return nil, err
	}
	method.Id = int(id)
	return method, nil
}
