package config

type Config struct {
	Http  HttpOptions  `mapstructure:"http"  json:"http" yaml:"http"`
	Mysql MysqlOptions `mapstructure:"mysql"  json:"mysql" yaml:"mysql"`
}

type HttpOptions struct {
	Mode   string `mapstructure:"mode" json:"mode" yaml:"mode"`
	Listen int    `mapstructure:"listen" json:"listen" yaml:"listen"`
}

type MysqlOptions struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname   string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
