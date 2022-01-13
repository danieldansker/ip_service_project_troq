package utils

import "os"

const ENV_VAR_DB_KEY = "DB_URL"
const ENV_VAR_RATE_LIMITER_KEY = "RATE_LIMITER"

func GetConfiguration() Configuration {
	dbUrl := os.Getenv(ENV_VAR_DB_KEY)
	rateLimiter := os.Getenv(ENV_VAR_RATE_LIMITER_KEY)

	var configuration = new(Configuration)
	configuration.Db = dbUrl
	configuration.RateLimiter = rateLimiter
	return *configuration
}

type Configuration struct {
	Db, RateLimiter string
}
