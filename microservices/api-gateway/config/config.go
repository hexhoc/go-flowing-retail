package config

import "github.com/spf13/viper"

type Config struct {
	Port           string `mapstructure:"PORT"`
	AuthServiceUrl string `mapstructure:"AUTH_SERVICE_URL"`
}

func LoadConfig() (Config, error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	// AutomaticEnv makes Viper check if environment variables match any of the existing keys
	// (config, default or flags). If matching env vars are found, they are loaded into Viper.
	viper.AutomaticEnv()

	c := Config{}

	err := viper.ReadInConfig()
	if err != nil {
		return c, err
	}

	err = viper.Unmarshal(&c)
	return c, err
}
