package config

import (
	"github.com/spf13/viper"
)

const PolygonSmartContract = "134B1BE34911E39A8397ec6289782989729807a4"

type Config struct {
	DBConnectionString      string `mapstructure:"BLOOCK_DB_CONNECTION_STRING"`
	APIKey                  string `mapstructure:"BLOOCK_API_KEY"`
	APIHost                 string `mapstructure:"BLOOCK_API_HOST"`
	APIPort                 string `mapstructure:"BLOOCK_API_PORT"`
	WebhookURL              string `mapstructure:"BLOOCK_WEBHOOK_URL"`
	WebhookSecretKey        string `mapstructure:"BLOOCK_WEBHOOK_SECRET_KEY"`
	WebhookEnforceTolerance bool   `mapstructure:"BLOOCK_ENFORCE_TOLERANCE"`
	PolygonProvider         string `mapstructure:"BLOOCK_POLYGON_PROVIDER"`
	DebugMode               bool   `mapstructure:"BLOOCK_API_DEBUG_MODE"`
}

func InitConfig() (*Config, error) {
	var cfg = &Config{}

	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		setDefaultConfigValues()
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		return &Config{}, err
	}

	return cfg, nil
}

func setDefaultConfigValues() {
	viper.SetDefault("bloock_db_connection_string", "file:managed?mode=memory&cache=shared&_fk=1")
	viper.SetDefault("bloock_api_key", "")
	viper.SetDefault("bloock_webhook_url", "")
	viper.SetDefault("bloock_webhook_secret_key", "")
	viper.SetDefault("bloock_api_host", "0.0.0.0")
	viper.SetDefault("bloock_api_port", "8080")
	viper.SetDefault("bloock_webhook_enforce_tolerance", false)
	viper.SetDefault("bloock_polygon_provider", "")
	viper.SetDefault("bloock_api_debug_mode", false)
}
