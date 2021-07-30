package config

type configRepository struct {
	//Port port number for the app to run
	Port int `env:"port" envdefault:"8080"`
	//Enviornment APP ENVIORNMENT
	Enviornment string `env:"enviornment" envdefault:"development"`

	// Google Places key
	GooglePlacesKey string `env:"services.google.places.key" envdefault:"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX4"`

	// Redis Connection Url
	RedisConnectionURL string `env:"redis.connection.url" envdefault:"localhost:6379"`
}
