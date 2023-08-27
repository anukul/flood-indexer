package indexer

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"golang.org/x/exp/slog"

	"github.com/anukul/flood-indexer/internal/chain"
	"github.com/anukul/flood-indexer/internal/types"
)

func (idx *Indexer) getPastOrders(ordersChannel chan *types.Order, fromBlock, toBlock uint64) error {
	pastOrderEvents, err := idx.getPastOrderEvents(fromBlock, toBlock, idx.cfg.RpcBatchSize)
	if err != nil {
		return err
	}
	for _, orderEvent := range pastOrderEvents {
		order, err := idx.getOrder(orderEvent)
		if err != nil {
			return err
		}
		ordersChannel <- order
	}
	return nil
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

func (idx *Indexer) listenForNewOrders(ctx context.Context, ordersChannel chan *types.Order) error {
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
			ordersChannel <- order
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
	ordersChannel := make(chan *types.Order)
	go func() {
		if err := idx.getPastOrders(ordersChannel, fromBlock, toBlock); err != nil {
			slog.Error(err.Error())
			close(ordersChannel)
			return
		}
	}()
	go func() {
		if err := idx.listenForNewOrders(ctx, ordersChannel); err != nil {
			slog.Error(err.Error())
			close(ordersChannel)
			return
		}
	}()
	for order := range ordersChannel {
		if err := idx.db.InsertOrders([]*types.Order{order}); err != nil {
			return err
		}
		if err := idx.db.SetLastIndexedBlock("orders", order.BlockNumber); err != nil {
			return err
		}
	}
	return nil
}
