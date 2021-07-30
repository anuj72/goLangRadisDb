package config

import (
	"fmt"
	"github.com/mountyco/envconf"
	"sync"
)

var (
	//Repository it containts application config
	repository Repository
	once       sync.Once
)

type env struct {
	config configRepository
}

func NewRepository() Repository {
	once.Do(func() {
		conf := configRepository{}
		err := envconf.Load(&conf)
		if err != nil {
			panic(fmt.Errorf("could not initialize config: %s", err.Error()))
		}
		repository = env{config: conf}
	})
	return repository
}

func (e env) Port() int {
	return e.config.Port
}

func (e env) Enviornment() string {
	return e.config.Enviornment
}

func (e env) GooglePlacesKey() string {
	return e.config.GooglePlacesKey
}

func (e env) RedisConnectionURL() string {
	return e.config.RedisConnectionURL
}
