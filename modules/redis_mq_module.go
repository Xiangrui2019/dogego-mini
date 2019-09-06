package modules

import (
	"dogego-mini/cache"
	"dogego-mini/global"
	"log"
)

type RedisMQ struct {
}

var RedisMQModule *RedisMQ

func (mq *RedisMQ) Publish(queuename string, message string) error {
	err := cache.CacheClient.LPush(global.QueueNameKey(queuename), message).Err()

	if err != nil {
		return err
	}

	return nil
}

func (mq *RedisMQ) Custome(queuename string, cb func(message string) error) error {
	go func() {
		for {
			message, err := cache.CacheClient.BRPop(0,
				global.QueueNameKey(queuename)).Result()

			if err != nil {
				log.Println(err)
			}

			err = cb(message[1])

			if err != nil {
				log.Printf("Execute Callback func Error: %s", err)
			}
		}
	}()

	return nil
}

func InitRedisMQModule() {
	RedisMQModule = new(RedisMQ)
}
