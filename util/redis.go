package util

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"sync"
	"time"
)

type RedisClient struct {
	Host string
	Password string
	DB int
	Client *redis.Client
}

/**
 * 连接redis
 */
func (redisClient *RedisClient) ConnectRedis()  {
	GetRedisConfig(redisClient)
	redisClient.Client = redis.NewClient(&redis.Options{
		Addr: redisClient.Host,
		Password: redisClient.Password,
		DB:       redisClient.DB,
	})
}

/**
 * 监听队列
 */
func (redisClient *RedisClient) OnQueue(wg *sync.WaitGroup,queue string,timeOut time.Duration,call func([]string,error))  {
	defer wg.Done()
	for {
		result,_ := redisClient.Client.BLPop(timeOut,queue).Result()
		call(result,nil)
	}
}

/**
 * 推入指定的队列
 */
func (redisClient *RedisClient) Push(queue string,data string)  {
	intCMd := redisClient.Client.LPush(queue,data)
	if intCMd.Err() != nil {
		Error(fmt.Sprintf("redis push error: %v",intCMd.Err()));
	}else{
		Info(fmt.Sprintf("redis push success：%v",data));
	}
}
