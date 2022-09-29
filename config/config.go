package config

import (
	"ethaddr-update/model"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Eth     model.EthConfig `json:"eth"`
	TgToken string          `json:"tgToken"`
}

var conf Config
var configName = flag.String("f", "config", "the config file")

func Init() {
	flag.Parse()

	dir, _ := filepath.Abs(`./config/`)
	// fmt.Println(dir + `/config`)
	viper.SetConfigName(*configName)
	viper.SetConfigType("json")
	viper.AddConfigPath(dir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file config: %s ", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Println(err)
	}
}

func GetConfig() Config {
	return conf
}
