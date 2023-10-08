package model

type RedisConfig struct {
	Address         string `json:"address"`
	Db              int    `json:"db"`
	IdleTimeout     string `json:"idleTimeout"`
	MaxConnLifetime string `json:"maxConnLifetime"`
	WaitTimeout     string `json:"waitTimeout"`
	DialTimeout     string `json:"dialTimeout"`
	ReadTimeout     string `json:"readTimeout"`
	WriteTimeout    string `json:"writeTimeout"`
	MaxActive       int    `json:"maxActive"`
}
