package config

import "github.com/seerwo/jushuitan/cache"

// Config config for erp
type Config struct {
	AppID          string `json:"app_id"`           //appid
	AppSecret      string `json:"app_secret"`       //appsecret
	AuthCode 	   string `json:"auth_code"`		//auth code
	Token          string `json:"token"`            //token
	EncodingAESKey string `json:"encoding_aes_key"` //EncodingAESKey
	Cache          cache.Cache
}
