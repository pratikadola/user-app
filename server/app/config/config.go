package config

import "github.com/spf13/viper"

type Config struct {
	Server   server   `mapstructure:"server"`
	Database database `mapstructure:"database"`
}

type server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type database struct {
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	DBName    string `mapstructure:"dbname"`
	Migration bool   `mapstructure:"migration" default:"false"`
}

func LoadConfig(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return
}
