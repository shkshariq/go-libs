package config

import (
	"flag"
	"gitlab.com/dishserve/go/libs/util/log"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	Port     int    `yaml:"port" json:"port"`
	Debug    bool   `yaml:"debug" json:"debug"`
	Timezone string `yaml:"timezone" json:"timezone"`
	Location *time.Location
	Log      string `yaml:"log" json:"log"`
	LogDir   string `yaml:"log_dir" json:"log_dir"`
}

var (
	AppConf AppConfig
)

func init() {
	flag.Parse()
}

func (c *AppConfig) Register() {
	ParseAppConfig()
}

//ParseAppConfig Parse application configs
func ParseAppConfig() {

	DefaultConfigurator.Load(`config/app`, &AppConf, func(config interface{}) {
		conf, _ := config.(*AppConfig)

		if conf.Timezone == `` {
			log.Fatal(`App timezone cannot be empty`)
		}
	})

	if os.Getenv(`APP_PORT`) != `` {
		if port, err := strconv.Atoi(os.Getenv(`APP_PORT`)); err == nil {
			AppConf.Port = port
		}

	}

	setDefaultTimeLocation(AppConf.Timezone)
}

//Set application default timezone
func setDefaultTimeLocation(timezone string) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatal(`Cannot load time location`, err)
	}

	AppConf.Location = location
}
