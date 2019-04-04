package eredis

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

type Eredis interface {
	Set(key string, value interface{}, expire int64)(interface{}, error)
	Get(key string)(interface{}, error)
	Del(key string)(interface{}, error)
}

type eredis struct {
	Con redis.Conn
}

func NewRedis(addr string, pw string, db int, option ...redis.DialOption) Eredis{
	option = append(option, redis.DialPassword(pw), redis.DialDatabase(db))
	c, err := redis.Dial("tcp", addr, option...)
	if err != nil {
		panic(err)
	}
	return &eredis{
		Con:c,
	}
}


func (e *eredis) Set(key string, value interface{}, expire int64)(interface{}, error){
	reply, err := e.Con.Do("SET", key, value)
	if err != nil {
		return reply, err
	}
	if expire > 0 {
		e.Con.Do("EXPIRE", key, expire)
	}
	return reply, err
}

func (e *eredis) Get(key string)(interface{}, error){
	return e.Con.Do("GET", key)
}

func (e *eredis) Del(key string)(interface{}, error){
	return e.Con.Do("DEL", key)
}

func ConvStr(reply interface{})(string, error){
	if reply == nil{
		return "", errors.New("this key not exits")
	}
	return reply.(string), nil

	switch reply.(type) {
	case string:
		return reply.(string), nil
	case []byte:
		return string(reply.([]byte)), nil
	case int64:
		return strconv.FormatInt(reply.(int64), 2), nil
	}
	return "", errors.New("decode fail")
}








