package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

// Build information -ldflags .
const (
	version    string = "dev"
	commitHash string = "-"
)

var cfg *Config

// GetConfigInstance returns service config
func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}

// Database - contains all parameters database connection.
type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SslMode  string `yaml:"sslmode"`
	Driver   string `yaml:"driver"`
}

type Retranslator struct {
	ChannelSize    uint64 `yaml:"channelSize"`
	ConsumerCount  uint64 `yaml:"consumerCount"`
	ConsumeSize    uint64 `yaml:"consumeSize"`
	ConsumeTimeout int    `yaml:"consumeTimeout"`
	ProducerCount  uint64 `yaml:"producerCount"`
	WorkerCount    int    `yaml:"workerCount"`
}

// Grpc - contains parameter address grpc.
type Grpc struct {
	Port              int    `yaml:"port"`
	MaxConnectionIdle int64  `yaml:"maxConnectionIdle"`
	Timeout           int64  `yaml:"timeout"`
	MaxConnectionAge  int64  `yaml:"maxConnectionAge"`
	Host              string `yaml:"host"`
}

// Rest - contains parameter rest json connection.
type Rest struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

// Project - contains all parameters project information.
type Project struct {
	Debug       bool   `yaml:"debug"`
	Name        string `yaml:"name"`
	ServiceName string `yaml:"serviceName"`
	Environment string `yaml:"environment"`
	Version     string
	CommitHash  string
}

// Metrics - contains all parameters metrics information.
type Metrics struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
	Path string `yaml:"path"`
}

// Jaeger - contains all parameters metrics information.
type Jaeger struct {
	Service string `yaml:"service"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
}

// Kafka - contains all parameters kafka information.
type Kafka struct {
	Capacity uint64   `yaml:"capacity"`
	Topic    string   `yaml:"topic"`
	GroupID  string   `yaml:"groupId"`
	Brokers  []string `yaml:"brokers"`
}

// Status config for service.
type Status struct {
	Port          int    `yaml:"port"`
	Host          string `yaml:"host"`
	VersionPath   string `yaml:"versionPath"`
	LivenessPath  string `yaml:"livenessPath"`
	ReadinessPath string `yaml:"readinessPath"`
}

type Telemetry struct {
	GraylogPath string `yaml:"graylogPath"`
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project      Project      `yaml:"project"`
	Grpc         Grpc         `yaml:"grpc"`
	Rest         Rest         `yaml:"rest"`
	Database     Database     `yaml:"database"`
	Metrics      Metrics      `yaml:"metrics"`
	Jaeger       Jaeger       `yaml:"jaeger"`
	Kafka        Kafka        `yaml:"kafka"`
	Status       Status       `yaml:"status"`
	Telemetry    Telemetry    `yaml:"telemetry"`
	Retranslator Retranslator `yaml:"retranslator"`
}

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(filePath string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	cfg.Project.Version = version
	cfg.Project.CommitHash = commitHash

	return nil
}

func (db *Database) GetDSN() string {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		db.Host,
		db.Port,
		db.User,
		db.Password,
		db.Name,
		db.SslMode,
	)

	return dsn
}
