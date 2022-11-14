package options

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/dnsjia/luban/cmd/config"
)

var (
	DB     *gorm.DB
	Config config.Config
	VP     *viper.Viper
)
