package server

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

const redisSource = "localhost:6666"

type SMSStore interface {
	Save(phone string, code int32) error
	Get(phone string) (int32, error)
	Valid(phone string, code int32) bool
}

type smsStore struct {
	pool *redis.Pool
}

func NewSMSStore() SMSStore {
	pool := openRedis(redisSource)
	return &smsStore{
		pool: pool,
	}
}

func (sms *smsStore) Save(phone string, code int32) error {
	conn := sms.pool.Get()
	defer conn.Close()
	key := "sms_" + phone
	_, err := conn.Do("SET", key, code)
	if err != nil {
		log.Println("conn.Do SET err: ", err, key, code)
		return err
	}
	expires := 5 * time.Minute
	_, err = conn.Do("EXPIRE", key, int64(expires.Seconds()))
	if err != nil {
		log.Println("conn.Do EXPIRE err: ", err)
	}
	return err
}

func (sms *smsStore) Get(phone string) (int32, error) {
	conn := sms.pool.Get()
	defer conn.Close()
	key := "sms_" + phone
	code, err := redis.Int64(conn.Do("GET", key))
	return int32(code), err
}

func (sms *smsStore) Valid(phone string, code int32) bool {
	conn := sms.pool.Get()
	defer conn.Close()

	c, err := sms.Get(phone)
	if err != nil {
		log.Println("sms.Get", c)
		return false
	}

	return c == code
}

// 连接 redis
func openRedis(source string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 5 * time.Minute,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", source)
		},
	}
}
