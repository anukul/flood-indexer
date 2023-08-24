package database

import (
	"database/sql"
	"time"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/samber/lo"

	"github.com/anukul/flood-indexer/internal/config"
	"github.com/anukul/flood-indexer/internal/database/flood/public/model"
	. "github.com/anukul/flood-indexer/internal/database/flood/public/table"
	"github.com/anukul/flood-indexer/internal/types"
)

type Client struct {
	conn *sql.DB
}

func NewClient(cfg *config.Config) (*Client, error) {
	connection, err := sql.Open("pgx", cfg.DatabaseUrl)
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: connection,
	}, nil
}

func (c *Client) GetLastIndexedBlock(entity string) (int, error) {
	query := IndexingStatus.SELECT(IndexingStatus.LastBlock).WHERE(IndexingStatus.Entity.EQ(String(entity)))
	var result model.IndexingStatus
	if err := query.Query(c.conn, &result); err == qrm.ErrNoRows {
		return -1, nil
	} else if err != nil {
		return -1, err
	}
	return int(*result.LastBlock), nil
}

func (c *Client) SetLastIndexedBlock(entity string, block uint64) error {
	stmt := IndexingStatus.UPDATE(IndexingStatus.LastBlock).SET(block).WHERE(IndexingStatus.Entity.EQ(String(entity)))
	_, err := stmt.Exec(c.conn)
	return err
}

func (c *Client) InsertOrders(orders []*types.Order) error {
	modelOrders := lo.Map(orders, func(order *types.Order, _ int) *model.Orders {
		return &model.Orders{
			UserAddress: order.UserAddress,
			Value:       order.UsdValue,
			TxHash:      order.TxHash,
			BlockTime:   time.Unix(int64(order.BlockTime), 0),
		}
	})
	stmt := Orders.INSERT(Orders.UserAddress, Orders.Value, Orders.TxHash, Orders.BlockTime).MODELS(modelOrders)
	_, err := stmt.Exec(c.conn)
	return err
}

func (c *Client) GetLeaderboardStats(fromTime time.Time) ([]*types.UserStats, error) {
	var leaderboard []*types.UserStats
	query := "select user_address, floor(sum(value)) as total from orders where block_time >= $1 group by user_address order by total desc"
	rows, err := c.conn.Query(query, fromTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u types.UserStats
		if err := rows.Scan(&u.Address, &u.TotalUSDValue); err != nil {
			return nil, err
		}
		leaderboard = append(leaderboard, &u)
	}
	return leaderboard, nil
}
