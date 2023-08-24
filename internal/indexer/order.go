package indexer

import (
	"context"
	"errors"
	"math/big"

	"github.com/anukul/flood-indexer/internal/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/anukul/flood-indexer/internal/chain"
)

func parseOrder(order interface{}) (struct {
	Offerer common.Address `json:"offerer"`
	Zone    common.Address `json:"zone"`
	Offer   []struct {
		Token  common.Address `json:"token"`
		Amount *big.Int       `json:"amount"`
	} `json:"offer"`
	Consideration []struct {
		Token  common.Address `json:"token"`
		Amount *big.Int       `json:"amount"`
	} `json:"consideration"`
	Deadline *big.Int `json:"deadline"`
	Nonce    *big.Int `json:"nonce"`
}, bool) {
	res, ok := order.(struct {
		Offerer common.Address `json:"offerer"`
		Zone    common.Address `json:"zone"`
		Offer   []struct {
			Token  common.Address `json:"token"`
			Amount *big.Int       `json:"amount"`
		} `json:"offer"`
		Consideration []struct {
			Token  common.Address `json:"token"`
			Amount *big.Int       `json:"amount"`
		} `json:"consideration"`
		Deadline *big.Int `json:"deadline"`
		Nonce    *big.Int `json:"nonce"`
	})
	return res, ok
}

func convertItems(items []struct {
	Token  common.Address `json:"token"`
	Amount *big.Int       `json:"amount"`
}) []chain.IFloodPlainItem {
	var result []chain.IFloodPlainItem
	for _, item := range items {
		result = append(result, chain.IFloodPlainItem{
			Token:  item.Token,
			Amount: item.Amount,
		})
	}
	return result
}

func (idx *Indexer) getChainOrder(
	txHash common.Hash,
) (*chain.IFloodPlainOrder, error) {
	bookAbi, err := chain.FloodPlainL2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	tx, _, err := idx.chain.TransactionByHash(context.TODO(), txHash)
	if err != nil {
		return nil, err
	}
	decoderId := tx.Data()[1]
	decoderAddress, err := idx.book.GetDecoder(&bind.CallOpts{}, big.NewInt(int64(decoderId)))
	// TODO: use abigen to staticcall DecoderWithRegistry fallback, once supported by ethclient
	// see https://github.com/ethereum/go-ethereum/issues/26254#issuecomment-1333476672
	callData, err := idx.chain.CallContract(
		context.TODO(),
		ethereum.CallMsg{
			To:   &decoderAddress,
			Data: tx.Data()[2:],
		},
		nil,
	)
	method, err := bookAbi.MethodById(callData[:4])
	if err != nil {
		return nil, err
	}
	args, err := method.Inputs.Unpack(callData[4:])
	if err != nil {
		return nil, err
	}
	order, ok := parseOrder(args[0])
	if !ok {
		return nil, errors.New("failed to parse order from transaction input")
	}
	return &chain.IFloodPlainOrder{
		Offerer:       order.Offerer,
		Zone:          order.Zone,
		Offer:         convertItems(order.Offer),
		Consideration: convertItems(order.Consideration),
		Deadline:      order.Deadline,
		Nonce:         order.Nonce,
	}, nil
}

func (idx *Indexer) getOrderValue(order *chain.IFloodPlainOrder, blockNumber uint64) (float64, error) {
	orderValue := 0.00
	for _, item := range order.Offer {
		// TODO: cache price
		price, err := getPrice(item.Token.String(), blockNumber)
		if err != nil {
			return 0.00, err
		}
		token, err := chain.NewERC20(item.Token, idx.chain)
		if err != nil {
			return 0.00, err
		}
		decimals, err := token.Decimals(&bind.CallOpts{})
		if err != nil {
			return 0.00, err
		}
		amount := big.NewInt(0).Div(
			item.Amount,
			big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil),
		)
		value, _ := big.NewFloat(0).Mul(big.NewFloat(0).SetInt(amount), big.NewFloat(price)).Float64()
		orderValue += value
	}
	return orderValue, nil
}

func (idx *Indexer) getOrder(orderEvent *chain.FloodPlainL2OrderFulfilled) (*types.Order, error) {
	rawOrder, err := idx.getChainOrder(orderEvent.Raw.TxHash)
	if err != nil {
		return nil, err
	}
	orderValue, err := idx.getOrderValue(rawOrder, orderEvent.Raw.BlockNumber)
	if err != nil {
		return nil, err
	}
	blockHeader, err := idx.chain.HeaderByHash(context.TODO(), orderEvent.Raw.BlockHash)
	if err != nil {
		return nil, err
	}
	order := &types.Order{
		UserAddress: orderEvent.Offerer.String(),
		TxHash:      orderEvent.Raw.TxHash.String(),
		BlockNumber: orderEvent.Raw.BlockNumber,
		UsdValue:    orderValue,
		BlockTime:   blockHeader.Time,
	}
	return order, nil
}
