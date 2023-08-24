package indexer

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/anukul/flood-indexer/internal/chain"
	"github.com/anukul/flood-indexer/internal/config"
	"github.com/anukul/flood-indexer/internal/database"
)

type Indexer struct {
	cfg   *config.Config
	chain *ethclient.Client
	db    *database.Client
	book  *chain.FloodPlainL2
}

func NewIndexer(cfg *config.Config, db *database.Client) (*Indexer, error) {
	client, err := ethclient.Dial(cfg.RpcUrl)
	if err != nil {
		return nil, err
	}
	book, err := chain.NewFloodPlainL2(common.HexToAddress(cfg.FloodContractAddress), client)
	if err != nil {
		return nil, err
	}
	return &Indexer{cfg: cfg, chain: client, book: book, db: db}, nil
}
