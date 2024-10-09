package cnf

import (
	"errors"
	log "github.com/jlentink/yaglogger"
	"github.com/spf13/viper"
)

var (
	_config        *viper.Viper
	ConfigLocation = ""
)

func Get() *viper.Viper {
	if _config == nil {
		_config = viper.New()
		if ConfigLocation == "" {
			_config.SetConfigType("toml")
			_config.AddConfigPath("/etc")
			_config.AddConfigPath(".")
			_config.SetConfigName("membaas")
		} else {
			_config.SetConfigFile(ConfigLocation)
		}
		_config.WatchConfig()
		if err := _config.ReadInConfig(); err != nil {
			var _t0 viper.ConfigFileNotFoundError
			if ok := errors.Is(err, _t0); ok {
				log.Fatalf("%s\n", err.Error())
			}
			log.Fatalf("Error parsing cnf file. (%s)\n", err.Error())
		}
	}
	return _config
}
