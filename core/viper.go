package core

import (
	"fmt"
	"self-discipline/global"
	"self-discipline/utils/env"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {

	v := viper.New()
	v.SetConfigName(env.Active().Value() + "_configs")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
