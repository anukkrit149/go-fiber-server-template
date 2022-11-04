package config

import (
	"go-rest-webserver-template/pkg/config"
	"go-rest-webserver-template/pkg/heimdall"
	"go-rest-webserver-template/pkg/spine/db"
	"os"
	"time"
)

var cfg Config

// Config contains application configuration values.
type Config struct {
	Core              Core
	PromServiceConfig PromServiceConfig
	Db                db.Config
	NameClient        HttpClientConfig `mapstructure:"name_client"`
}

type Core struct {
	AppEnv          string
	ServiceName     string
	Hostname        string
	Port            int
	ShutdownTimeout int
	ShutdownDelay   int
	GitCommitHash   string
	Secret          string
}

type PromServiceConfig struct {
	AppEnv          string
	ServiceName     string
	Hostname        string
	Port            int
	ShutdownTimeout int
	ShutdownDelay   int
	GitCommitHash   string
}

type HttpClientConfig struct {
	Host                  string
	Auth                  Auth
	Mock                  bool
	Timeout               time.Duration
	HttpClient            heimdall.Config
	HttpClientName        string
	HttpRetryAttempts     int
	HttpRetryWindow       time.Duration
	HttpMaxJitterInterval time.Duration
}
type Auth struct {
	Username string
	Password string
}

func InitConfig(env string) error {
	// Init config
	err := config.NewDefaultConfig().Load(env, &cfg)
	if err != nil {
		return err
	}

	// Puts git commit hash into config.
	// This is not read automatically because env variable is not in expected format.
	cfg.Core.GitCommitHash = os.Getenv("GIT_COMMIT_HASH")

	return nil
}

func GetConfig() Config {
	return cfg
}
