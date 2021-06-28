package gredis

import "schoolServer/pkg/setting"

var RedisConn *Cacher

func Setup()  {
	options := Options{
		Addr:setting.RedisSetting.Host,
		Password:setting.RedisSetting.Password,
		MaxIdle:setting.RedisSetting.MaxIdle,
		MaxActive:setting.RedisSetting.MaxActive,
		IdleTimeout:setting.RedisSetting.IdleTimeout,
	}
	RedisConn,_= New(options)
}

