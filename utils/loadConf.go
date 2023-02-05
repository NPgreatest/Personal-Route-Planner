package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("an error occurs when load conf file:%v", err))
	}

}
