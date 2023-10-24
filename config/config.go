package config

import "os"

type Config struct {
	RedisAddress  string
	RedisPassword string
	RedisDB       int
}

func LoadConfig() *Config {
	redisAddress := os.Getenv("REDIS_ADDRESS")
	if redisAddress == "" {
		redisAddress = "localhost:6379" // default
	}
	redisPassword := os.Getenv("REDIS_PASSWORD") // it's okay if it's empty
	redisDB := 0                                 // Default value, can be made configurable too

	return &Config{
		RedisAddress:  redisAddress,
		RedisPassword: redisPassword,
		RedisDB:       redisDB,
	}
}
