package config

type Config struct {
	Host  string    `mapstructure:"Host"`
	Port  int       `mapstructure:"Port"`
	Mode  string    `mapstructure:"Mode"`
	Mysql MysqlInfo `mapstructure:"Mysql"`
}

type MysqlInfo struct {
	Host     string `mapstructure:"Host"`
	Port     int    `mapstructure:"Port"`
	Database string `mapstructure:"Database"`
	Username string `mapstructure:"Username"`
	Password string `mapstructure:"Password"`
}
