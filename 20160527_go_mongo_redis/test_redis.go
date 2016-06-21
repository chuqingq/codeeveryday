package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

func main() {
	pool := &redis.Pool{
		MaxIdle:     100,              // 最大空闲连接数
		MaxActive:   1024,             // 最大活跃连接数
		Wait:        true,             // 在没连接资源情况下，等待连接
		IdleTimeout: 30 * time.Second, // 空闲连接等待最大时间
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:56379")
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	redisCli := pool.Get()
	defer redisCli.Close()

	// 增加数据
	// for i := 0; i < 100000000; i++ {
	// 	if i%10000 == 0 {
	// 		log.Printf("i: %v\n", i)
	// 	}
	// 	// ret, err := redis.Values(redisCli.Do("SADD", i, i, i+1, i+2, i+3, i+4))
	// 	_, err := redisCli.Do("SADD", i, i, i+1, i+2, i+3, i+4)
	// 	if err != nil {
	// 		log.Printf("SADD error: %v\n", err)
	// 	}
	// 	// log.Printf("ret: %v, err: %v\n", ret, err)
	// }

	totalErr := 0

	startTime := time.Now()

	for i := 0; i < 100000; i++ {
		_, err := redisCli.Do("SMEMBERS", i)
		if err != nil {
			log.Printf("SMEMBERS error: %v\n", err)
			totalErr += 1
		}
	}

	log.Printf("duration: %v, totalErr: %v", time.Now().Sub(startTime), totalErr)
}

/*
100w*5*60秒
3秒1000个指纹
1亿条数据：占用内存13.7 GB

3.9秒
*/
