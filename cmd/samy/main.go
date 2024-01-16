package main

import (
	"fmt"

	_viper "github.com/jeancarloshp/samy-go/pkg/config/viper"
	"github.com/spf13/viper"
)

func main() {
	env, err := _viper.LoadViperConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(viper.GetString("ENV"))

	fmt.Println(env)
}
