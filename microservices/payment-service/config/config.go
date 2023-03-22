package config

import "github.com/spf13/viper"

type Config struct {
	Port      string `mapstructure:"PORT"`
	LogLevel  string `mapstructure:"LOG_LEVEL"`
	KafkaAddr string `mapstructure:"KAFKA_ADDR"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = viper.Unmarshal(c)

	return c, err
}
