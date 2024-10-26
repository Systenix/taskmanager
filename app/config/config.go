package config

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/errors"
	"cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/yaml"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	TaskQueueSystemUseCase string = "task_queue_system_use_case"
	TaskQueueSystemData    string = "task_queue_system_data"
	NoSqlDB                string = "no_sql_db"
	Redis                  string = "redis"
)

// AppConfig represents the application config
type AppConfig struct {
	ServiceListener Listener      `mapstructure:"ServiceListener"`
	UseCaseConfig   UseCaseConfig `mapstructure:"useCase"`
}

// Listener represents the service's listener
type Listener struct {
	Code           string   `mapstructure:"code"`
	Protocol       string   `mapstructure:"protocol"`
	Address        string   `mapstructure:"address"`
	TrustedProxies []string `mapstructure:"proxies"`
}

// UseCaseConfig represents the service's use cases
type UseCaseConfig struct {
	TaskQueue TaskQueueUseCaseConfig `mapstructure:"task_queue_system"`
}

// TaskQueueUseCaseConfig represents the task queue use case
type TaskQueueUseCaseConfig struct {
	Code                string     `mapstructure:"code"`
	TaskQueueDataConfig DataConfig `mapstructure:"TaskQueueSystemDataConfig"`
}

// DataConfig represents a data service
type DataConfig struct {
	Code            string          `mapstructure:"code"`
	DataStoreConfig DataStoreConfig `mapstructure:"dataStoreConfig"`
}

// DataStoreConfig represents handlers for data store. It can be a database, a gRPC connection, or a HTTP connection
type DataStoreConfig struct {
	Code string `mapstructure:"code"`
	// For database, this is datasource name; for grpc, it is target url
	Address string `mapstructure:"address"`
	// To indicate whether support transaction or not. "true" means supporting transaction
	Tx bool `mapstructure:"tx"`
	// Password for database connection
	Password string `mapstructure:"password"`
	// Db number for redis database
	Db int `mapstructure:"db"`
}

var (
	config        map[string]interface{}
	decodedConfig interface{}
)

func ValidateConfig(cfgFile string) {
	if cfgFile == "" {
		cfgFile = "./app/config/config_dev.yml"
	}

	ctx := cuecontext.New()

	// load schema
	bInst := load.Instances([]string{"./app/config/config.cue"}, nil)
	insts, err := ctx.BuildInstances(bInst)
	if err != nil {
		log.Fatal(err.Error())
	}
	schemaVal := insts[0]

	// load config file
	yamlFile, err := yaml.Extract(cfgFile, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	bOpts := []cue.BuildOption{}
	yamlVal := ctx.BuildFile(yamlFile, bOpts...)

	// check config against schema
	unifiedVals := schemaVal.Unify(yamlVal)
	opts := []cue.Option{
		cue.Attributes(true),
		cue.Definitions(true),
		cue.Hidden(true),
		cue.Concrete(true),
	}
	err = unifiedVals.Validate(opts...)

	// report errors
	if err != nil {
		p := message.NewPrinter(language.Tag{})
		format := func(w io.Writer, format string, args ...interface{}) {
			p.Fprintf(w, format, args...)
		}

		cwd, osErr := os.Getwd()
		if osErr != nil {
			log.Fatal(err.Error())
		}
		w := &bytes.Buffer{}
		errors.Print(w, err, &errors.Config{
			Format:  format,
			Cwd:     cwd,
			ToSlash: false,
		})

		log.Fatal("Incorrect service configuration: " + w.String())
	}
}

func LoadConfig(envPrefix, cfgFile string) {
	if cfgFile == "" {
		viper.SetConfigName("config_dev")
		viper.AddConfigPath("./app/config")
	} else {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", zap.String("err", err.Error()))
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("unable to decode into struct", zap.String("err", err.Error()))
	}
}

func GetConfig[T any]() *T {
	if config == nil {
		log.Fatal("config not loaded")
	}

	if decodedConfig == nil {
		var out T
		decoderCfg := &mapstructure.DecoderConfig{
			Result:           &out,
			DecodeHook:       mapstructure.StringToTimeDurationHookFunc(),
			WeaklyTypedInput: true,
		}
		decoder, err := mapstructure.NewDecoder(decoderCfg)
		if err != nil {
			log.Fatal("unable to create decoder", zap.String("err", err.Error()))
		}
		if err := decoder.Decode(config); err != nil {
			log.Fatal("unable to decode into struct", zap.String("err", err.Error()))
		}

		decodedConfig = &out
	}

	return decodedConfig.(*T)
}
