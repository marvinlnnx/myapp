package model

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

// Config provide the config struct which will be parsed using yaml lib
// to be used by other function
type Config struct {
	App struct {
		Version     string `yaml:"version"`
		Description string `yaml:"desc"`
	} `yaml:"app"`
	PDB struct {
		Addr string `yaml:"addr"`
	} `yaml:"pdb"`
	RDB struct {
		Addr string `yaml:"addr"`
	} `yaml:"rdb"`
	Log struct {
		Path string `yaml:"path"`
	} `yaml:"log"`
	ApiServer struct {
		Bind string `yaml:"bind"`
	} `yaml:"api_server"`
	PDBClient *pgxpool.Pool
	RDBClient *redis.Client
}
