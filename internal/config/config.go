package config

import (
	"fmt"
	"github.com/bloock/bloock-sdk-go/v2"
	"github.com/mcuadros/go-defaults"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"reflect"
	"regexp"
	"strings"
)

const (
	PublicPolygonProvider = "https://polygon-mumbai.infura.io/v3/2IaPG8cEELd8sDCXUaIYTeSpAAF"
)

type APIConfig struct {
	Host       string `mapstructure:"host" default:"0.0.0.0"`
	Port       string `mapstructure:"port" default:"8080"`
	DebugMode  bool   `mapstructure:"debug_mode" default:"false"`
	PublicHost string `mapstructure:"public_host"`
}

type AuthConfig struct {
	Secret string `mapstructure:"secret"`
}

type DBConfig struct {
	ConnectionString string `mapstructure:"connection_string" default:"file:managed?mode=memory&cache=shared&_fk=1"`
}

type BloockConfig struct {
	ApiHost          string `mapstructure:"api_host" default:"https://api.bloock.com"`
	ApiKey           string `mapstructure:"api_key"`
	WebhookSecretKey string `mapstructure:"webhook_secret_key"`
}

type BlockchainConfig struct {
	SmartContract  string `mapstructure:"smart_contract" default:"0x134B1BE34911E39A8397ec6289782989729807a4"`
	Provider       string `mapstructure:"provider" default:"https://polygon.bloock.dev"`
	ResolverPrefix string `mapstructure:"resolver_prefix" default:"polygon:mumbai"`
}

type IssuerConfig struct {
	Name            string            `mapstructure:"name"`
	Description     string            `mapstructure:"description"`
	Image           string            `mapstructure:"image"`
	PublishInterval int64             `mapstructure:"publish_interval"`
	Key             KeyConfig         `mapstructure:"key"`
	IssuerDid       string            `mapstructure:"issuer_did"`
	DidMetadata     DidMetadataConfig `mapstructure:"did_metadata"`
}

type DidMetadataConfig struct {
	Method     string `mapstructure:"method"`
	Blockchain string `mapstructure:"blockchain"`
	Network    string `mapstructure:"network"`
}

type KeyConfig struct {
	Key string `mapstructure:"key"`
}

type Config struct {
	Api        APIConfig
	Auth       AuthConfig
	Db         DBConfig
	Bloock     BloockConfig
	Blockchain BlockchainConfig
	Issuer     IssuerConfig
}

var Configuration = Config{}

func InitConfig(logger zerolog.Logger) (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("bloock")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		switch err.(type) {
		default:
			return nil, fmt.Errorf("fatal error loading config file: %s", err)
		case viper.ConfigFileNotFoundError:
			logger.Warn().Msg("No config file found. Using defaults and environment variables")
		}
	}

	bindEnvs(Configuration)

	err = viper.Unmarshal(&Configuration)
	if err != nil {
		return nil, fmt.Errorf("fatal error loading config file: %s", err)
	}
	defaults.SetDefaults(&Configuration)

	bloock.ApiHost = Configuration.Bloock.ApiHost
	bloock.DisableAnalytics = true

	return &Configuration, nil
}

func bindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)

		var tv string
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			tv = toSnakeCase(t.Name)
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, tv)...)
		default:
			viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
