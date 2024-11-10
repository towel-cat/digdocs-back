package cfg

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Config struct {
	PgHost   string `split_words:"true"`
	PgPort   string `split_words:"true"`
	PgUser   string `split_words:"true"`
	PgPass   string `split_words:"true"`
	PgDbName string `split_words:"true"`
}

var Env Config

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("env load error: %v", err)
	}

	Env = Config{}
	if err := envconfig.Process("", &Env); err != nil {
		fmt.Printf("config error: %v", err)
		os.Exit(1)
	}
}
