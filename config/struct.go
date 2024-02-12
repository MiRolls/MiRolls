package config

type JsonConfig struct {
	Type string  `json:"type"`
	Key  KeyInfo `json:"key"`
}

type KeyInfo struct {
	Site    interface{} `json:"site"`
	DB      DBInfo      `json:"db"`
	Redis   RedisInfo   `json:"redis"`
	Mirolls MirollsInfo `json:"mirolls"`
	Web     WebInfo     `json:"web"`
}

type DBInfo struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Database string `json:"Database"`
}

type RedisInfo struct {
	Type      string `json:"type"`
	Address   string `json:"address"`
	Password  string `json:"password"`
	Database  int    `json:"database"`
	TLS       bool   `json:"tls"`
	Prefix    string `json:"prefix"`
	MaxIdle   int    `json:"maxidle"`
	MaxActive int    `json:"maxactive"`
}

type MirollsInfo struct {
	Expire int  `json:"expire"`
	Debug  bool `json:"debug"`
}

type WebInfo struct {
	Address                  string `json:"address"`
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`
}
