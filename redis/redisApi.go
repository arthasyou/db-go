package redis

// Set key value to redis
func Set(key string, value interface{}) error {
	return Cli.Set(Ctx, key, value, 0).Err()
}

// Get key value to redis
func Get(key string) (string, error) {
	return Cli.Get(Ctx, key).Result()
}
