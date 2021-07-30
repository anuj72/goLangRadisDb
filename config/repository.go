package config

type Repository interface {
	GooglePlacesKey() string

	// Redis Connection Url
	RedisConnectionURL() string
}
