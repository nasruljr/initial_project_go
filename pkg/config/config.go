package config

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	rootpath   = filepath.Join(filepath.Dir(b), "../../configs/")
)

func initConfig() {
	viper.SetConfigType("yml")
	viper.AddConfigPath(rootpath)
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetConfig(value string) string {
	initConfig()
	return viper.GetString(value)
}

func initMessage() {
	viper.SetConfigType("yml")
	viper.AddConfigPath(rootpath)
	viper.SetConfigName("message")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetMessage(value string) string {
	initMessage()
	return viper.GetString(value)
}
