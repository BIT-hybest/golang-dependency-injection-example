package mvcExample

type RedisConfig struct {
	Host string
}

type RedisClient struct {
	config RedisConfig
}

func (c *RedisClient) Get() string {
	return "testData"
}

func NewRedisClient(config RedisConfig) CacheClient {
	return &RedisClient{config: config}
}

type CacheClient interface {
	Get() string
}
