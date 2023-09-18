package config

import (
	"github.com/spf13/viper"
)

const PolygonSmartContract = "134B1BE34911E39A8397ec6289782989729807a4"

type Config struct {
	DBConnectionString  string `mapstructure:"BLOOCK_DB_CONNECTION_STRING"`
	APIKey              string `mapstructure:"BLOOCK_API_KEY"`
	APIHost             string `mapstructure:"BLOOCK_API_HOST"`
	APIPort             string `mapstructure:"BLOOCK_API_PORT"`
	WebhookSecretKey    string `mapstructure:"BLOOCK_WEBHOOK_SECRET_KEY"`
	LocalPrivateKey     string `mapstructure:"BLOOCK_LOCAL_PRIVATE_KEY"`
	LocalPublicKey      string `mapstructure:"BLOOCK_LOCAL_PUBLIC_KEY"`
	ManagedKeyID        string `mapstructure:"BLOOCK_MANAGED_KEY_ID"`
	PublicHost          string `mapstructure:"BLOOCK_PUBLIC_HOST"`
	IssuerDidMethod     string `mapstructure:"BLOOCK_ISSUER_DID_METHOD"`
	IssuerDidBlockchain string `mapstructure:"BLOOCK_ISSUER_DID_BLOCKCHAIN"`
	IssuerDidNetwork    string `mapstructure:"BLOOCK_ISSUER_DID_NETWORK"`
	DebugMode           bool   `mapstructure:"BLOOCK_API_DEBUG_MODE"`
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
	viper.SetDefault("bloock_webhook_secret_key", "")
	viper.SetDefault("bloock_api_host", "0.0.0.0")
	viper.SetDefault("bloock_api_port", "8080")
	viper.SetDefault("bloock_local_private_key", "")
	viper.SetDefault("bloock_local_public_key", "")
	viper.SetDefault("bloock_managed_key_id", "")
	viper.SetDefault("bloock_public_host", "")
	viper.SetDefault("bloock_issuer_did_method", "")
	viper.SetDefault("bloock_issuer_did_blockchain", "")
	viper.SetDefault("bloock_issuer_did_network", "")
	viper.SetDefault("bloock_api_debug_mode", false)
}
