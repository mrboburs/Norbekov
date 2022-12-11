package config

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically

	"github.com/spf13/cast"
)

var (
	instance *Configuration
	once     sync.Once
)

// Config ...
func Config() *Configuration {
	once.Do(func() {
		instance = load()
	})

	return instance
}

// Configuration ...
type Configuration struct {
	AppName     string
	AppVersion  string
	AppURL      string
	Environment string
	ServerPort  int
	ServerHost  string

	AccessTokenDuration        time.Duration
	RefreshTokenDuration       time.Duration
	RefreshPasswdTokenDuration time.Duration

	CasbinConfigPath    string
	MiddlewareRolesPath string

	// context timeout in seconds
	CtxTimeout        int
	SigninKey         string
	ServerReadTimeout int

	JWTSecretKey              string
	JWTSecretKeyExpireMinutes int
	JWTRefreshKey             string
	JWTRefreshKeyExpireHours  int

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	PostgresSSLMode  string

	MinioAccessKeyID string
	MinioSecretKey   string
	MinioEndpoint    string
	MinioBucketName  string
	MinioLocation    string
	MinioHost        string
	MinioPort        int
	MinioUSeSSL      bool
}

func load() *Configuration {
	return &Configuration{
		AppName:             cast.ToString(getOrReturnDefault("APP_NAME", "Navoi Taxi")),
		AppVersion:          cast.ToString(getOrReturnDefault("APP_VERSION", "1.0")),
		AppURL:              cast.ToString(getOrReturnDefault("APP_URL", "localhost:8000")),
		ServerHost:          cast.ToString(getOrReturnDefault("SERVER_HOST", "localhost")),
		ServerPort:          cast.ToInt(getOrReturnDefault("SERVER_PORT", "9000")),
		Environment:         cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop")),
		CtxTimeout:          cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7)),
		CasbinConfigPath:    cast.ToString(getOrReturnDefault("CASBIN_CONFIG_PATH", "./config/rbac_model.conf")),
		MiddlewareRolesPath: cast.ToString(getOrReturnDefault("MIDDLEWARE_ROLES_PATH", "./config/models.csv")),
		SigninKey:           cast.ToString(getOrReturnDefault("SIGNIN_KEY", "")),
		ServerReadTimeout:   cast.ToInt(getOrReturnDefault("SERVER_READ_TIMEOUT", "")),

		JWTSecretKey:              cast.ToString(getOrReturnDefault("JWT_SECRET_KEY", "")),
		JWTSecretKeyExpireMinutes: cast.ToInt(getOrReturnDefault("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT", 720)),
		JWTRefreshKey:             cast.ToString(getOrReturnDefault("JWT_REFRESH_KEY", "")),
		JWTRefreshKeyExpireHours:  cast.ToInt(getOrReturnDefault("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT", 24)),
		PostgresHost:              cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost")),
		PostgresPort:              cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432)),
		PostgresDatabase:          cast.ToString(getOrReturnDefault("POSTGRES_DB", "")),
		PostgresUser:              cast.ToString(getOrReturnDefault("POSTGRES_USER", "")),
		PostgresPassword:          cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "")),
		PostgresSSLMode:           cast.ToString(getOrReturnDefault("POSTGRES_SSLMODE", "disable")),
		MinioAccessKeyID:          cast.ToString(getOrReturnDefault("MINIO_ACCESS_KEY_ID", "")),
		MinioSecretKey:            cast.ToString(getOrReturnDefault("MINIO_SECRET_KEY", "")),
		MinioEndpoint:             cast.ToString(getOrReturnDefault("MINIO_ENDPOINT", "")),
		MinioBucketName:           cast.ToString(getOrReturnDefault("MINIO_BUCKET_NAME", "")),
		MinioLocation:             cast.ToString(getOrReturnDefault("MINIO_LOCATION", "")),
		MinioHost:                 cast.ToString(getOrReturnDefault("MINIO_HOST", "")),
		MinioPort:                 cast.ToInt(getOrReturnDefault("MINIO_PORT", "")),
		MinioUSeSSL:               cast.ToBool(getOrReturnDefault("MINIO_USE_SSL", false)),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
