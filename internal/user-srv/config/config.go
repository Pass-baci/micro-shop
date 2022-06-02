package config

type Config struct {
	Host   string     `mapstructure:"Host"`
	Port   int        `mapstructure:"Port"`
	Mode   string     `mapstructure:"Mode"`
	Mysql  MysqlInfo  `mapstructure:"Mysql"`
	Logger LoggerInfo `mapstructure:"Logger"`
}

type MysqlInfo struct {
	Host     string `mapstructure:"Host"`
	Port     int    `mapstructure:"Port"`
	Database string `mapstructure:"Database"`
	Username string `mapstructure:"Username"`
	Password string `mapstructure:"Password"`
}

type LoggerInfo struct {
	EncodeTime string `mapstructure:"EncodeTime"`
	LoggerFile struct {
		LogPath    string `mapstructure:"LogPath"`
		ErrPath    string `mapstructure:"ErrPath"`
		MaxSize    int    `mapstructure:"MaxSize"`
		MaxBackups int    `mapstructure:"MaxBackups"`
		MaxAge     int    `mapstructure:"MaxAge"`
		Compress   bool   `mapstructure:"Compress"`
	} `mapstructure:"LoggerFile"`
}
