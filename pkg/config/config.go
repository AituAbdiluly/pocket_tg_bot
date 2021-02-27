package config

import "github.com/spf13/viper"

type Config struct {
	TelegramToken     string
	PocketConsumerKey string
	AuthServerURL     string
	TelegramBotURL    string `mapstructure:"bot_url"`
	DBPath            string `mapstructure:"db_file"`

	Messages Messages
}

type Messages struct {
	Errors
	Responses
}

type Errors struct {
	Default      string `mapstructure:"default"`
	InvalidURL   string `mapstructure:"invalid_url"`
	Unauthorized string `mapstructure:"unauthorized"`
	UnableToSave string `mapstructure:"unable_to_save"`
}

type Responses struct {
	Start             string `mapstructure:"start"`
	AlreadyAuthorized string `mapstructure:"already_authorized"`
	SavedSuccessfully string `mapstructure:"saved_successfully"`
	UnknowCommand     string `mapstructure:"unknown_command"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &config.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &config.Messages.Errors); err != nil {
		return nil, err
	}

	if err := parseEnv(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func parseEnv(config *Config) error {
	if err := viper.BindEnv("token"); err != nil {
		return err
	}
	if err := viper.BindEnv("comsumer_key"); err != nil {
		return err
	}
	if err := viper.BindEnv("rediract_url"); err != nil {
		return err
	}

	config.TelegramToken = viper.GetString("token")
	config.PocketConsumerKey = viper.GetString("consumer_key")
	config.AuthServerURL = viper.GetString("redirect_url")

	return nil
}
