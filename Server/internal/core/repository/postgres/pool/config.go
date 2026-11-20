package core_postgres_pool

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host     string        `envconfig:"HOST" requiresd:"true"`
	Port     string        `envconfig:"PORT" default:"5432"`
	User     string        `envconfig:"USER" requiresd:"true"`
	Password string        `envconfig:"PASSWORD" requiresd:"true"`
	Database string        `envconfig:"DB" requiresd:"true"`
	Timeout  time.Duration `envconfig:"TIMEOUT" requiresd:"true"`
}


func NewConfig() (Config,error){
	var config Config

	if err := envconfig.Process("POSTGRES", &config); err != nil{
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}

	return config,nil
}

func NewCofigMusst() Config{
	config, err := NewConfig()
	if err != nil{
		err = fmt.Errorf("getPostgres connection pool config: %w", err)
		panic(err)
	}
	return config
}