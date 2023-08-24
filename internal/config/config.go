package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	Testing              bool   `env:"TESTING" envDefault:"true"`
	ApiPort              uint64 `env:"API_PORT" envDefault:"8000"`
	DatabaseUrl          string `env:"DATABASE_URL,required"`
	RpcUrl               string `env:"RPC_URL,required"`
	RpcBatchSize         uint64 `env:"RPC_BATCH_SIZE" envDefault:"2000"`
	FloodContractAddress string `env:"FLOOD_CONTRACT_ADDRESS" envDefault:"0x00000000cf83d487A9cDfd8E8C121AF9f5ddE55B"`
	// TODO: find deployment block by binary searching over ethclient.CodeAt(address, block)
	FloodContractStartBlock uint64 `env:"FLOOD_CONTRACT_START_BLOCK" envDefault:"111508091"`
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
