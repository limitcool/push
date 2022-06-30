package main

import (
	"fmt"

	"github.com/limitcool/push/engine"
	. "github.com/limitcool/push/global"
	"github.com/spf13/viper"
)

func init() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.Unmarshal(&Config)
}

func main() {
	p := engine.Select(Config.Enable)
	fmt.Println(p.Name(), p.GetTitle())
}
