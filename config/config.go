package config

import (
	"context"
	"my-app/model"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v2"
)

func ReadConfig(confFile string) (cfg model.Config, err error) {
	cFile, err := os.ReadFile(confFile)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(cFile, &cfg)
	if err != nil {
		return
	}

	cfg.PDBClient, err = connectPDB(cfg.PDB.Addr)
	if err != nil {
		return
	}

	cfg.RDBClient, err = connectRDB(cfg.RDB.Addr)
	if err != nil {
		return
	}

	return
}

func connectPDB(addr string) (client *pgxpool.Pool, err error) {
	pgxCfg, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return
	}

	client, err = pgxpool.NewWithConfig(context.Background(), pgxCfg)
	if err != nil {
		return
	}

	return
}

func connectRDB(addr string) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:       addr,
		ClientName: "My-Backend-Server",
	})

	_, err = client.Ping(context.Background()).Result()

	return
}
