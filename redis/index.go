package redis

import (
	"github.com/JabinGP/ssr-sub-hub/config"
	"github.com/go-redis/redis"
	"sync"
)

var Enable bool
var Client *redis.Client
var once sync.Once

func init() {
	once.Do(func() {
		Enable = checkEnable()
		if !Enable {
			return
		}
		Client = newRedisClient()
		if Client == nil {
			panic("Rdis")
		}
	})
}

func checkEnable() bool {
	return config.Viper.GetBool("redis.enable")
}

func newRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: config.Viper.GetString("redis.host") +
			":" + config.Viper.GetString("redis.port"),
		Password: config.Viper.GetString("redis.password"),
		DB:       config.Viper.GetInt("redis.db"),
	})
	if err := pingClient(client); err != nil {
		panic(err)
	}
	return client
}

func pingClient(client *redis.Client) error {
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
