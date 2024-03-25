package config
import (
	"github.com/spf13/viper"
)
type Config struct {
	DbPort     string `mapstructure:"DB_PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbUsername string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`
	Port       string `mapstructure:"PORT"`
	// RedisHost  string `mapstructure:"REDIS_HOST"`
	// RedisPort  string `mapstructure:"REDIS_PORT"`
	// RedisDB    int    `mapstructure:"REDIS_DB"`
	RedisAddress string `mapstructure:"REDIS_ADDRESS"`
}

var envs = []string{
	"DB_PORT", "DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME",
	"PORT", "REDIS_ADDRESS",
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return nil, err
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
