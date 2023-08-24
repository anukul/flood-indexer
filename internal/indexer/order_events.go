package indexer

import (
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"golang.org/x/exp/slog"

	"github.com/anukul/flood-indexer/internal/chain"
)

func getBlockBounds(startBlock, batchNumber, batchSize uint64) (uint64, uint64) {
	// add 1 to exclude block from previous batch
	fromBlock := startBlock + batchNumber*(batchSize) + 1
	// no previous batch for batchNumber = 0
	if batchNumber == 0 {
		fromBlock = fromBlock - 1
	}
	toBlock := fromBlock + batchSize
	return fromBlock, toBlock
}

func exhaustIterator(iterator *chain.FloodPlainL2OrderFulfilledIterator) ([]*chain.FloodPlainL2OrderFulfilled, error) {
	var orderEvents []*chain.FloodPlainL2OrderFulfilled
	for iterator.Next() {
		orderEvents = append(orderEvents, iterator.Event)
	}
	return orderEvents, iterator.Error()
}

func (idx *Indexer) getPastOrderEvents(fromBlock, toBlock, batchSize uint64) ([]*chain.FloodPlainL2OrderFulfilled, error) {
	numBatches := int(math.Ceil(float64(toBlock-fromBlock) / float64(batchSize)))
	var orderEvents []*chain.FloodPlainL2OrderFulfilled
	for i := 0; i < numBatches; i++ {
		currentFromBlock, currentToBlock := getBlockBounds(fromBlock, uint64(i), batchSize)
		currentOrderEventsIterator, err := idx.book.FilterOrderFulfilled(&bind.FilterOpts{
			Start: currentFromBlock,
			End:   &currentToBlock,
		}, nil, nil, nil)
		if err != nil {
			return nil, err
		}
		currentOrderEvents, err := exhaustIterator(currentOrderEventsIterator)
		if err != nil {
			return nil, err
		}
		orderEvents = append(orderEvents, currentOrderEvents...)
		slog.Info("batch",
			"num", fmt.Sprintf("%d/%d", i+1, numBatches),
			"from", currentFromBlock,
			"to", currentToBlock,
			"order_count", len(currentOrderEvents),
		)
	}
	return orderEvents, nil
}
