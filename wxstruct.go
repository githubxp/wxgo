package main

import "time"

type Jsapi struct {
	AppID     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
	Url		  string `json:"url"`
}

type Resmsg struct {
	Code int `json:"code"`
	Data Jsapi `json:"data"`
}

type Reserr struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type Configuration struct {
	Weixin  WechatConfig
	Redis	RedisConfig
}

type WechatConfig struct {
	Appid		string
	Appsecret	string
}

type RedisConfig struct {
	Host		string
	Password	string
}

type Todos []Todo

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}