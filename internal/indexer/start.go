package indexer

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"golang.org/x/exp/slog"

	"github.com/anukul/flood-indexer/internal/chain"
	"github.com/anukul/flood-indexer/internal/types"
)

func (idx *Indexer) getPastOrders(fromBlock, toBlock uint64) ([]*types.Order, error) {
	pastOrderEvents, err := idx.getPastOrderEvents(fromBlock, toBlock, idx.cfg.RpcBatchSize)
	if err != nil {
		return nil, err
	}
	var pastOrders []*types.Order
	for _, orderEvent := range pastOrderEvents {
		order, err := idx.getOrder(orderEvent)
		if err != nil {
			return nil, err
		}
		pastOrders = append(pastOrders, order)
	}
	return pastOrders, nil
}

func (idx *Indexer) getBlockBounds(entity string) (uint64, uint64, error) {
	var fromBlock uint64
	lastIndexedBlock, err := idx.db.GetLastIndexedBlock(entity)
	if err != nil {
		return 0, 0, err
	}
	if lastIndexedBlock == -1 {
		fromBlock = idx.cfg.FloodContractStartBlock
	} else {
		fromBlock = uint64(lastIndexedBlock + 1)
	}

	latestBlock, err := idx.chain.BlockNumber(context.TODO())
	if err != nil {
		return 0, 0, err
	}
	toBlock := latestBlock
	return fromBlock, toBlock, nil
}

func (idx *Indexer) listenForNewOrders(ctx context.Context) error {
	slog.Info("listening for new order events")
	orderEventChannel := make(chan *chain.FloodPlainL2OrderFulfilled)
	subscription, err := idx.book.WatchOrderFulfilled(&bind.WatchOpts{}, orderEventChannel, nil, nil, nil)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-subscription.Err():
			return err
		case orderEvent := <-orderEventChannel:
			slog.Info("received new order event", orderEvent.OrderHash)
			order, err := idx.getOrder(orderEvent)
			if err != nil {
				return err
			}
			if err := idx.db.InsertOrders([]*types.Order{order}); err != nil {
				return err
			}
			if err := idx.db.SetLastIndexedBlock("orders", orderEvent.Raw.BlockNumber); err != nil {
				return err
			}
		}
	}
}

func (idx *Indexer) Start(ctx context.Context) error {
	fromBlock, toBlock, err := idx.getBlockBounds("orders")
	if err != nil {
		return err
	}
	slog.Info("fetching past orders events", "from", fromBlock, "to", toBlock)
	slog.Info("block bounds", "from", fromBlock, "to", toBlock)
	if idx.cfg.Testing {
		fromBlock, toBlock = uint64(112759760), uint64(112768387)
	}
	pastOrders, err := idx.getPastOrders(fromBlock, toBlock)
	if err != nil {
		return err
	}
	if err := idx.db.InsertOrders(pastOrders); err != nil {
		return err
	}
	if err := idx.db.SetLastIndexedBlock("orders", toBlock); err != nil {
		return err
	}
	return idx.listenForNewOrders(ctx)
}
