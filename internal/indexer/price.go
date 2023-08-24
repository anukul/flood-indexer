package indexer

import (
	"context"
	"errors"
	"strings"

	"github.com/machinebox/graphql"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

const (
	UniswapSubgraphEndpoint = "https://api.thegraph.com/subgraphs/name/ianlapham/arbitrum-minimal"
)

var (
	Stablecoins = []string{
		// USDC
		"0xff970a61a04b1ca14834a43f5de4533ebddb5cc8",
		// USDT
		"0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9\n",
		// DAI
		"0xda10009cbd5d07dd0cecc66161fc93d7c9000da1",
	}
	ErrPoolNotFound = errors.New("pool not found")
)

func getPrice(token string, blockNumber uint64) (float64, error) {
	token = strings.ToLower(token)
	if slices.Contains(Stablecoins, token) {
		return 1, nil
	}
	for _, stablecoin := range Stablecoins {
		token0, token1 := stablecoin, token
		if strings.Compare(token0, token1) > 0 {
			token0, token1 = token1, token0
		}
		priceToken0, priceToken1, err := fetchPrices(token0, token1, blockNumber)
		if err == ErrPoolNotFound {
			continue
		} else if err != nil {
			return 0, err
		}
		var price float64
		if strings.Compare(token, token0) == 0 {
			price = priceToken0
		} else {
			price = priceToken1
		}
		return price, nil
	}
	return 0, ErrPoolNotFound
}

type PoolResponse struct {
	Token0Price  float64 `json:"token0Price,string"`
	Token1Price  float64 `json:"token1Price,string"`
	VolumeToken0 float64 `json:"volumeToken0,string"`
	VolumeToken1 float64 `json:"volumeToken1,string"`
}

func fetchPrices(token0, token1 string, blockNumber uint64) (float64, float64, error) {
	client := graphql.NewClient(UniswapSubgraphEndpoint)
	req := graphql.NewRequest(`
query ($token0: String!, $token1: String!, $blockNumber: Int!) {
  pools(
    where: {
		token0: $token0,
		token1: $token1,
    },
	block: {
		number: $blockNumber
	},
  ) {
	token0Price,
	token1Price
	volumeToken0,
	volumeToken1,
  }
}
	`)
	req.Var("token0", token0)
	req.Var("token1", token1)
	req.Var("blockNumber", blockNumber)

	var response map[string][]PoolResponse
	if err := client.Run(context.TODO(), req, &response); err != nil {
		return 0, 0, err
	}
	pools, ok := response["pools"]
	if !ok || len(pools) == 0 {
		return 0, 0, ErrPoolNotFound
	}
	priceToken0 := lo.Reduce(pools, func(agg float64, item PoolResponse, _ int) float64 {
		return agg + item.Token0Price*item.VolumeToken0
	}, 0.00) / lo.Reduce(pools, func(agg float64, item PoolResponse, _ int) float64 {
		return agg + item.VolumeToken0
	}, 0.00)
	priceToken1 := lo.Reduce(pools, func(agg float64, item PoolResponse, _ int) float64 {
		return agg + item.Token1Price*item.VolumeToken1
	}, 0.00) / lo.Reduce(pools, func(agg float64, item PoolResponse, _ int) float64 {
		return agg + item.VolumeToken1
	}, 0.00)
	return priceToken0, priceToken1, nil
}
