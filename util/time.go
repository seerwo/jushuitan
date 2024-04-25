package util

import "time"

//GetCurrTS return current timestamps
func GetCurrTS() int64 {
	return time.Now().Unix()
}

//GetCurrStr return current timestamps
func GetCurrStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
