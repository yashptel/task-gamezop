package config

type RedisConfig struct {
	Host string `envconfig:"REDIS_HOST"`
	Pass string `envconfig:"REDIS_PASS"`
	Db   int    `envconfig:"REDIS_DB"`
}
