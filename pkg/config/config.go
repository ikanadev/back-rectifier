package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// Server stores server configuration
type Server struct {
	Port                int  `yaml:"port"`
	Debug               bool `yaml:"debug"`
	ReadTimeoutSeconds  int  `yaml:"readTimeoutSeconds"`
	WriteTimeoutSeconds int  `yaml:"writeTimeoutSeconds"`
}

// DB database configuration
type DB struct {
	LogQueries     bool `yaml:"logQueries"`
	TimeoutSeconds int  `yaml:"timeoutSeconds"`
	Username       string
	DBName         string
	Password       string
	Host           string
}

// JWT json web token configuration
type JWT struct {
	DurationMinutes        int    `yaml:"durationMinutes"`
	RefreshDurationMinutes int    `yaml:"refreshDurationMinutes"`
	MaxRefreshMinutes      int    `yaml:"maxRefreshMinutes"`
	SigningAlgorithm       string `yaml:"signingAlgorithm"`
	MinSecretLength        int    `yaml:"minSecretLength"`
	Key                    string
}

// AWS all aws config vars
type AWS struct {
	BucketName string
}

// Config main application config
type Config struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"database"`
	JWT    JWT    `yaml:"jwt"`
	AWS    AWS
}

func readFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// GetConfig given a path with yaml config, returns the parsed version of it
func GetConfig(path string) (Config, error) {
	config := Config{}
	fileContent, err := readFile(path)
	if err != nil {
		return config, err
	}
	if err := yaml.Unmarshal(fileContent, &config); err != nil {
		return config, err
	}
	config.DB.Username = os.Getenv("POSTGRES_USER")
	config.DB.Password = os.Getenv("POSTGRES_PASSWORD")
	config.DB.DBName = os.Getenv("POSTGRES_DB")
	config.DB.Host = os.Getenv("POSTGRES_HOST")
	config.JWT.Key = os.Getenv("JWT_SECRET")
	config.AWS = AWS{BucketName: os.Getenv("AWS_BUCKET_NAME")}
	return config, nil
}
