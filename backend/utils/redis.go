package utils

import (
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"
)

//TODO: move this file to CONFIG

const (
	SETNX_RESCODE_FAILED  = 0
	SETNX_RESCODE_SUCCESS = 1
	SETNX_EXPIRED_TIME    = 60 * 60 * 24

	REDIS_DB_SYSTEM   = 0 //系统库（默认）
	REDIS_DB_REALTIME = 1 //实时控制库
)

type dialFuncHandle func() (redis.Conn, error)

type RedisPoolConfig struct {
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	DialFunc    dialFuncHandle
}

func defaultDialFunc() (redis.Conn, error) {
	return redis.Dial("tcp", "47.106.145.145:6379")
}

func DefaultRedisPoolConfig() *RedisPoolConfig {
	return &RedisPoolConfig{
		MaxIdle:     16,
		MaxActive:   1024,
		IdleTimeout: 300,
		DialFunc:    defaultDialFunc,
	}
}

var (
	redisIns  *RepoRedis
	redisOnce sync.Once
	repoRedis RepoRedis
)

type RepoRedis struct {
	pool *redis.Pool
}

func GetRedisRepo() RepoRedis {
	return *NewRepoRedis()
}

func NewRepoRedis() *RepoRedis {
	redisOnce.Do(func() {
		cfg := DefaultRedisPoolConfig()

		pool := &redis.Pool{
			MaxIdle:     cfg.MaxIdle,
			MaxActive:   cfg.MaxActive,
			IdleTimeout: cfg.IdleTimeout,
			Dial:        cfg.DialFunc,
		}

		redisIns = &RepoRedis{
			pool: pool,
		}
	})

	return redisIns
}

func (this *RepoRedis) HExists(key string, value string) (bool, *Error) {
	c := this.pool.Get()
	defer c.Close()

	isExists, err := redis.Bool(c.Do("HEXISTS", key, value))
	if err != nil {
		Errorf("HExists HEXISTS error: key=%s, value=%s, error=%s", key, value, err)
		return false, NewError(ErrorCodeInnerError, ErrorDescInnerError)
	}

	return isExists, nil
}

func (this *RepoRedis) Exists(key string) (bool, *Error) {
	c := this.pool.Get()
	defer c.Close()

	isExists, err := redis.Bool(c.Do("Exists", key))
	if err != nil {
		Errorf("Exists Exists error: key=%s, error=%s", key, err)
		return false, NewError(ErrorCodeInnerError, ErrorDescInnerError)
	}

	return isExists, nil
}

func (this *RepoRedis) HmsetString(name string, values map[string]string) (utilsErr *Error) {
	c := this.pool.Get()
	defer c.Close()

	_, err := c.Do("HMSET", redis.Args{}.Add(name).AddFlat(values)...)
	if err != nil {
		Errorf("HmsetInt64 HMSET error: err=%s, name=%s, value=%+v", err, name, values)
		return NewError(ErrorCodeInnerError, ErrorDescInnerError)
	}

	return nil
}

func (this *RepoRedis) SetLocker(key string, value string, expired int) (ok bool, utilsErr *Error) {
	c := this.pool.Get()
	defer c.Close()

	_, err := redis.String(c.Do("SET", key, value, "EX", expired, "NX"))
	if err == redis.ErrNil {
		return false, nil
	}
	if err != nil {
		Warnf("DeviceRepositoryRedis SetLocker SET error: key=%s, value=%s, err=%s", key, value, err)
		utilsErr = NewError(ErrorCodeInnerError, ErrorDescInnerError)
		return
	}

	ok = true
	return
}

func (this *RepoRedis) GetString(key string) (string, *Error) {
	c := this.pool.Get()
	defer c.Close()

	value, err := redis.String(c.Do("GET", key))
	if err != nil {
		Errorf("DeviceRepositoryRedis GetString GET error: key=%s, err=%s", key, err)
		if err == redis.ErrNil {
			return "", NewError(ErrorCodeNotFound, ErrorDescNotFound)
		}
		return "", NewError(ErrorCodeInnerError, ErrorDescInnerError)
	}

	return value, nil
}

func (this *RepoRedis) DelKey(key string) (utilsErr *Error) {
	c := this.pool.Get()
	defer c.Close()

	_, err := c.Do("DEL", key)
	if err != nil {
		Warnf("DeviceRepositoryRedis DelKey Do DEL error: key=%s, err=%s", key, err)
		return NewError(ErrorCodeInnerError, ErrorDescInnerError)
	}

	return nil
}
