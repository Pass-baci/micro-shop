package conf

import (
	"github.com/spf13/viper"
)

func ResolveConfig(configStruct interface{}, path, fileName string) error {
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(configStruct)
	if err != nil {
		return err
	}
	return nil
}
