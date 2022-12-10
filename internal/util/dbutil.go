package util

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v9"
)

func GetParamString(param string, defaultValue string) string {
	env := os.Getenv(param)
	if env != "" {
		return env
	}
	return defaultValue
}

func getParamInt(param string, defaultValue int) int {
	env := os.Getenv(param)
	v, err := strconv.Atoi(env)
	if err != nil {
		return defaultValue
	}
	return v
}

func GetConnectionString() string {
	host := GetParamString("MYSQL_DB_HOST", "localhost")
	port := GetParamString("MYSQL_PORT", "3306")
	user := GetParamString("MYSQL_USER", "")
	pass := GetParamString("MYSQL_PASSWORD", "")
	dbname := GetParamString("MYSQL_DATABASE", "")
	protocol := GetParamString("MYSQL_PROTOCOL", "tcp")
	dbargs := GetParamString("MYSQL_DBARGS", "")

	if dbargs != "" {
		dbargs = "?" + dbargs
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, dbname, dbargs)
}

func GetRedisOptions() *redis.Options {
	host := GetParamString("REDIS_HOST", "localhost")
	port := GetParamString("REDIS_PORT", "6379")
	pass := GetParamString("REDIS_PASSWORD", "")
	db := getParamInt("REDIS_DB", 0)

	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: pass,
		DB:       db,
	}
}

func GetKafkaAddress() string {
	host := GetParamString("KAFKA_HOST", "localhost")
	port := GetParamString("KAFKA_PORT", "9092")
	return fmt.Sprintf("%s:%s", host, port)
}

func GetMailServerAddress() string {
	host := GetParamString("MAIL_HOST", "localhost")
	port := GetParamString("MAIL_PORT", "1025")
	return fmt.Sprintf("%s:%s", host, port)
}

func getJaegerURL() string {
	host := GetParamString("JAEGER_HOST", "jaeger")
	port := GetParamString("JAEGER_PORT", "14268")
	return fmt.Sprintf("http://%s:%s/api/traces", host, port)
}
