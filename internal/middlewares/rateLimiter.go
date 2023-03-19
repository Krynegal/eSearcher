package middlewares

import (
	"eSearcher/configs"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type redisType struct {
	Pool *redis.Pool
}

var (
	redisClient *redisType
	once        sync.Once
)

// GetRedisConn creates a Singleton for redis connection pool and returns Redis connection instance
func GetRedisConn(cfg *configs.Config) redis.Conn {
	once.Do(func() {
		redisPool := &redis.Pool{
			MaxActive: 100,
			Dial: func() (redis.Conn, error) {
				redisURL := fmt.Sprintf("%s:%s", cfg.Redis.RedisHost, cfg.Redis.RedisPort)
				rc, err := redis.Dial("tcp", redisURL)
				if err != nil {
					fmt.Println("Error connecting to redis:", err.Error())
					return nil, err
				}
				return rc, nil
			},
		}
		redisClient = &redisType{
			Pool: redisPool,
		}
	})
	return redisClient.Pool.Get()
}

// GetKey returns a Key to be stored in Redis.
// It appends the minute value of Unix time stamp to create buckets for Key
func GetKey(cfg *configs.Config, IP string) string {
	bucket := time.Now().Unix() / int64(cfg.Redis.Bucket)
	IP = IP + strconv.FormatInt(bucket, 10)
	return IP
}

// Middle ware to checkif the Threshould per IP is reached
func RateLimitMiddleWare(cfg *configs.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			conn := GetRedisConn(cfg)
			defer conn.Close()
			IPAddress := r.Header.Get("X-Real-Ip")
			if IPAddress == "" {
				IPAddress = r.Header.Get("X-Forwarded-For")
			}
			if IPAddress == "" {
				IPAddress = r.RemoteAddr
			}
			IPAddress = GetKey(cfg, IPAddress)
			fmt.Println("IP:", IPAddress)
			val, err := redis.Int(conn.Do("GET", IPAddress))
			if err != nil {
				conn.Do("SET", IPAddress, 1)
				conn.Do("EXPIRE", IPAddress, cfg.Redis.Expiry)
			} else {
				if val > cfg.Redis.Threshold {
					err = errors.New("Max Rate Limiting Reached, Please try after some time")
					w.Write([]byte(err.Error()))
					return
				}
				conn.Do("SET", IPAddress, val+1)
			}
			fmt.Println("IP count:", val)
			next.ServeHTTP(w, r)
		})
	}
}
