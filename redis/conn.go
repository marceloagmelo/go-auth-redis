package redis

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/marceloagmelo/go-auth-redis/logger"
	"github.com/marceloagmelo/go-auth-redis/utils"
)

const (
	redisExpire  = 60
	redisService = "REDIS_SERVICE"
)

var pool *redis.Pool
var redisMaxIdle = 10

func init() {
	if !utils.IsEmpty(os.Getenv("REDIS_MAXIDLE")) {
		redisMaxIdle, _ = strconv.Atoi(os.Getenv("REDIS_MAXIDLE"))
	}
	pool = &redis.Pool{
		MaxIdle:     redisMaxIdle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv(redisService))
		},
	}
}

//redisConnect Conexão com o redis
func redisConnect() (redis.Conn, error) {
	c, err := redis.Dial("tcp", os.Getenv(redisService))
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Conectando com o redis", err)
		logger.Erro.Println(mensagem)
		return nil, err
	}
	return c, nil
}

//Set enviar chave para o redis
func Set(key string, value []byte) error {

	/*conn, err := redisConnect()
	if err != nil {
		return err
	}*/
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, []byte(value))
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Gravando chave no redis", err)
		logger.Erro.Println(mensagem)
		return err
	}

	conn.Do("EXPIRE", key, redisExpire) //10 Minutes

	return err
}

//Get receber conteúdo da chave do redis
func Get(key string) ([]byte, error) {

	/*conn, err := redisConnect()
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Recuperando chave no redis", err)
		logger.Erro.Println(mensagem)
		return nil, err
	}*/
	conn := pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		mensagem := fmt.Sprintf("%s [%s] no redis: %s", "Recuperando chave", key, err)
		logger.Erro.Println(mensagem)
		return data, err
	}
	return data, err
}
